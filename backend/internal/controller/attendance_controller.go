package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"test-webdev-suiten-25/internal/models/dto"
	"test-webdev-suiten-25/internal/service"
	"test-webdev-suiten-25/internal/util"
	"time"
)

type AttendanceController interface {
	GetByDateAndDivision(w http.ResponseWriter, r *http.Request)
	BulkUpsert(w http.ResponseWriter, r *http.Request)
}

type attendanceControllerImpl struct {
	service service.AttendanceService
	log     *log.Logger
}

func ProvideAttendanceController(
	service service.AttendanceService,
	logger *log.Logger,
) AttendanceController {
	return &attendanceControllerImpl{
		service: service,
		log:     logger,
	}
}

// GetByDateAndDivision godoc
// @Summary      List attendance by date and division
// @Description  Get attendance list filtered by date (YYYY-MM-DD) and division_id
// @Tags         attendance
// @Produce      json
// @Param        date         query     string  true  "Date (YYYY-MM-DD)" example(2025-10-03)
// @Param        division_id  query     int     true  "Division ID" example(1)
// @Success      200  {array}   dto.AttendanceDTO
// @Failure      400  {object}  util.ErrorResponse
// @Failure      500  {object}  util.ErrorResponse
// @Router       /api/attendance [get]
func (c *attendanceControllerImpl) GetByDateAndDivision(w http.ResponseWriter, r *http.Request) {
	dateStr := strings.TrimSpace(r.URL.Query().Get("date"))
	divisionStr := strings.TrimSpace(r.URL.Query().Get("division_id"))

	if dateStr == "" {
		util.Error(w, http.StatusBadRequest, "date is required (YYYY-MM-DD)")
		return
	}
	if divisionStr == "" {
		util.Error(w, http.StatusBadRequest, "division_id is required")
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		util.Error(w, http.StatusBadRequest, "invalid date format (use YYYY-MM-DD)")
		return
	}

	divisionID, err := strconv.Atoi(divisionStr)
	if err != nil || divisionID <= 0 {
		util.Error(w, http.StatusBadRequest, "invalid division_id")
		return
	}

	data, err := c.service.GetByDateAndDivision(date, divisionID)
	if err != nil {
		c.log.Printf("error get attendance by date+division: %v\n", err)
		util.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.JSON(w, http.StatusOK, data)
}

// BulkUpsert godoc
// @Summary      Bulk upsert attendance
// @Description  Save multiple attendance rows at once (create or update by employee_id + date)
// @Tags         attendance
// @Accept       json
// @Produce      json
// @Param        request  body      dto.BulkAttendanceUpsertDTO  true  "Bulk attendance payload"
// @Success      200      {object}  BulkUpsertResponse
// @Failure      400      {object}  util.ErrorResponse
// @Failure      500      {object}  util.ErrorResponse
// @Router       /api/attendance/bulk [post]
func (c *attendanceControllerImpl) BulkUpsert(w http.ResponseWriter, r *http.Request) {
	var req dto.BulkAttendanceUpsertDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if len(req.Items) == 0 {
		util.Error(w, http.StatusBadRequest, "items is required")
		return
	}

	if err := c.service.BulkUpsert(r.Context(), req); err != nil {
		c.log.Printf("error bulk upsert attendance: %v\n", err)
		util.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	util.JSON(w, http.StatusOK, map[string]any{
		"updated": len(req.Items),
	})
}

type BulkUpsertResponse struct {
	Updated int `json:"updated"`
}
