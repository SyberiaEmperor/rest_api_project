package service

import (
	restwebprj "github.com/SyberiaEmperor/rest_web_prj"
	"github.com/SyberiaEmperor/rest_web_prj/pkg/repository"
)

type TodoItemService struct {
	repo repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item restwebprj.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)

	if err != nil {
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]restwebprj.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId, itemId int) (restwebprj.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *TodoItemService) Update(userId, itemId int, input restwebprj.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}