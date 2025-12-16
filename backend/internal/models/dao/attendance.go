package dao

import "time"

type Attendance struct {
	ID            int `gorm:"primaryKey"`
	EmployeeID    int
	Employee      Employee
	Date          time.Time `gorm:"type:date"`
	HomeTime      string    `gorm:"type:time"`
	TotalOvertime string    `gorm:"type:text"`
	Note          string    `gorm:"type:text"`
}

func (Attendance) TableName() string {
	return "attendance"
}
