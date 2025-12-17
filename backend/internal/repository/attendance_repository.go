package repository

import (
	"context"
	"test-webdev-suiten-25/internal/models/dao"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AttendanceRepository interface {
	GetAll() ([]dao.Attendance, error)
	GetByID(id int) (dao.Attendance, error)
	Create(attendance *dao.Attendance) error
	Update(attendance *dao.Attendance) error
	Delete(id int) error
	BulkUpsert(ctx context.Context, items []dao.Attendance) error
	GetByDateAndDivision(date time.Time, divisionID int) ([]dao.Attendance, error)
}

type attendanceRepositoryImpl struct {
	db *gorm.DB
}

func ProvideAttendanceRepository(db *gorm.DB) AttendanceRepository {
	return &attendanceRepositoryImpl{db: db}
}

func (r *attendanceRepositoryImpl) GetAll() ([]dao.Attendance, error) {
	var data []dao.Attendance
	err := r.db.
		Preload("Employee").
		Preload("Employee.Division").
		Find(&data).Error
	return data, err
}

func (r *attendanceRepositoryImpl) GetByID(id int) (dao.Attendance, error) {
	var data dao.Attendance
	err := r.db.
		Preload("Employee").
		Preload("Employee.Division").
		First(&data, id).Error
	return data, err
}

func (r *attendanceRepositoryImpl) Create(attendance *dao.Attendance) error {
	return r.db.Create(attendance).Error
}

func (r *attendanceRepositoryImpl) Update(attendance *dao.Attendance) error {
	return r.db.Save(attendance).Error
}

func (r *attendanceRepositoryImpl) Delete(id int) error {
	return r.db.Delete(&dao.Attendance{}, id).Error
}

func (r *attendanceRepositoryImpl) BulkUpsert(ctx context.Context, items []dao.Attendance) error {
	if len(items) == 0 {
		return nil
	}

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "employee_id"},
				{Name: "date"},
			},
			DoUpdates: clause.AssignmentColumns([]string{
				"home_time",
				"total_overtime",
				"note",
			}),
		}).Create(&items).Error
	})
}

func (r *attendanceRepositoryImpl) GetByDateAndDivision(date time.Time, divisionID int) ([]dao.Attendance, error) {
	var data []dao.Attendance

	err := r.db.
		Model(&dao.Attendance{}).
		Preload("Employee").
		Preload("Employee.Division").
		Joins("Employee").
		Where("attendance.date = ?", date).
		Where("Employee.division_id = ?", divisionID).
		Find(&data).Error

	return data, err
}
