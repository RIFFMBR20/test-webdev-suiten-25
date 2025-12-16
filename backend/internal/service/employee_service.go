package service

import (
	"errors"
	"log"
	"test-webdev-suiten-25/internal/models/dto"
	"test-webdev-suiten-25/internal/models/mapper"
	"test-webdev-suiten-25/internal/repository"
)

type EmployeeService interface {
	GetAll() ([]dto.EmployeeDTO, error)
	GetByID(id int) (dto.EmployeeDTO, error)
	Upsert(id int, employeeDTO dto.EmployeeInputModifyDTO) (dto.EmployeeDTO, error)
	Delete(id int) error
}

type employeeServiceImpl struct {
	repo         repository.EmployeeRepository
	divisionRepo repository.DivisionRepository
	log          *log.Logger
}

func ProvideEmployeeService(
	repo repository.EmployeeRepository,
	divisionRepo repository.DivisionRepository,
	logger *log.Logger,
) EmployeeService {
	return &employeeServiceImpl{
		repo:         repo,
		divisionRepo: divisionRepo,
		log:          logger,
	}
}

func (s *employeeServiceImpl) GetAll() ([]dto.EmployeeDTO, error) {
	s.log.Println("get all employees")

	employees, err := s.repo.GetAll()
	if err != nil {
		s.log.Println("error get all employees:", err)
		return nil, err
	}

	var result []dto.EmployeeDTO
	for _, e := range employees {
		result = append(result, mapper.EmployeeToDTO(e))
	}

	return result, nil
}

func (s *employeeServiceImpl) GetByID(id int) (dto.EmployeeDTO, error) {
	s.log.Printf("get employee by id: %d\n", id)

	employee, err := s.repo.GetByID(id)
	if err != nil {
		s.log.Printf("employee not found: %d\n", id)
		return dto.EmployeeDTO{}, errors.New("employee not found")
	}

	return mapper.EmployeeToDTO(employee), nil
}

func (s *employeeServiceImpl) Upsert(id int, employeeDTO dto.EmployeeInputModifyDTO) (dto.EmployeeDTO, error) {
	if employeeDTO.Name == "" || employeeDTO.DivisionID <= 0 {
		return dto.EmployeeDTO{}, errors.New("employee name and valid division ID are required")
	}

	division, err := s.divisionRepo.GetByID(employeeDTO.DivisionID)
	if err != nil {
		return dto.EmployeeDTO{}, errors.New("invalid DivisionID: division not found")
	}

	if id == 0 {
		s.log.Printf("Create employee: %s", employeeDTO.Name)

		employee := mapper.EmployeeInputToDao(employeeDTO)
		employee.Division = division

		err := s.repo.Create(&employee)
		if err != nil {
			s.log.Println("error create employee:", err)
			return dto.EmployeeDTO{}, err
		}
		employee.Division = division
		return mapper.EmployeeToDTO(employee), nil

	} else {
		s.log.Printf("Update employee id=%d", id)

		employee, err := s.repo.GetByID(id)
		if err != nil {
			s.log.Printf("employee not found for update: %d\n", id)
			return dto.EmployeeDTO{}, errors.New("employee not found")
		}

		employee = mapper.UpdateEmployeeFromInputDTO(employee, employeeDTO)
		employee.Division = division

		err = s.repo.Update(&employee)
		if err != nil {
			s.log.Println("error update employee:", err)
			return dto.EmployeeDTO{}, err
		}
		employee.Division = division
		return mapper.EmployeeToDTO(employee), nil
	}
}

func (s *employeeServiceImpl) Delete(id int) error {
	s.log.Printf("delete employee id=%d\n", id)

	_, err := s.repo.GetByID(id)
	if err != nil {
		s.log.Printf("employee not found for delete: %d\n", id)
		return errors.New("employee not found")
	}
	
	err = s.repo.Delete(id)
	if err != nil {
		s.log.Println("error soft delete employee:", err)
		return err
	}

	s.log.Printf("employee soft deleted id=%d\n", id)
	return nil
}
