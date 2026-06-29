package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model

	Task        string
	Finished    bool
	TodoGroupID uint
}
