package answer

import "../entity"

// AnswerService specifies customer Answer related service
type AnswerService interface {
	Answers() ([]entity.Answer, []error)
	Answer(id uint) (*entity.Answer, []error)
	UpdateAnswer(Answer *entity.Answer) (*entity.Answer, []error)
	DeleteAnswer(id uint) (*entity.Answer, []error)
	StoreAnswer(Answer *entity.Answer) (*entity.Answer, []error)
}
