package dto

type DivisionDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required"`
}
