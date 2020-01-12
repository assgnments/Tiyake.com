package user

import "teyake/entity"

type UserRepository interface {
	User(id uint) (*entity.User,[]error)
	StoreUser(user *entity.User) (*entity.User,[]error)
	UserByEmail(email string) (*entity.User, [] error)

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
	DeleteSession(sessionId string ) (*entity.Session, []error)
}


