package repository

import (
	restwebprj "github.com/SyberiaEmperor/rest_web_prj"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user restwebprj.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}