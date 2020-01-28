package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"teyake/entity"
	"teyake/user"
)

type MockUserRepo struct {
	users map[uint]*entity.User
}

func NewMockUserRepo() user.UserRepository {
	return &MockUserRepo{
		map[uint]*entity.User{
			0: {
				Model: gorm.Model{
					ID: 0,
				},
				FullName: "Nathan",
				RoleID:   0,
			},
			1: {
				Model: gorm.Model{
					ID: 1,
				},
				FullName: "Habib",
				RoleID:   1,
			},
		},
	}
}

func (userRepo *MockUserRepo) User(id uint) (*entity.User, []error) {
	user := userRepo.users[id]
	if user == nil {
		return nil, []error{
			errors.New("User not found"),
		}
	}
	return user, nil

}

func (userRepo *MockUserRepo) Users() ([]entity.User, []error) {
		return nil,nil
}

func (userRepo *MockUserRepo) DeleteUser(id uint) (*entity.User, []error) {
	return nil,nil
}

func (userRepo *MockUserRepo) StoreUser(user *entity.User) (*entity.User, []error) {
	return nil,nil
}

func (userRepo *MockUserRepo) UpdateUser(user *entity.User) (*entity.User, []error) {
	return nil,nil
}

func (userRepo *MockUserRepo) UserByEmail(email string) (*entity.User, []error) {
	return nil,nil
}

// UserRoles returns list of application roles that a given user has
func (userRepo *MockUserRepo) UserRoles(user *entity.User) ([]entity.Role, []error) {
	return nil,nil
}
