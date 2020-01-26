package repository

import (
	"github.com/jinzhu/gorm"
	"teyake/entity"
	"teyake/user"
)

// SessionGormRepo implements user.SessionRepository interface
type SessionGormRepo struct {
	conn *gorm.DB
}

// NewSessionGormRepo  returns a new SessionGormRepo object
func NewSessionGormRepo(db *gorm.DB) user.SessionRepository {
	return &SessionGormRepo{conn: db}
}

// Session returns a given stored session
func (sr *SessionGormRepo) Session(sessionId string) (*entity.Session, []error) {
	session := entity.Session{}
	errs := sr.conn.Find(&session, "session_id=?", sessionId).GetErrors()
	return &session, errs
}

// Returns all the sessions
func (sr *SessionGormRepo) Sessions() ([]entity.Session, []error) {
	sessions := []entity.Session{}
	errs := sr.conn.Find(&sessions).GetErrors()
	return sessions, errs
}

// StoreSession stores a given session
func (sr *SessionGormRepo) StoreSession(session *entity.Session) (*entity.Session, []error) {
	sess := session
	errs := sr.conn.Save(sess).GetErrors()
	return sess, errs
}

// DeleteSession deletes a given session
func (sr *SessionGormRepo) DeleteSession(sessionId string) (*entity.Session, []error) {
	sess, errs := sr.Session(sessionId)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = sr.conn.Delete(sess, "session_id=?", sessionId).GetErrors()
	return sess, errs
}
