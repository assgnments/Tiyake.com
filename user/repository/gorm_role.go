package repository

import (
	"github.com/jinzhu/gorm"
	"teyake/entity"
	"teyake/user"
)

type RoleGormRepo struct {
	conn *gorm.DB
}

func NewRoleGormRepo(db *gorm.DB)  user.RoleRepository {
	return &RoleGormRepo{conn: db,}
}

func (roleRepo *RoleGormRepo) Role(id uint)  (*entity.Role,[]error){
	role :=entity.Role{}
	errs:=roleRepo.conn.First(&role,id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &role, errs

}

func (roleRepo *RoleGormRepo) StoreRole(role *entity.Role) (*entity.Role,[]error){
	errs:=roleRepo.conn.Create(&role).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return role, errs
}