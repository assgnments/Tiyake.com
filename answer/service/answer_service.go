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
func (as *AnswerService) Answers() ([]entity.Answer, []error) {
	return as.AnswerRepo.Answers()
}

// Answer retrieves stored Answer by its id
func (as *AnswerService) Answer(id uint) (*entity.Answer, []error) {
	return as.AnswerRepo.Answer(id)
}

// UpdateAnswer updates a given Answer
func (as *AnswerService) UpdateAnswer(Answer *entity.Answer) (*entity.Answer, []error) {
	return as.AnswerRepo.UpdateAnswer(Answer)
}

// DeleteAnswer deletes a given Answer
func (as *AnswerService) DeleteAnswer(id uint) (*entity.Answer, []error) {
	return as.AnswerRepo.DeleteAnswer(id)
}

// StoreAnswer stores a given Answer
func (as *AnswerService) StoreAnswer(Answer *entity.Answer) (*entity.Answer, []error) {
	return  as.AnswerRepo.StoreAnswer(Answer)

}
