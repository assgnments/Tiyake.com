package service

import (
	"teyake/answer"
	"teyake/entity"
)

// AnswerService implements menu.AnswerService interface
type AnswerService struct {
	AnswerRepo answer.AnswerRepository
}

// NewAnswerService returns a new AnswerService object
func NewAnswerService(commRepo answer.AnswerRepository) answer.AnswerService {
	return &AnswerService{AnswerRepo: commRepo}
}

// Answers returns all stored Answers
func (cs *AnswerService) Answers() ([]entity.Answer, []error) {
	cmnts, errs := cs.AnswerRepo.Answers()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnts, errs
}

// Answer retrieves stored Answer by its id
func (cs *AnswerService) Answer(id uint) (*entity.Answer, []error) {
	cmnt, errs := cs.AnswerRepo.Answer(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

// UpdateAnswer updates a given Answer
func (cs *AnswerService) UpdateAnswer(Answer *entity.Answer) (*entity.Answer, []error) {
	cmnt, errs := cs.AnswerRepo.UpdateAnswer(Answer)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

// DeleteAnswer deletes a given Answer
func (cs *AnswerService) DeleteAnswer(id uint) (*entity.Answer, []error) {
	cmnt, errs := cs.AnswerRepo.DeleteAnswer(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

// StoreAnswer stores a given Answer
func (cs *AnswerService) StoreAnswer(Answer *entity.Answer) (*entity.Answer, []error) {
	cmnt, errs := cs.AnswerRepo.StoreAnswer(Answer)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}
