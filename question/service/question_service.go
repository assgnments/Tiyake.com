package service

import (
	"teyake/entity"
	"teyake/question"
)

// QuestionService implements menu.QuestionService interface
type QuestionService struct {
	QuestionRepo question.QuestionRepository
}

// NewQuestionService returns a new QuestionService object
func NewQuestionService(commRepo question.QuestionRepository) question.QuestionService {
	return &QuestionService{QuestionRepo: commRepo}
}

// Questions returns all stored Questions
func (cs *QuestionService) Questions() ([]entity.Question, []error) {
	cmnts, errs := cs.QuestionRepo.Questions()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnts, errs
}

// Question retrieves stored Question by its id
func (cs *QuestionService) Question(id uint) (*entity.Question, []error) {
	cmnt, errs := cs.QuestionRepo.Question(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

// UpdateQuestion updates a given Question
func (cs *QuestionService) UpdateQuestion(Question *entity.Question) (*entity.Question, []error) {
	cmnt, errs := cs.QuestionRepo.UpdateQuestion(Question)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

// DeleteQuestion deletes a given Question
func (cs *QuestionService) DeleteQuestion(id uint) (*entity.Question, []error) {
	cmnt, errs := cs.QuestionRepo.DeleteQuestion(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

// StoreQuestion stores a given Question
func (cs *QuestionService) StoreQuestion(Question *entity.Question) (*entity.Question, []error) {
	cmnt, errs := cs.QuestionRepo.StoreQuestion(Question)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}
