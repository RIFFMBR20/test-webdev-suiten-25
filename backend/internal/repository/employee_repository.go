package repository

import (
	"test-webdev-suiten-25/internal/models/dao"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	GetAll() ([]dao.Employee, error)
	GetByID(id int) (dao.Employee, error)
	Create(employee *dao.Employee) error
	Update(employee *dao.Employee) error
	Delete(id int) error
}

type employeeRepositoryImpl struct {
	db *gorm.DB
}

func ProvideEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepositoryImpl{db: db}
}

func (r *employeeRepositoryImpl) GetAll() ([]dao.Employee, error) {
	var data []dao.Employee
	err := r.db.
		Preload("Division").
		Where("is_deleted = ?", false).
		Find(&data).Error
	return data, err
}

func (r *employeeRepositoryImpl) GetByID(id int) (dao.Employee, error) {
	var data dao.Employee
	err := r.db.
		Preload("Division").
		Where("id = ? AND is_deleted = ?", id, false).
		First(&data).Error
	return data, err
}

func (r *employeeRepositoryImpl) Create(employee *dao.Employee) error {
	return r.db.Create(employee).Error
}

func (r *employeeRepositoryImpl) Update(employee *dao.Employee) error {
	return r.db.Save(employee).Error
}

func (r *employeeRepositoryImpl) Delete(id int) error {
	return r.db.Model(&dao.Employee{}).
		Where("id = ?", id).
		Update("is_deleted", true).Error
}
