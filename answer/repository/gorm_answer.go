package repository

import (
	"teyake/answer"
	"teyake/entity"

	"github.com/jinzhu/gorm"
)

// AnswerGormRepo implements menu.AnswerRepository interface
type AnswerGormRepo struct {
	conn *gorm.DB
}

// NewAnswerGormRepo returns new object of AnswerGormRepo
func NewAnswerGormRepo(db *gorm.DB) answer.AnswerRepository {
	return &AnswerGormRepo{conn: db}
}

// Answers returns all customer Answers stored in the database
func (ansRepo *AnswerGormRepo) Answers() ([]entity.Answer, []error) {
	cmnts := []entity.Answer{}
	errs := ansRepo.conn.Find(&cmnts).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnts, errs
}

// Answer retrieves a customer Answer from the database by its id
func (ansRepo *AnswerGormRepo) Answer(id uint) (*entity.Answer, []error) {
	cmnt := entity.Answer{}
	errs := ansRepo.conn.Set("gorm:auto_preload", true).First(&cmnt, id).GetErrors()
	return &cmnt, errs
}

// UpdateAnswer updates a given customer Answer in the database
func (ansRepo *AnswerGormRepo) UpdateAnswer(answer *entity.Answer) (*entity.Answer, []error) {
	errs := ansRepo.conn.Save(answer).GetErrors()
	return answer, errs
}

// DeleteAnswer deletes a given customer Answer from the database
func (ansRepo *AnswerGormRepo) DeleteAnswer(id uint) (*entity.Answer, []error) {
	cmnt, errs := ansRepo.Answer(id)

	if len(errs) > 0 {
		return nil, errs
	}
	errs = ansRepo.conn.Delete(cmnt, id).GetErrors()
	return cmnt, errs
}

// StoreAnswer stores a given customer Answer in the database
func (ansRepo *AnswerGormRepo) StoreAnswer(answer *entity.Answer) (*entity.Answer, []error) {
	errs := ansRepo.conn.Create(answer).GetErrors()
	return answer, errs
}
