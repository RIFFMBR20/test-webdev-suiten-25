package dao

import "gorm.io/gorm"

type Employee struct {
	ID int `gorm:"primaryKey"`

	DivisionID int
	Division   Division `gorm:"foreignKey:DivisionID"`

	Name             string `gorm:"type:varchar(100);not null"`
	PhoneNumber      string `gorm:"type:varchar(20)"`
	AccountNumber    string `gorm:"type:varchar(50)"`
	BankName         string `gorm:"type:varchar(50)"`
	Shift            string `gorm:"type:varchar(20)"`
	Salary           int
	PeriodSalary     string `gorm:"type:varchar(20)"`
	DailySalary      int    `gorm:"default:0"`
	MealAllowance    int
	RedMealAllowance int
	Overtime         int
	RedOvertime      int
	IsDeleted        bool
	gorm.Model
}

func (Employee) TableName() string {
	return "employee"
}
