package service

import (
	"teyake/entity"
	"teyake/user"
)

// SessionServiceImpl implements user.SessionService interface
type SessionServiceImpl struct {
	sessionRepo user.SessionRepository
}


// NewSessionService  returns a new SessionService object
func NewSessionService(sessRepository user.SessionRepository) user.SessionService {
	return &SessionServiceImpl{sessionRepo: sessRepository}
}

// Session returns a given stored session
func (ss *SessionServiceImpl) Session(sessionId string) (*entity.Session, []error) {
	return ss.sessionRepo.Session(sessionId)
}
// Returns all the sessions
func (ss *SessionServiceImpl) Sessions() ([]entity.Session, []error) {
	return  ss.sessionRepo.Sessions()
}

// StoreSession stores a given session
func (ss *SessionServiceImpl) StoreSession(session *entity.Session) (*entity.Session, []error) {
	return ss.sessionRepo.StoreSession(session)
}

// DeleteSession deletes a given session
func (ss *SessionServiceImpl) DeleteSession(sessionId string) (*entity.Session, []error) {
	return  ss.sessionRepo.DeleteSession(sessionId)
}
