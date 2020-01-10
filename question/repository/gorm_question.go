package repository

import (
	"teyake/entity"
	"teyake/question"

	"github.com/jinzhu/gorm"
)

// QuestionGormRepo implements menu.QuestionRepository interface
type QuestionGormRepo struct {
	conn *gorm.DB
}

// NewQuestionGormRepo returns new object of QuestionGormRepo
func NewQuestionGormRepo(db *gorm.DB) question.QuestionRepository {
	return &QuestionGormRepo{conn: db}
}

// Questions returns all customer Questions stored in the database
func (cmntRepo *QuestionGormRepo) Questions() ([]entity.Question, []error) {
	cmnts := []entity.Question{}
	errs := cmntRepo.conn.Find(&cmnts).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnts, errs
}

// Question retrieves a customer Question from the database by its id
func (cmntRepo *QuestionGormRepo) Question(id uint) (*entity.Question, []error) {
	cmnt := entity.Question{}
	errs := cmntRepo.conn.First(&cmnt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &cmnt, errs
}

// UpdateQuestion updates a given customer Question in the database
func (cmntRepo *QuestionGormRepo) UpdateQuestion(Question *entity.Question) (*entity.Question, []error) {
	cmnt := Question
	errs := cmntRepo.conn.Save(cmnt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

// DeleteQuestion deletes a given customer Question from the database
func (cmntRepo *QuestionGormRepo) DeleteQuestion(id uint) (*entity.Question, []error) {
	cmnt, errs := cmntRepo.Question(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = cmntRepo.conn.Delete(cmnt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

// StoreQuestion stores a given customer Question in the database
func (cmntRepo *QuestionGormRepo) StoreQuestion(Question *entity.Question) (*entity.Question, []error) {
	cmnt := Question
	errs := cmntRepo.conn.Create(cmnt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}
