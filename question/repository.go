package question

import (
	"teyake/entity"
)

// QuestionRepository specifies customer Question related database operations
type QuestionRepository interface {
	Questions() ([]entity.Question, []error)
	Question(id uint) (*entity.Question, []error)
	UpdateQuestion(Question *entity.Question) (*entity.Question, []error)
	DeleteQuestion(id uint) (*entity.Question, []error)
	StoreQuestion(Question *entity.Question) (*entity.Question, []error)
}
