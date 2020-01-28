package repository

import (
	"errors"
	"teyake/entity"
	"teyake/question"
)

// MockQuestionRepo implements menu.QuestionRepository interface
type MockQuestionRepo struct {
	questions map[uint]*entity.Question
}

// NewMockQuestionRepo returns new object of MockQuestionRepo
func NewMockQuestionRepo(questions map[uint]*entity.Question) question.QuestionRepository {
	return &MockQuestionRepo{questions}
}

// Questions returns all customer Questions stored in the database
func (questionRepo *MockQuestionRepo) Questions() ([]entity.Question, []error) {
	questions := []entity.Question{}
	for _, v := range questionRepo.questions {
		questions = append(questions, *v)
	}
	return questions, nil
}

// Question retrieves a customer Question from the database by its id
func (questionRepo *MockQuestionRepo) Question(id uint) (*entity.Question, []error) {
	return questionRepo.questions[id], nil
}

// UpdateQuestion updates a given customer Question in the database
func (questionRepo *MockQuestionRepo) UpdateQuestion(question *entity.Question) (*entity.Question, []error) {
	questionToUpdate := questionRepo.questions[question.ID]
	if questionToUpdate == nil {
		return nil, []error{
			errors.New("Question not found"),
		}
	}
	return questionToUpdate, nil
}

// DeleteQuestion deletes a given customer Question from the database
func (questionRepo *MockQuestionRepo) DeleteQuestion(id uint) (*entity.Question, []error) {
	question := questionRepo.questions[id]
	if question == nil {
		return nil, []error{
			errors.New("Question not found"),
		}
	}
	return question, nil

}

// StoreQuestion stores a given customer Question in the database
func (questionRepo *MockQuestionRepo) StoreQuestion(question *entity.Question) (*entity.Question, []error) {
	if question == nil {
		return nil, []error{
			errors.New("Can't create an empty question"),
		}
	}
	question.ID = uint(len(questionRepo.questions))
	questionRepo.questions[question.ID] = question
	return question, nil
}
func (questionRepo *MockQuestionRepo) QuestionByCategory(categoryId uint) ([]entity.Question, []error) {
	//questions := []entity.Question{}
	//errs := questionRepo.conn.Find(&questions, "category_id=?", categoryId).GetErrors()
	//return questions, errs
	return nil, nil
}

func (questionRepo *MockQuestionRepo) SearchQuestions(searcheable string) ([]entity.Question, []error) {
	return nil, nil
}

func (questionRepo *MockQuestionRepo) SearchByTitle(searcheable string) ([]entity.Question, []error) {
	return nil, nil
}

func (questionRepo *MockQuestionRepo) SearchByDescription(searcheable string) ([]entity.Question, []error) {
	return nil, nil
}
