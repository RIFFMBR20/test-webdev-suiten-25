package repository

import (
	"test-webdev-suiten-25/internal/models/dao"

	"gorm.io/gorm"
)

type AttendanceRepository interface {
	GetAll() ([]dao.Attendance, error)
	GetByID(id int) (dao.Attendance, error)
	Create(attendance *dao.Attendance) error
	Update(attendance *dao.Attendance) error
	Delete(id int) error
}

type attendanceRepositoryImpl struct {
	db *gorm.DB
}

func ProvideAttendanceRepository(db *gorm.DB) AttendanceRepository {
	return &attendanceRepositoryImpl{db: db}
}

func (r *attendanceRepositoryImpl) GetAll() ([]dao.Attendance, error) {
	var data []dao.Attendance
	err := r.db.Preload("Employee").Find(&data).Error
	return data, err
}

func (r *attendanceRepositoryImpl) GetByID(id int) (dao.Attendance, error) {
	var data dao.Attendance
	err := r.db.
		Preload("Employee").
		First(&data, id).Error
	return data, err
}

func (r *attendanceRepositoryImpl) Create(attendance *dao.Attendance) error {
	return r.db.Create(attendance).Error
}

func (r *attendanceRepositoryImpl) Update(attendance *dao.Attendance) error {
	return r.db.Save(attendance).Error
}

func (r *attendanceRepositoryImpl) Delete(id int) error {
	return r.db.Delete(&dao.Attendance{}, id).Error
}
