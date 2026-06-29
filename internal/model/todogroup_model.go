package model

import "gorm.io/gorm"

type TodoGroup struct {
	gorm.Model

	Title string
	Todos []Todo
}
