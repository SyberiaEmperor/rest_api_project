package service

import (
	restwebprj "github.com/SyberiaEmperor/rest_web_prj"
	"github.com/SyberiaEmperor/rest_web_prj/pkg/repository"
)

type Authorization interface {
	CreateUser(user restwebprj.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
