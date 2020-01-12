package handler

import (
	"html/template"
	"net/http"
	"strconv"
	"teyake/category"
	"teyake/entity"
	"teyake/question"
)

const categoryTag="cat"
type IndexHandler struct {
	tmpl            *template.Template
	questionService question.QuestionService
	categoryService category.CategoryService
}

func NewIndexHandler(tmpl *template.Template, questionService question.QuestionService, categoryService category.CategoryService) *IndexHandler {
	return &IndexHandler{
		tmpl:            tmpl,
		questionService: questionService,
		categoryService: categoryService,
	}
}

func (indexHandler *IndexHandler) Index(w http.ResponseWriter, r *http.Request) {
	categoryId:=r.URL.Query().Get(categoryTag)
	questionList:=[]entity.Question{}
	if categoryId =="" {
		questions, errs := indexHandler.questionService.Questions()
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		questionList=questions
	}else{
		categoryId,err:=strconv.Atoi(categoryId)
		if err!=nil{
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		questions, errs := indexHandler.questionService.QuestionByCategory(uint(categoryId))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		questionList=questions
	}


	categories, errs := indexHandler.categoryService.Catagories()
	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	data := struct {
		Categories []entity.Category
		Questions  []entity.Question
	}{
		Categories: categories,
		Questions:  questionList,
	}
	indexHandler.tmpl.ExecuteTemplate(w, "index.layout", data)
}
