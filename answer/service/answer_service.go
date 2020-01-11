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
	return cs.AnswerRepo.Answers()
}

// Answer retrieves stored Answer by its id
func (cs *AnswerService) Answer(id uint) (*entity.Answer, []error) {
	return cs.AnswerRepo.Answer(id)
}

// UpdateAnswer updates a given Answer
func (cs *AnswerService) UpdateAnswer(Answer *entity.Answer) (*entity.Answer, []error) {
	return cs.AnswerRepo.UpdateAnswer(Answer)
}

// DeleteAnswer deletes a given Answer
func (cs *AnswerService) DeleteAnswer(id uint) (*entity.Answer, []error) {
	return cs.AnswerRepo.DeleteAnswer(id)
}

// StoreAnswer stores a given Answer
func (cs *AnswerService) StoreAnswer(Answer *entity.Answer) (*entity.Answer, []error) {
	return  cs.AnswerRepo.StoreAnswer(Answer)

}
