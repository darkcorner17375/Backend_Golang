package service

import (
	"fmt"
	"log"

	"github.com/ed/gaintime/dto"
	"github.com/ed/gaintime/entity"
	"github.com/ed/gaintime/repository.go"
	"github.com/mashingan/smapping"
)

type TodoService interface {
	Insert(t dto.TodoCreateDTO) entity.Todo
	Update(t dto.TodoUpdateDTO) entity.Todo
	Delete(t entity.Todo)
	All() []entity.Todo
	FindByID(todo uint64) entity.Todo
	IsAllowedToEdit(userID string, todoID uint64) bool
}

type todoService struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return &todoService{
		todoRepository: todoRepo,
	}
}

func (service *todoService) Insert(t dto.TodoCreateDTO) entity.Todo {
	todo := entity.Todo{}
	err := smapping.FillStruct(&todo, smapping.MapFields(&t))
	if err != nil {
		log.Fatalf("Failed mapping %v", err)
	}
	res := service.todoRepository.InsertTodo(todo)
	return res
}

func (service *todoService) Update(t dto.TodoUpdateDTO) entity.Todo {
	todo := entity.Todo{}
	err := smapping.FillStruct(&todo, smapping.MapFields(&t))
	if err != nil {
		log.Fatalf("Failed mapping %v", err)
	}
	res := service.todoRepository.UpdateTodo(todo)
	return res
}

func (service *todoService) Delete(t entity.Todo) {
	service.todoRepository.DeleteTodo(t)
}

func (service *todoService) All() []entity.Todo {
	return service.todoRepository.AllTodo()
}

func (service *todoService) FindByID(todoID uint64) entity.Todo {
	return service.todoRepository.FindTodoByID(todoID)
}

func (service *todoService) IsAllowedToEdit(userID string, todoID uint64) bool {
	t := service.todoRepository.FindTodoByID(todoID)
	id := fmt.Sprintf("%v", t.UserID)
	return userID == id
}
