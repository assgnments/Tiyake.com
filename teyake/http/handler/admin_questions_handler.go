package handler

import (
	"html/template"
	"net/http"
	"strconv"
	"teyake/category"
	"teyake/entity"
	"teyake/question"
)

type AdminQuestionHandler struct {
	tmpl            *template.Template
	questionService question.QuestionService
	categoryService category.CategoryService
	csrfSigningKey  []byte
}

func NewAdminQuestionHandler(tmpl *template.Template, questionService question.QuestionService,
	categoryService category.CategoryService, csrfSigningKey []byte) *AdminQuestionHandler {
	return &AdminQuestionHandler{
		tmpl:            tmpl,
		questionService: questionService,
		csrfSigningKey:  csrfSigningKey,
		categoryService: categoryService,
	}
}

func (aqh *AdminQuestionHandler) AdminQuestionHandler(w http.ResponseWriter, r *http.Request) {
	questionList := []entity.Question{}

	questions, errs := aqh.questionService.Questions()
	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	questionList = questions

	categories, errs := aqh.categoryService.Catagories()
	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	data := struct {
		Questions  []entity.Question
		Categories []entity.Category
	}{
		Questions:  questionList,
		Categories:categories,
	}
	aqh.tmpl.ExecuteTemplate(w, "admin.questions.layout", data)
}

func (aqh *AdminQuestionHandler) AdminQuestionsDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		_, errs := aqh.questionService.DeleteQuestion(uint(id))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}

	http.Redirect(w, r, "/admin/questions", http.StatusSeeOther)
}
