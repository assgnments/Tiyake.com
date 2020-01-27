package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"teyake/category"
	"teyake/entity"
	"teyake/question"
)

const categoryTag = "cat"

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
	categoryId := r.URL.Query().Get(categoryTag)
	questionList := []entity.Question{}
	if categoryId == "" {
		questions, errs := indexHandler.questionService.Questions()
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		questionList = questions
	} else {
		categoryId, err := strconv.Atoi(categoryId)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		questions, errs := indexHandler.questionService.QuestionByCategory(uint(categoryId))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		questionList = questions
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

func (indexHandler *IndexHandler) SearchQuestions(w http.ResponseWriter, r *http.Request) {
	searchtype := r.URL.Query().Get("type")
	search := r.URL.Query().Get("searchable")
	fmt.Println(search)
	fmt.Println(("^^is the search query text"))
	fmt.Println(("search type is " + searchtype))
	questions := []entity.Question{}

	if searchtype == "Title" {
		questions, _ = indexHandler.questionService.SearchByTitle(search)
	} else if searchtype == "Description" {
		questions, _ = indexHandler.questionService.SearchByDescription(search)
	} else {
		questions, _ = indexHandler.questionService.SearchQuestions(search)
	}

	data := struct {
		Categories []entity.Category
		Questions  []entity.Question
	}{
		Categories: nil,
		Questions:  questions,
	}

	if len(questions) == 0 {
		data = struct {
			Categories []entity.Category
			Questions  []entity.Question
		}{
			Categories: nil,
			Questions:  nil,
		}
	}

	indexHandler.tmpl.ExecuteTemplate(w, "index.layout", data)
}
