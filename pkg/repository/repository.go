package repository

import (
	restwebprj "github.com/SyberiaEmperor/rest_web_prj"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user restwebprj.User) (int, error)
	GetUser(username, password string) (restwebprj.User, error)
}

type TodoList interface {
	Create(userId int, list restwebprj.TodoList) (int, error)
	GetAll(userId int) ([]restwebprj.TodoList, error)
	GetById(userId, listId int) (restwebprj.TodoList, error)
}

type TodoItem interface {
	Create(listId int, list restwebprj.TodoItem) (int, error)
	GetAll(userID, listId int) ([]restwebprj.TodoItem, error)
	GetById(userId, itemId int) (restwebprj.TodoItem, error)
	Update(userId, itemId int, input restwebprj.UpdateItemInput) error
	Delete(userId, itemId int) error 
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList: NewTodoListPostgres(db),
		TodoItem: NewTodoItemPostgres(db),
	}
}