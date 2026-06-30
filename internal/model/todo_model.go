package model

type Todo struct {
	ID          uint `gorm:"primarykey"`
	Task        string
	Finished    bool
	TodoGroupID uint
}
