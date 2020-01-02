package user

import "teyake/entity"

type UserRepository interface {
	//user by id
	User(id uint) (*entity.User,[]error)
	//user create
	StoreUser(user *entity.User) (*entity.User,[]error)
	UserByEmail(email string) (*entity.User,[]error)

}
type RoleRepository interface {
	//role by id
	Role(id uint) (*entity.Role,[]error)
	//role create
	StoreRole(role *entity.Role) (*entity.Role,[]error)

}

