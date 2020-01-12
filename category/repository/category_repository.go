package repository

import (
	"github.com/jinzhu/gorm"
	"teyake/category"
	"teyake/entity"
)

type CategoryGormRepo struct {
	conn *gorm.DB
}

func NewCategoryGormRepo(conn *gorm.DB)  category.CategoryRepo{
	return  CategoryGormRepo{conn:conn}
}


func (c CategoryGormRepo) Catagories() ([]entity.Category, []error) {
	categories:=[]entity.Category{}
	errs:= c.conn.Find(&categories).GetErrors()
	return categories,errs
}

func (c CategoryGormRepo) Category(id uint) (*entity.Category, []error) {
	category:=entity.Category{}
	errs:=c.conn.Find(&category,id).GetErrors()
	return &category,errs
}

func (c CategoryGormRepo) UpdateCategory(Category *entity.Category) (*entity.Category, []error) {
	errs := c.conn.Save(Category).GetErrors()
	return Category, errs
}

func (c CategoryGormRepo) DeleteCategory(id uint) (*entity.Category, []error) {
	category, errs := c.Category(id)

	if len(errs) > 0 {
		return nil, errs
	}
	errs = c.conn.Delete(category, id).GetErrors()
	return category, errs
}

func (c CategoryGormRepo) StoreCategory(Category *entity.Category) (*entity.Category, []error) {
	errs := c.conn.Create(Category).GetErrors()
	return Category, errs
}