package user

import "teyake/entity"

type UserService interface {
	User(id uint) (*entity.User,[]error)
	StoreUser(user *entity.User) (*entity.User,[]error)
	UserByEmail(email string) (*entity.User, [] error)
	EmailExists(email string) bool

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
	Session(id uint) (*entity.Session, []error)
	Sessions() ([]entity.Session, []error)
	StoreSession(session *entity.Session) (*entity.Session, []error)
	DeleteSession(id uint) (*entity.Session, []error)
}
