package dto

import "time"

type AttendanceDTO struct {
	ID            int       `json:"id"`
	EmployeeID    int       `json:"employee_id"`
	EmployeeName  string    `json:"employee_name"`
	DivisionID    int       `json:"division_id"`
	DivisionName  string    `json:"division_name"`
	Date          time.Time `json:"date"`
	HomeTime      string    `json:"home_time"`
	TotalOvertime string    `json:"total_overtime"`
	Note          string    `json:"note"`
}

type AttendanceUpsertDTO struct {
	EmployeeID int       `json:"employee_id"`
	Date       time.Time `json:"date"`
	HomeTime   string    `json:"home_time"`
	Note       string    `json:"note"`
}

type BulkAttendanceUpsertDTO struct {
	Items []AttendanceUpsertDTO `json:"items"`
}
