package dao

import "gorm.io/gorm"

type Division struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"size:255;unique"`
	IsDeleted bool
	gorm.Model
}

func (Division) TableName() string {
	return "division"
}
