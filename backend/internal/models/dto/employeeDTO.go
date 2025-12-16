package dto

type EmployeeDTO struct {
	ID               int         `json:"id"`
	Name             string      `json:"name"`
	Division         DivisionDTO `json:"division"`
	PhoneNumber      string      `json:"phone_number"`
	AccountNumber    string      `json:"account_number"`
	BankName         string      `json:"bank_name"`
	Shift            string      `json:"shift"`
	Salary           int         `json:"salary"`
	PeriodSalary     string      `json:"period_salary"`
	DailySalary      int         `json:"daily_salary"`
	MealAllowance    int         `json:"meal_allowance"`
	RedMealAllowance int         `json:"red_meal_allowance"`
	Overtime         int         `json:"overtime"`
	RedOvertime      int         `json:"red_overtime"`
}

type EmployeeInputModifyDTO struct {
	DivisionID       int    `json:"division_id" validate:"required"`
	Name             string `json:"name" validate:"required"`
	PhoneNumber      string `json:"phone_number"`
	AccountNumber    string `json:"account_number"`
	BankName         string `json:"bank_name"`
	Shift            string `json:"shift"`
	Salary           int    `json:"salary" validate:"required"`
	PeriodSalary     string `json:"period_salary"`
	DailySalary      int    `json:"daily_salary"`
	MealAllowance    int    `json:"meal_allowance"`
	RedMealAllowance int    `json:"red_meal_allowance"`
	Overtime         int    `json:"overtime"`
	RedOvertime      int    `json:"red_overtime"`
}
