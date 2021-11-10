package repository

import (
	"github.com/ed/gaintime/entity"
	"gorm.io/gorm"
)

type TodoRepository interface {
	InsertTodo(t entity.Todo) entity.Todo
	UpdateTodo(t entity.Todo) entity.Todo
	DeleteTodo(t entity.Todo)
	AllTodo() []entity.Todo
	FindTodoByID(TodoID uint64) entity.Todo
}

type todoConnection struct {
	connection *gorm.DB
}

func NewTodoRepository(dbConn *gorm.DB) TodoRepository {
	return &todoConnection{
		connection: dbConn,
	}
}

func (db *todoConnection) InsertTodo(t entity.Todo) entity.Todo {
	db.connection.Save(&t)
	db.connection.Preload("User").Find(&t)
	return t
}

func (db *todoConnection) UpdateTodo(t entity.Todo) entity.Todo {
	db.connection.Save(&t)
	db.connection.Preload("User").Find(&t)
	return t
}

func (db *todoConnection) DeleteTodo(t entity.Todo) {
	db.connection.Delete(&t)
}

func (db *todoConnection) FindTodoByID(TodoID uint64) entity.Todo {
	var todo entity.Todo
	db.connection.Preload("User").Find(&todo, TodoID)
	return todo
}

func (db *todoConnection) AllTodo() []entity.Todo {
	var todos []entity.Todo
	db.connection.Preload("User").Find(&todos)
	return todos
}
