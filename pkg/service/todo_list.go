package service

import (
	restwebprj "github.com/SyberiaEmperor/rest_web_prj"
	"github.com/SyberiaEmperor/rest_web_prj/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list restwebprj.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]restwebprj.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int) (restwebprj.TodoList, error) {
	return s.repo.GetById(userId, listId)
}