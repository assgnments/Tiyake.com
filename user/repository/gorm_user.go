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
	return &UserGormRepo{conn: db,}
}

func (userRepo *UserGormRepo) User(id uint)  (*entity.User,[]error){
	user :=entity.User{}
	errs:=userRepo.conn.First(&user,id).GetErrors()
	return &user, errs

}

func (userRepo *UserGormRepo) StoreUser(user *entity.User) (*entity.User,[]error){
	errs:=userRepo.conn.Create(user).GetErrors()
	return user, errs
}
func (userRepo *UserGormRepo) UserByEmail(email string) (*entity.User, []error) {
	user :=entity.User{}
	errs:=userRepo.conn.First(&user,"email=?",email).GetErrors()
	return &user, errs
}

