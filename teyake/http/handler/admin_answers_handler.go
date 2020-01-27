package handler

import (
	"html/template"
	"net/http"
	"strconv"
	"teyake/answer"
	"teyake/entity"
	"teyake/question"
)

type AdminAnswerHandler struct {
	tmpl            *template.Template
	answerService   answer.AnswerService
	questionService question.QuestionService
	csrfSigningKey  []byte
}

func NewAdminAnswerHandler(
	tmpl *template.Template,
	answerService answer.AnswerService,
	questionService question.QuestionService,
	csrfSigningKey []byte) *AdminAnswerHandler {
	return &AdminAnswerHandler{
		tmpl:            tmpl,
		answerService:   answerService,
		questionService: questionService,
		csrfSigningKey:  csrfSigningKey,
	}
}

func (aah *AdminAnswerHandler) AdminAnswerHandler(w http.ResponseWriter, r *http.Request) {
	answerList := []entity.Answer{}

	answers, errs := aah.answerService.Answers()
	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	answerList = answers

	questions, errs := aah.questionService.Questions()
	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	data := struct {
		Answers   []entity.Answer
		Questions []entity.Question
	}{
		Questions: questions,
		Answers:   answerList,
	}
	aah.tmpl.ExecuteTemplate(w, "admin.answers.layout", data)
}

func (aah *AdminAnswerHandler) AdminAnswersDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		_, errs := aah.answerService.DeleteAnswer(uint(id))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}

	http.Redirect(w, r, "/admin/answers", http.StatusSeeOther)
}
