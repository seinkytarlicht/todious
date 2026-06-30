package service

import (
	"github.com/seinkytarlicht/todious/internal/model"
	"github.com/seinkytarlicht/todious/internal/repository"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type TodoService struct {
	app   *application.App
	todoR repository.TodoRepository
}

func (t *TodoService) GetAllTodoGroup() []model.TodoGroup {
	return t.todoR.GetAllTodoGroup()
}

func (t *TodoService) CreateTodoGroup(title string) model.TodoGroup {
	todoG := t.todoR.CreateTodoGroup(title)

	t.app.Event.Emit("todoService:changes")
	return todoG
}

func (t *TodoService) UpdateTodoGroup(groupId uint, title string) {
	t.todoR.UpdateTodoGroup(groupId, title)

	t.app.Event.Emit("todoService:changes")
}

func (t *TodoService) DeleteTodoGroup(groupId uint) {
	t.todoR.DeleteTodoGroup(groupId)

	t.app.Event.Emit("todoService:changes")
}

func (t *TodoService) AddTodo(groupId uint, task string) model.Todo {
	todo := t.todoR.AddTodo(groupId, task)

	t.app.Event.Emit("todoService:taskCreated")
	return todo
}

func (t *TodoService) UpdateTodo(todoId uint, task string, finished bool) {
	t.todoR.UpdateTodo(todoId, task, finished)

	t.app.Event.Emit("todoService:taskUpdated")
}

func (t *TodoService) RemoveTodo(todoId uint) {
	t.todoR.RemoveTodo(todoId)

	t.app.Event.Emit("todoService:taskRemoved")
}

func NewTodoService(todoRepo repository.TodoRepository, app *application.App) *TodoService {
	return &TodoService{
		todoR: todoRepo,
		app:   app,
	}
}
