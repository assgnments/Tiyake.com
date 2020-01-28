package handler

import (
	"html/template"
	"net/http"
	"teyake/category"
	"teyake/entity"
	"teyake/question"
)
type ProfileHandler struct {
	tmpl            *template.Template
	questionService question.QuestionService
	categoryService category.CategoryService
}

func NewProfileHandler(tmpl *template.Template, questionService question.QuestionService, categoryService category.CategoryService) *ProfileHandler {
	return &ProfileHandler{
		tmpl:            tmpl,
		questionService: questionService,
		categoryService: categoryService,
	}
}

func (ph *ProfileHandler) UserQuestions(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		userid := r.FormValue("user_id")
		questions, errs := ph.questionService.SearchByID(userid)
		if len(errs) > 0 {
			panic(errs)
		}

		data := struct {
			Categories []entity.Category
			Questions  []entity.Question
		}{
			Categories: nil,
			Questions:  questions,
		}

		ph.tmpl.ExecuteTemplate(w, "user.profile.layout", data)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
