package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"teyake/answer"
	"teyake/category"
	"teyake/entity"
	"teyake/form"
	"teyake/question"
	"teyake/upvote"
)

// const questionKey = "id"
// const answerKey = "answer"

// const questionTitleKey = "question_title"
// const questionDescriptionKey = "question_description"
// const categoryFormKey = "category"
// const questionImagekey = "questionImagekey"

type QuestionHandler struct {
	tmpl            *template.Template
	questionService question.QuestionService
	answerService   answer.AnswerService
	upvoteService   upvote.UpVoteService
	categoryService category.CategoryService
	csrfSigningKey  []byte
}

type QuestionForm struct {
	FormInput form.Input
	Question  entity.Question
}
type NewQuestionForm struct {
	FormInput  form.Input
	Categories []entity.Category
}

func NewQuestionHandler(tmpl *template.Template, questionService question.QuestionService, answerService answer.AnswerService, categoryService category.CategoryService, upvoteService upvote.UpVoteService, csrfSigningKey []byte) *QuestionHandler {
	return &QuestionHandler{
		tmpl:            tmpl,
		questionService: questionService,
		csrfSigningKey:  csrfSigningKey,
		answerService:   answerService,
		upvoteService:   upvoteService,
		categoryService: categoryService,
	}
}

func (questionHandler *QuestionHandler) getQuestion(w http.ResponseWriter, r *http.Request) *QuestionForm {
	// questionIdString := r.URL.Query().Get(questionKey)
	// questionId, err := strconv.Atoi(questionIdString)
	questions, errs := questionHandler.questionService.Questions()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return nil
	}

	output, err := json.MarshalIndent(questions, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return nil

}
