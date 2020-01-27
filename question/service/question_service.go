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
func (qs *QuestionService) Questions() ([]entity.Question, []error) {
	return qs.QuestionRepo.Questions()
}

// Question retrieves stored Question by its id
func (qs *QuestionService) Question(id uint) (*entity.Question, []error) {
	return qs.QuestionRepo.Question(id)
}

// UpdateQuestion updates a given Question
func (qs *QuestionService) UpdateQuestion(Question *entity.Question) (*entity.Question, []error) {
	return qs.QuestionRepo.UpdateQuestion(Question)
}

// DeleteQuestion deletes a given Question
func (qs *QuestionService) DeleteQuestion(id uint) (*entity.Question, []error) {
	return qs.QuestionRepo.DeleteQuestion(id)
}

// StoreQuestion stores a given Question
func (qs *QuestionService) StoreQuestion(Question *entity.Question) (*entity.Question, []error) {
	return qs.QuestionRepo.StoreQuestion(Question)
}
func (qs *QuestionService) QuestionByCategory(categoryId uint) ([]entity.Question, []error) {
	return qs.QuestionRepo.QuestionByCategory(categoryId)
}

func (qs *QuestionService) SearchQuestions(searchable string) ([]entity.Question, []error) {
	return qs.QuestionRepo.SearchQuestions(searchable)
}

func (qs *QuestionService) SearchByTitle(searchable string) ([]entity.Question, []error) {
	return qs.QuestionRepo.SearchByTitle(searchable)
}

func (qs *QuestionService) SearchByDescription(searchable string) ([]entity.Question, []error) {
	return qs.QuestionRepo.SearchByDescription(searchable)
}
