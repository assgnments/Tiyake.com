package question

import (
	"teyake/entity"
)

// QuestionRepository specifies customer Question related database operations
type QuestionRepository interface {
	QuestionByCategory(categoryId uint) ([]entity.Question, []error)
	Questions() ([]entity.Question, []error)
	Question(id uint) (*entity.Question, []error)
	UpdateQuestion(Question *entity.Question) (*entity.Question, []error)
	DeleteQuestion(id uint) (*entity.Question, []error)
	StoreQuestion(Question *entity.Question) (*entity.Question, []error)
	SearchQuestions(searchable string) ([]entity.Question, []error)
	SearchByTitle(searchable string) ([]entity.Question, []error)
	SearchByDescription(searchable string) ([]entity.Question, []error)
}
