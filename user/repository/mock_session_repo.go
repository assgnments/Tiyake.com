package repository

import (
	"errors"
	"strconv"
	"teyake/entity"
	"teyake/user"
)

// MockSessionRepo implements user.SessionRepository interface
type MockSessionRepo struct {
	sessions map[uint]*entity.Session
}

// NewMockSessionRepo  returns a new MockSessionRepo object
func NewMockSessionRepo( sessions map[uint]*entity.Session) user.SessionRepository {
	return &MockSessionRepo{
		sessions,
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
