package service

import (
	restwebprj "github.com/SyberiaEmperor/rest_web_prj"
	"github.com/SyberiaEmperor/rest_web_prj/pkg/repository"
)

type Authorization interface {
	CreateUser(user restwebprj.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list restwebprj.TodoList) (int, error)
	GetAll(userId int) ([]restwebprj.TodoList, error)
	GetById(userId, listId int) (restwebprj.TodoList, error)
	//Delete(userId, listId int) error
	//Update(userId, listId int, input restwebprj.UpdateListInput)
}

type TodoItem interface {
	Create(userId, listId int, list restwebprj.TodoItem) (int, error)
	GetAll(userID, listId int) ([]restwebprj.TodoItem, error)
	GetById(userId, itemId int) (restwebprj.TodoItem, error)
	Update(userId, itemId int, input restwebprj.UpdateItemInput) error
	Delete(userId, itemId int) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList: NewTodoListService(repos.TodoList),
		TodoItem: NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
