package repository

import (
	"github.com/jinzhu/gorm"
	"teyake/entity"
	"teyake/user"
)

type UserGormRepo struct {
	conn *gorm.DB
}

func NewUserGormRepo(db *gorm.DB) user.UserRepository {
	return &UserGormRepo{conn: db}
}

func (userRepo *UserGormRepo) User(id uint) (*entity.User, []error) {
	user := entity.User{}
	errs := userRepo.conn.First(&user, id).GetErrors()
	return &user, errs

}

func (userRepo *UserGormRepo) Users() ([]entity.User, []error) {
	users := []entity.User{}
	errs := userRepo.conn.Find(&users).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return users, errs
}

func (userRepo *UserGormRepo) DeleteUser(id uint) (*entity.User, []error) {
	user, errs := userRepo.User(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = userRepo.conn.Delete(user, user.ID).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return user, errs
}

func (userRepo *UserGormRepo) StoreUser(user *entity.User) (*entity.User, []error) {
	errs := userRepo.conn.Create(user).GetErrors()
	return user, errs
}

func (userRepo *UserGormRepo) UpdateUser(user *entity.User) (*entity.User, []error) {
	usr := user
	errs := userRepo.conn.Save(&usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

func (userRepo *UserGormRepo) UserByEmail(email string) (*entity.User, []error) {
	user := entity.User{}
	errs := userRepo.conn.First(&user, "email=?", email).GetErrors()
	return &user, errs
}

// UserRoles returns list of application roles that a given user has
func (userRepo *UserGormRepo) UserRoles(user *entity.User) ([]entity.Role, []error) {
	userRoles := []entity.Role{}
	errs := userRepo.conn.Model(user).Related(&userRoles).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return userRoles, errs
}
