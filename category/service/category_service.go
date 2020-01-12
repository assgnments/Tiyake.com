package service

import (
	"teyake/category"
	"teyake/entity"
)

type CategoryService struct {
	categoryRepo category.CategoryRepo
}

func NewCategoryService(categoryRepo category.CategoryRepo) category.CategoryRepo {
	return CategoryService{categoryRepo: categoryRepo}
}
func (c CategoryService) Catagories() ([]entity.Category, []error) {
	return c.categoryRepo.Catagories()
}

func (c CategoryService) Category(id uint) (*entity.Category, []error) {
	return c.categoryRepo.Category(id)
}

func (c CategoryService) UpdateCategory(Category *entity.Category) (*entity.Category, []error) {
	return c.categoryRepo.UpdateCategory(Category)
}

func (c CategoryService) DeleteCategory(id uint) (*entity.Category, []error) {
	return c.categoryRepo.DeleteCategory(id)
}

func (c CategoryService) StoreCategory(Category *entity.Category) (*entity.Category, []error) {
	return  c.categoryRepo.StoreCategory(Category)
}
