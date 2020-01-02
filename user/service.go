package user

import "teyake/entity"

type UserService interface {
	//user by id
	User(id uint) (*entity.User,[]error)
	//user create
	StoreUser(user *entity.User) (*entity.User,[]error)

	EmailExists(email string) bool
	UserByEmail(email string) (*entity.User, [] error)

}
type RoleService interface {
	//role by id
	Role(id uint) (*entity.Role,[]error)
	//role create
	StoreRole(role *entity.Role) (*entity.Role,[]error)

}