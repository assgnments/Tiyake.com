package answer

import "../entity"

// AnswerRepository specifies customer Answer related database operations
type AnswerRepository interface {
	Answers() ([]entity.Answer, []error)
	Answer(id uint) (*entity.Answer, []error)
	UpdateAnswer(Answer *entity.Answer) (*entity.Answer, []error)
	DeleteAnswer(id uint) (*entity.Answer, []error)
	StoreAnswer(Answer *entity.Answer) (*entity.Answer, []error)
}
