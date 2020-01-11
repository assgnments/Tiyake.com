package handler

import (
	"html/template"
	"net/http"
	"teyake/question"
)

type IndexHandler struct {
	tmpl *template.Template
	questionService question.QuestionService
}
func NewIndexHandler(tmpl *template.Template, questionService  question.QuestionService) *IndexHandler{
	return &IndexHandler{
		tmpl:tmpl,
		questionService:questionService,
	}
}

func (indexHandler *IndexHandler) Index(w http.ResponseWriter, r *http.Request){
	questions,errs:=indexHandler.questionService.Questions()
	if len(errs) > 0{
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	indexHandler.tmpl.ExecuteTemplate(w,"index.layout",questions)
}