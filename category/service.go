package category

import "teyake/entity"

type CategoryService interface {
	Catagories() ([]entity.Category, []error)
	Category(id uint) (*entity.Category, []error)
	UpdateCategory(Category *entity.Category) (*entity.Category, []error)
	DeleteCategory(id uint) (*entity.Category, []error)
	StoreCategory(Category *entity.Category) (*entity.Category, []error)
}
