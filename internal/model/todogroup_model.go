package model

type TodoGroup struct {
	ID    uint `gorm:"primarykey"`
	Title string
	Todos []Todo `gorm:"constraint:OnDelete:CASCADE;"`
}
