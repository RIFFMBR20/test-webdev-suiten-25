package service

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"test-webdev-suiten-25/internal/models/dao"
	"test-webdev-suiten-25/internal/models/dto"
	"test-webdev-suiten-25/internal/models/mapper"
	"test-webdev-suiten-25/internal/repository"
)

type AttendanceService interface {
	GetAll() ([]dto.AttendanceDTO, error)
	GetByID(id int) (dto.AttendanceDTO, error)

	GetByDateAndDivision(date time.Time, divisionID int) ([]dto.AttendanceDTO, error)
	BulkUpsert(ctx context.Context, req dto.BulkAttendanceUpsertDTO) error
}

type attendanceServiceImpl struct {
	repo repository.AttendanceRepository
	log  *log.Logger
}

func ProvideAttendanceService(repo repository.AttendanceRepository, logger *log.Logger) AttendanceService {
	return &attendanceServiceImpl{
		repo: repo,
		log:  logger,
	}
}

func (s *attendanceServiceImpl) GetAll() ([]dto.AttendanceDTO, error) {
	rows, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	out := make([]dto.AttendanceDTO, 0, len(rows))
	for _, a := range rows {
		out = append(out, mapper.AttendanceToDTO(a))
	}

	return out, nil
}

func (s *attendanceServiceImpl) GetByID(id int) (dto.AttendanceDTO, error) {
	if id <= 0 {
		return dto.AttendanceDTO{}, fmt.Errorf("invalid id")
	}

	a, err := s.repo.GetByID(id)
	if err != nil {
		return dto.AttendanceDTO{}, err
	}

	return mapper.AttendanceToDTO(a), nil
}

func (s *attendanceServiceImpl) GetByDateAndDivision(date time.Time, divisionID int) ([]dto.AttendanceDTO, error) {
	if date.IsZero() {
		return nil, fmt.Errorf("invalid date")
	}
	if divisionID <= 0 {
		return nil, fmt.Errorf("invalid division_id")
	}

	rows, err := s.repo.GetByDateAndDivision(date, divisionID)
	if err != nil {
		return nil, err
	}

	out := make([]dto.AttendanceDTO, 0, len(rows))
	for _, a := range rows {
		out = append(out, mapper.AttendanceToDTO(a))
	}

	return out, nil
}

func (s *attendanceServiceImpl) BulkUpsert(ctx context.Context, req dto.BulkAttendanceUpsertDTO) error {
	if len(req.Items) == 0 {
		return nil
	}

	items := make([]dao.Attendance, 0, len(req.Items))
	for _, it := range req.Items {
		if it.EmployeeID <= 0 {
			return fmt.Errorf("invalid employee_id")
		}
		if it.Date.IsZero() {
			return fmt.Errorf("invalid date")
		}

		var homeTimePtr *time.Time
		if strings.TrimSpace(it.HomeTime) != "" {
			tm, err := parseHomeTime(it.Date, it.HomeTime)
			if err != nil {
				return err
			}
			homeTimePtr = &tm
		}

		entity := dao.Attendance{
			EmployeeID: it.EmployeeID,
			Date:       it.Date,
			HomeTime:   homeTimePtr,
			Note:       it.Note,
		}

		if entity.HomeTime != nil {
			entity.TotalOvertime = calcTotalOvertime(it.Date, entity.HomeTime)
		} else {
			entity.TotalOvertime = "1+0"
		}

		items = append(items, entity)
	}

	return s.repo.BulkUpsert(ctx, items)
}

func calcTotalOvertime(date time.Time, homeTime *time.Time) string {
	if homeTime == nil {
		return "1+0"
	}

	base := time.Date(date.Year(), date.Month(), date.Day(), 17, 0, 0, 0, date.Location())

	if homeTime.Before(base) {
		return "1+0"
	}

	diff := homeTime.Sub(base)
	overtimeHours := int(diff / time.Hour)

	workDays := 1 + (overtimeHours / 5)
	overtimeRemainder := overtimeHours % 5

	return fmt.Sprintf("%d+%d", workDays, overtimeRemainder)
}

func parseHomeTime(date time.Time, homeTime string) (time.Time, error) {
	h, m, err := parseClock(homeTime)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid home_time: %w", err)
	}

	return time.Date(date.Year(), date.Month(), date.Day(), h, m, 0, 0, date.Location()), nil
}

func parseClock(s string) (hour int, minute int, err error) {
	v := strings.TrimSpace(s)
	v = strings.ReplaceAll(v, ".", ":")

	if v == "" {
		return 0, 0, fmt.Errorf("empty")
	}

	parts := strings.Split(v, ":")
	if len(parts) == 1 {
		h, e := strconv.Atoi(parts[0])
		if e != nil {
			return 0, 0, e
		}
		if h < 0 || h > 23 {
			return 0, 0, fmt.Errorf("hour out of range")
		}
		return h, 0, nil
	}

	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid format")
	}

	h, e := strconv.Atoi(parts[0])
	if e != nil {
		return 0, 0, e
	}
	m, e := strconv.Atoi(parts[1])
	if e != nil {
		return 0, 0, e
	}

	if h < 0 || h > 23 {
		return 0, 0, fmt.Errorf("hour out of range")
	}
	if m < 0 || m > 59 {
		return 0, 0, fmt.Errorf("minute out of range")
	}

	return h, m, nil
}
