package repository

import (
	"errors"
	"teyake/entity"
	"teyake/user"
)

type MockRoleRepo struct {
	roles map[string]*entity.Role
}


func NewMockRoleRepo(roles map[string]*entity.Role)  user.RoleRepository {
return  &MockRoleRepo{roles}
}

func (roleRepo *MockRoleRepo) Role(id uint)  (*entity.Role,[]error){
	return  nil,nil
}

func (roleRepo *MockRoleRepo) Roles() ([]entity.Role, []error) {
	return  nil,nil
}

func (roleRepo *MockRoleRepo) RoleByName(name string) (*entity.Role, []error) {
	role := roleRepo.roles[name]
	if role == nil {
		return nil, []error{
			errors.New("Role not found"),
		}
	}
	return role, nil
}

func (roleRepo *MockRoleRepo) StoreRole(role *entity.Role) (*entity.Role,[]error){
	return  nil,nil
}

func (roleRepo *MockRoleRepo) UpdateRole(role *entity.Role) (*entity.Role, []error) {
	return  nil,nil
}

func (roleRepo *MockRoleRepo) DeleteRole(id uint) (*entity.Role, []error) {
	return  nil,nil
}

