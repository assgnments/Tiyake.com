package user

import "teyake/entity"

type UserService interface {
	Users() ([]entity.User, []error)
	User(id uint) (*entity.User, []error)
	DeleteUser(id uint) (*entity.User, []error)
	UpdateUser(user *entity.User) (*entity.User, []error)
	StoreUser(user *entity.User) (*entity.User, []error)
	UserByEmail(email string) (*entity.User, []error)
	EmailExists(email string) bool
	UserRoles(user *entity.User) ([]entity.Role, []error)
}
type RoleService interface {
	Roles() ([]entity.Role, []error)
	Role(id uint) (*entity.Role, []error)
	RoleByName(name string) (*entity.Role, []error)
	UpdateRole(role *entity.Role) (*entity.Role, []error)
	DeleteRole(id uint) (*entity.Role, []error)
	StoreRole(role *entity.Role) (*entity.Role, []error)
}

// SessionService specifies logged in user session related service
type SessionService interface {
	Session(sessionId string) (*entity.Session, []error)
	Sessions() ([]entity.Session, []error)
	StoreSession(session *entity.Session) (*entity.Session, []error)
	DeleteSession(sessionId string) (*entity.Session, []error)
}
