package handler

import (
	"html/template"
	"net/http"
	"teyake/entity"
	"teyake/form"
	"teyake/question"
	"teyake/util"
)

const Titlekey = "title"
const Descriptionkey = "description"
const UserIDkey = "user_ID"
const CreatedAtkey = "created_at"
const Imagekey = "image"

type QuestionHandler struct {
	tmpl            *template.Template
	questionService question.QuestionService
}

func NewQuestionHandler(tmpl *template.Template, questionService question.QuestionService) *QuestionHandler {
	return &QuestionHandler{
		tmpl:            tmpl,
		questionService: questionService,
	}
}

func (QuestionHandler *QuestionHandler) AddQuestion(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		QuestionHandler.tmpl.ExecuteTemplate(w, "", nil)
		return
	}
	if r.Method == http.MethodPost && util.ParseForm(w, r) {
		loginForm := form.Input{Values: r.PostForm, VErrors: form.VaildationErros{}}
		loginForm.ValidateRequiredFields(Titlekey, Descriptionkey, UserIDkey)
		Title := r.FormValue(Titlekey)
		Description := r.FormValue(Descriptionkey)
		Image := r.FormValue(Imagekey)
		//CreatedAt := not sure what to write here
		UserID := r.FormValue(UserIDkey)

		question := entity.Question{
			Title:       Title,
			Description: Description,
			Image:       Image,
			UserID:      UserID,
			//CreatedAt: ,
		}

		QuestionHandler.questionService.StoreQuestion(&question)
	}
}
