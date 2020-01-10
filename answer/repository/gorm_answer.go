package repository

import (
	"teyake/entity"
	"teyake/answer"
	"github.com/jinzhu/gorm"
)

// AnswerGormRepo implements menu.AnswerRepository interface
type AnswerGormRepo struct {
	conn *gorm.DB
}

// NewAnswerGormRepo returns new object of AnswerGormRepo
func NewAnswerGormRepo(db *gorm.DB) Answer.AnswerRepository {
	return &AnswerGormRepo{conn: db}
}

// Answers returns all customer Answers stored in the database
func (cmntRepo *AnswerGormRepo) Answers() ([]entity.Answer, []error) {
	cmnts := []entity.Answer{}
	errs := cmntRepo.conn.Find(&cmnts).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnts, errs
}

// Answer retrieves a customer Answer from the database by its id
func (cmntRepo *AnswerGormRepo) Answer(id uint) (*entity.Answer, []error) {
	cmnt := entity.Answer{}
	errs := cmntRepo.conn.First(&cmnt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &cmnt, errs
}

// UpdateAnswer updates a given customer Answer in the database
func (cmntRepo *AnswerGormRepo) UpdateAnswer(Answer *entity.Answer) (*entity.Answer, []error) {
	cmnt := Answer
	errs := cmntRepo.conn.Save(cmnt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

// DeleteAnswer deletes a given customer Answer from the database
func (cmntRepo *AnswerGormRepo) DeleteAnswer(id uint) (*entity.Answer, []error) {
	cmnt, errs := cmntRepo.Answer(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = cmntRepo.conn.Delete(cmnt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

// StoreAnswer stores a given customer Answer in the database
func (cmntRepo *AnswerGormRepo) StoreAnswer(Answer *entity.Answer) (*entity.Answer, []error) {
	cmnt := Answer
	errs := cmntRepo.conn.Create(cmnt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}
