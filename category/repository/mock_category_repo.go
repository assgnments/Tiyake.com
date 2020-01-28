package repository

import (
	"teyake/category"
	"teyake/entity"
)

type MockCategoryRepo struct {
	categories map[uint]*entity.Category
}


func NewMockCategoryRepo(categories map[uint]*entity.Category)  category.CategoryRepo{
	return  MockCategoryRepo{ categories}
}


func (c MockCategoryRepo) Catagories() ([]entity.Category, []error) {
	categories := []entity.Category{}
	for _, v := range c.categories {
		categories = append(categories, *v)
	}
	return categories, nil
}

func (c MockCategoryRepo) Category(id uint) (*entity.Category, []error) {
	return  nil,nil
}

func (c MockCategoryRepo) UpdateCategory(Category *entity.Category) (*entity.Category, []error) {
	return  nil,nil
}

func (c MockCategoryRepo) DeleteCategory(id uint) (*entity.Category, []error) {
	return  nil,nil
}

func (c MockCategoryRepo) StoreCategory(Category *entity.Category) (*entity.Category, []error) {
	return  nil,nil
}