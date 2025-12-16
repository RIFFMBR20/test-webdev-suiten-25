package mapper

import (
	"test-webdev-suiten-25/internal/models/dao"
	"test-webdev-suiten-25/internal/models/dto"
)

func DivisionToDTO(d dao.Division) dto.DivisionDTO {
	return dto.DivisionDTO{
		ID:   d.ID,
		Name: d.Name,
	}
}

func EmployeeToDTO(d dao.Employee) dto.EmployeeDTO {
	return dto.EmployeeDTO{
		ID:               d.ID,
		Division:         DivisionToDTO(d.Division),
		Name:             d.Name,
		PhoneNumber:      d.PhoneNumber,
		AccountNumber:    d.AccountNumber,
		BankName:         d.BankName,
		Shift:            d.Shift,
		Salary:           d.Salary,
		PeriodSalary:     d.PeriodSalary,
		DailySalary:      d.DailySalary,
		MealAllowance:    d.MealAllowance,
		RedMealAllowance: d.RedMealAllowance,
		Overtime:         d.Overtime,
		RedOvertime:      d.RedOvertime,
	}
}

func EmployeeInputToDao(input dto.EmployeeInputModifyDTO) dao.Employee {
	return dao.Employee{
		DivisionID:       input.DivisionID,
		Name:             input.Name,
		PhoneNumber:      input.PhoneNumber,
		AccountNumber:    input.AccountNumber,
		BankName:         input.BankName,
		Shift:            input.Shift,
		Salary:           input.Salary,
		PeriodSalary:     input.PeriodSalary,
		DailySalary:      input.DailySalary,
		MealAllowance:    input.MealAllowance,
		RedMealAllowance: input.RedMealAllowance,
		Overtime:         input.Overtime,
		RedOvertime:      input.RedOvertime,
	}
}

func UpdateEmployeeFromInputDTO(existing dao.Employee, input dto.EmployeeInputModifyDTO) dao.Employee {
	existing.DivisionID = input.DivisionID
	existing.Name = input.Name
	existing.PhoneNumber = input.PhoneNumber
	existing.AccountNumber = input.AccountNumber
	existing.BankName = input.BankName
	existing.Shift = input.Shift
	existing.Salary = input.Salary
	existing.PeriodSalary = input.PeriodSalary
	existing.DailySalary = input.DailySalary
	existing.MealAllowance = input.MealAllowance
	existing.RedMealAllowance = input.RedMealAllowance
	existing.Overtime = input.Overtime
	existing.RedOvertime = input.RedOvertime
	return existing
}
