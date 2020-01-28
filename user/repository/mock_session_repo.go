package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strconv"
	"teyake/entity"
	"teyake/user"
)

// MockSessionRepo implements user.SessionRepository interface
type MockSessionRepo struct {
	sessions map[uint]*entity.Session
}

// NewMockSessionRepo  returns a new MockSessionRepo object
func NewMockSessionRepo() user.SessionRepository {
	return &MockSessionRepo{
		map[uint]*entity.Session{
			0: {
				Model: gorm.Model{
					ID: 0,
				},
				UUID:       0,
				SessionId:  "0",
				SigningKey: []byte("demo_key"),
			},
			1: {
				Model: gorm.Model{
					ID: 1,
				},
				UUID:       1,
				SessionId:  "1",
				SigningKey: []byte("demo_key"),
			},
		},
	}
}

// Session returns a given stored session
func (sr *MockSessionRepo) Session(sessionId string) (*entity.Session, []error) {
	id, err := strconv.Atoi(sessionId)
	if err != nil {
		return nil, []error{
			err,
		}
	}
	session := sr.sessions[uint(id)]
	if session == nil {
		return nil, []error{
			errors.New("session not found"),
		}
	}
	return session, nil
}

// Returns all the sessions
func (sr *MockSessionRepo) Sessions() ([]entity.Session, []error) {
	return nil, nil
}

// StoreSession stores a given session
func (sr *MockSessionRepo) StoreSession(session *entity.Session) (*entity.Session, []error) {
	return nil, nil
}

// DeleteSession deletes a given session
func (sr *MockSessionRepo) DeleteSession(sessionId string) (*entity.Session, []error) {
	return nil, nil
}
