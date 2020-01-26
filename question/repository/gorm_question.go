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
func (questionRepo *QuestionGormRepo) Questions() ([]entity.Question, []error) {
	questions := []entity.Question{}
	errs := questionRepo.conn.Set("gorm:auto_preload", true).Find(&questions).GetErrors()
	return questions, errs
}

// Question retrieves a customer Question from the database by its id
func (questionRepo *QuestionGormRepo) Question(id uint) (*entity.Question, []error) {
	question := entity.Question{}
	errs := questionRepo.conn.Set("gorm:auto_preload", true).First(&question, id).GetErrors()
	return &question, errs
}

// UpdateQuestion updates a given customer Question in the database
func (questionRepo *QuestionGormRepo) UpdateQuestion(question *entity.Question) (*entity.Question, []error) {
	errs := questionRepo.conn.Save(question).GetErrors()
	return question, errs
}

// DeleteQuestion deletes a given customer Question from the database
func (questionRepo *QuestionGormRepo) DeleteQuestion(id uint) (*entity.Question, []error) {
	question, errs := questionRepo.Question(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = questionRepo.conn.Delete(question, id).GetErrors()
	return question, errs
}

// StoreQuestion stores a given customer Question in the database
func (questionRepo *QuestionGormRepo) StoreQuestion(question *entity.Question) (*entity.Question, []error) {
	errs := questionRepo.conn.Create(question).GetErrors()
	return question, errs
}
func (questionRepo *QuestionGormRepo) QuestionByCategory(categoryId uint) ([]entity.Question, []error) {
	questions := []entity.Question{}
	errs := questionRepo.conn.Find(&questions, "category_id=?", categoryId).GetErrors()
	return questions, errs
}

func (questionRepo *QuestionGormRepo) SearchQuestions(searcheable string) ([]entity.Question, []error) {
	questions := []entity.Question{}

	errs := questionRepo.conn.Set("gorm:auto_preload", true).Where("Description like ? or Title like ? ", "%"+searcheable+"%", "%"+searcheable+"%").Find(&questions).GetErrors()
	return questions, errs
}

func (questionRepo *QuestionGormRepo) SearchByTitle(searcheable string) ([]entity.Question, []error) {
	questions := []entity.Question{}
	errs := questionRepo.conn.Set("gorm:auto_preload", true).Where("Title like ?", "%"+searcheable+"%").Find(&questions).GetErrors()
	return questions, errs
}

func (questionRepo *QuestionGormRepo) SearchByDescription(searcheable string) ([]entity.Question, []error) {
	questions := []entity.Question{}
	errs := questionRepo.conn.Set("gorm:auto_preload", true).Where("Description like ?", "%"+searcheable+"%").Find(&questions).GetErrors()
	return questions, errs
}
