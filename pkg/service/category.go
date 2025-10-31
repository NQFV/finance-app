package service

import (
	"github.com/NQFV/p/pkg/models"
	"github.com/NQFV/p/pkg/repository"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) Create(Id int, category models.Category) (int, error) {
	return s.repo.Create(Id, category)
}

func (s *CategoryService) GetAll(userId int) ([]models.Category, error) {
	return s.repo.GetAll(userId)
}

func (s *CategoryService) Update(userId, categoryId int, input models.Category) error {
	return s.repo.Update(userId, categoryId, input)
}

func (s *CategoryService) Delete(userId, categoryId int) error {
	return s.repo.Delete(userId, categoryId)
}
