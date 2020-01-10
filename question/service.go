package question

import "teyake/entity"

// QuestionService specifies customer Question related service
type QuestionService interface {
	Questions() ([]entity.Question, []error)
	Question(id uint) (*entity.Question, []error)
	UpdateQuestion(Question *entity.Question) (*entity.Question, []error)
	DeleteQuestion(id uint) (*entity.Question, []error)
	StoreQuestion(Question *entity.Question) (*entity.Question, []error)
}
