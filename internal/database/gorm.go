package database

import (
	"github.com/seinkytarlicht/todious/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("todious.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Exec("PRAGMA foreign_keys = ON")

	db.AutoMigrate(&model.TodoGroup{}, &model.Todo{})

	return db
}
