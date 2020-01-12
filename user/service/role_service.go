package service

import (
	"teyake/entity"
	"teyake/user"
)

// RoleService implements menu.RoleService interface
type RoleService struct {
	roleRepo user.RoleRepository
}
func NewRoleService(RoleRepo user.RoleRepository) user.RoleService {
	return &RoleService{roleRepo: RoleRepo}
}

func (rs *RoleService) Roles() ([]entity.Role, []error) {
	return rs.roleRepo.Roles()
}

func (rs *RoleService) RoleByName(name string) (*entity.Role, []error) {
	return rs.roleRepo.RoleByName(name)
}

func (rs *RoleService) Role(id uint) (*entity.Role, []error) {
	return rs.roleRepo.Role(id)
}

func (rs *RoleService) UpdateRole(role *entity.Role) (*entity.Role, []error) {
	return rs.roleRepo.UpdateRole(role)
}

func (rs *RoleService) DeleteRole(id uint) (*entity.Role, []error) {
	return rs.roleRepo.DeleteRole(id)
}

func (rs *RoleService) StoreRole(role *entity.Role) (*entity.Role, []error) {
	return rs.roleRepo.StoreRole(role)
}
