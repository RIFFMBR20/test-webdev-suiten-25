package repository

import (
	"test-webdev-suiten-25/internal/models/dao"

	"gorm.io/gorm"
)

type DivisionRepository interface {
	GetAll() ([]dao.Division, error)
	GetByID(id int) (dao.Division, error)
	Create(division *dao.Division) error
	Update(division *dao.Division) error
	Delete(id int) error
}

type divisionRepositoryImpl struct {
	db *gorm.DB
}

func ProvideDivisionRepository(db *gorm.DB) DivisionRepository {
	return &divisionRepositoryImpl{db: db}
}

func (r *divisionRepositoryImpl) GetAll() ([]dao.Division, error) {
	var data []dao.Division
	err := r.db.Where("is_deleted = ?", false).Find(&data).Error
	return data, err
}

func (r *divisionRepositoryImpl) GetByID(id int) (dao.Division, error) {
	var data dao.Division
	err := r.db.
		Where("id = ? AND is_deleted = ?", id, false).
		First(&data).Error
	return data, err
}

func (r *divisionRepositoryImpl) Create(division *dao.Division) error {
	return r.db.Create(division).Error
}

func (r *divisionRepositoryImpl) Update(division *dao.Division) error {
	return r.db.Save(division).Error
}

func (r *divisionRepositoryImpl) Delete(id int) error {
	return r.db.Model(&dao.Division{}).
		Where("id = ?", id).
		Update("is_deleted", true).Error
}
