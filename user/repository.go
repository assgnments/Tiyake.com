package user

import "teyake/entity"

type UserRepository interface {
	Users() ([]entity.User, []error)
	User(id uint) (*entity.User, []error)
	DeleteUser(id uint) (*entity.User, []error)
	UpdateUser(user *entity.User) (*entity.User, []error)
	StoreUser(user *entity.User) (*entity.User, []error)
	UserByEmail(email string) (*entity.User, []error)
	UserRoles(user *entity.User) ([]entity.Role, []error)
}
type RoleRepository interface {
	Roles() ([]entity.Role, []error)
	Role(id uint) (*entity.Role, []error)
	RoleByName(name string) (*entity.Role, []error)
	UpdateRole(role *entity.Role) (*entity.Role, []error)
	DeleteRole(id uint) (*entity.Role, []error)
	StoreRole(role *entity.Role) (*entity.Role, []error)
}
type SessionRepository interface {
	Session(sessionId string) (*entity.Session, []error)
	Sessions() ([]entity.Session, []error)
	StoreSession(session *entity.Session) (*entity.Session, []error)
	DeleteSession(sessionId string) (*entity.Session, []error)
}
