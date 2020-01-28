package repository

import (
	"github.com/jinzhu/gorm"
	"teyake/category"
	"teyake/entity"
)

type MockCategoryRepo struct {
	categories map[uint]*entity.Category
}

func NewMockCategoryRepo()  category.CategoryRepo{
	return  MockCategoryRepo{ map[uint]*entity.Category{
		0: {
			Model: gorm.Model{
				ID:0,
			},
			Name:  "Science",
		},
		1: {
			Model: gorm.Model{
				ID:0,
			},
			Name:  "Programming",
		},
	}}
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