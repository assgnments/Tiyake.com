package handler

import (
"encoding/json"
"html/template"
"net/http"
"teyake/answer"
"teyake/category"
"teyake/question"
"teyake/upvote"
)

// const questionKey = "id"
// const answerKey = "answer"

// const questionTitleKey = "question_title"
// const questionDescriptionKey = "question_description"
// const categoryFormKey = "category"
// const questionImagekey = "questionImagekey"


type QuestionAPIHandler struct {
	tmpl            *template.Template
	questionService question.QuestionService
	answerService   answer.AnswerService
	upvoteService   upvote.UpVoteService
	categoryService category.CategoryService
}


func NewQuestionAPIHandler(
	tmpl *template.Template,
	questionService question.QuestionService,
	answerService answer.AnswerService,
	categoryService category.CategoryService,
	) *QuestionAPIHandler {
	return &QuestionAPIHandler{
		tmpl:            tmpl,
		questionService: questionService,
		answerService:   answerService,
		categoryService: categoryService,
	}
}

 func (qah *QuestionAPIHandler) GetQuestionAPI(w http.ResponseWriter, r *http.Request){
	// questionIdString := r.URL.Query().Get(questionKey)
	// questionId, err := strconv.Atoi(questionIdString)
	r.PostForm.Get("")
	questions, errs := qah.questionService.Questions()
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(questions, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
