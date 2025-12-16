package service

import (
	"errors"
	"log"
	"test-webdev-suiten-25/internal/models/dao"
	"test-webdev-suiten-25/internal/models/dto"
	"test-webdev-suiten-25/internal/models/mapper"
	"test-webdev-suiten-25/internal/repository"
)

type DivisionService interface {
	GetAll() ([]dto.DivisionDTO, error)
	GetByID(id int) (dto.DivisionDTO, error)
	Create(name string) (dto.DivisionDTO, error)
	Update(id int, name string) (dto.DivisionDTO, error)
	Delete(id int) error
}

type divisionServiceImpl struct {
	repo repository.DivisionRepository
	log  *log.Logger
}

func ProvideDivisionService(
	repo repository.DivisionRepository,
	logger *log.Logger,
) DivisionService {
	return &divisionServiceImpl{
		repo: repo,
		log:  logger,
	}
}

func (s *divisionServiceImpl) GetAll() ([]dto.DivisionDTO, error) {
	s.log.Println("get all divisions")

	divisions, err := s.repo.GetAll()
	if err != nil {
		s.log.Println("error get all divisions:", err)
		return nil, err
	}

	var result []dto.DivisionDTO
	for _, d := range divisions {
		result = append(result, mapper.DivisionToDTO(d))
	}

	return result, nil
}

func (s *divisionServiceImpl) GetByID(id int) (dto.DivisionDTO, error) {
	s.log.Printf("get division by id: %d\n", id)

	division, err := s.repo.GetByID(id)
	if err != nil {
		s.log.Printf("division not found: %d\n", id)
		return dto.DivisionDTO{}, errors.New("division not found")
	}

	return mapper.DivisionToDTO(division), nil
}

func (s *divisionServiceImpl) Create(name string) (dto.DivisionDTO, error) {
	s.log.Printf("create division: %s\n", name)

	if name == "" {
		return dto.DivisionDTO{}, errors.New("division name is required")
	}

	division := dao.Division{
		Name: name,
	}

	err := s.repo.Create(&division)
	if err != nil {
		s.log.Println("error create division:", err)
		return dto.DivisionDTO{}, err
	}

	s.log.Printf("division created with id: %d\n", division.ID)

	return mapper.DivisionToDTO(division), nil
}

func (s *divisionServiceImpl) Update(id int, name string) (dto.DivisionDTO, error) {
	s.log.Printf("update division id=%d\n", id)

	if name == "" {
		return dto.DivisionDTO{}, errors.New("division name is required")
	}

	division, err := s.repo.GetByID(id)
	if err != nil {
		s.log.Printf("division not found: %d\n", id)
		return dto.DivisionDTO{}, errors.New("division not found")
	}

	division.Name = name

	err = s.repo.Update(&division)
	if err != nil {
		s.log.Println("error update division:", err)
		return dto.DivisionDTO{}, err
	}

	return mapper.DivisionToDTO(division), nil
}

func (s *divisionServiceImpl) Delete(id int) error {
	s.log.Printf("delete division id=%d\n", id)

	_, err := s.repo.GetByID(id)
	if err != nil {
		s.log.Printf("division not found: %d\n", id)
		return errors.New("division not found")
	}

	err = s.repo.Delete(id)
	if err != nil {
		s.log.Println("error delete division:", err)
		return err
	}

	s.log.Printf("division deleted id=%d\n", id)
	return nil
}
