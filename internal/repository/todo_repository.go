package repository

import (
	"context"

	"github.com/seinkytarlicht/todious/internal/model"
	"gorm.io/gorm"
)

type TodoRepository interface {
	GetAllTodoGroup() []model.TodoGroup
	CreateTodoGroup(title string) model.TodoGroup
	UpdateTodoGroup(groupId uint, title string)
	DeleteTodoGroup(groupId uint)
	AddTodo(groupId uint, task string) model.Todo
	UpdateTodo(todoId uint, task string, finished bool)
	RemoveTodo(todoId uint)
}

type TodoRepositoryImpl struct {
	db  *gorm.DB
	ctx context.Context
}

func (t *TodoRepositoryImpl) GetAllTodoGroup() []model.TodoGroup {
	todoG, err := gorm.G[model.TodoGroup](t.db).Preload("Todos", nil).Find(t.ctx)

	if err != nil {
		panic(err)
	}

	return todoG
}

func (t *TodoRepositoryImpl) CreateTodoGroup(title string) model.TodoGroup {
	todoG := model.TodoGroup{
		Title: title,
	}

	err := gorm.G[model.TodoGroup](t.db).Create(t.ctx, &todoG)

	if err != nil {
		panic(err)
	}

	return todoG
}

func (t *TodoRepositoryImpl) UpdateTodoGroup(groupId uint, title string) {
	_, err := gorm.G[model.TodoGroup](t.db).Where("id = ?", groupId).Update(t.ctx, "Title", title)

	if err != nil {
		panic(err)
	}
}

func (t *TodoRepositoryImpl) DeleteTodoGroup(groupId uint) {
	_, err := gorm.G[model.TodoGroup](t.db).Where("id = ?", groupId).Delete(t.ctx)
	if err != nil {
		panic(err)
	}
}

func (t *TodoRepositoryImpl) AddTodo(groupId uint, task string) model.Todo {
	todo := model.Todo{
		TodoGroupID: groupId,
		Task:        task,
	}

	err := gorm.G[model.Todo](t.db).Create(t.ctx, &todo)

	if err != nil {
		panic(err)
	}

	return todo
}

func (t *TodoRepositoryImpl) UpdateTodo(todoId uint, task string, finished bool) {
	_, err := gorm.G[model.Todo](t.db).Where("id = ?", todoId).Updates(t.ctx, model.Todo{
		Task:     task,
		Finished: finished,
	})

	if err != nil {
		panic(err)
	}
}

func (t *TodoRepositoryImpl) RemoveTodo(todoId uint) {
	_, err := gorm.G[model.Todo](t.db).Where("id = ?", todoId).Delete(t.ctx)
	if err != nil {
		panic(err)
	}
}

func NewTodoRepository(db *gorm.DB, ctx context.Context) TodoRepository {
	return &TodoRepositoryImpl{
		db:  db,
		ctx: ctx,
	}
}
