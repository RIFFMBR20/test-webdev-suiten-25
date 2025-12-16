package dto

type AttendanceDTO struct {
	ID       int         `json:"id"`
	Name     EmployeeDTO `json:"name"`
	Date     string      `json:"date"`
	Division DivisionDTO `json:"division"`
	HomeTime string      `json:"home_time"`
	Note     string      `json:"note"`
}
