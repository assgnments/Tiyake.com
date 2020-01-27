package handler

import (
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"teyake/answer"
	"teyake/category"
	"teyake/entity"
	"teyake/form"
	"teyake/question"
	"teyake/upvote"
	"teyake/util"
	"teyake/util/token"
)

const questionKey = "id"
const answerKey = "answer"

const questionTitleKey = "question_title"
const questionDescriptionKey = "question_description"
const categoryFormKey = "category"
const questionImagekey = "questionImagekey"

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

func (questionHandler *QuestionHandler) QuestionHandler(w http.ResponseWriter, r *http.Request) {
	answerForm := questionHandler.getQuestion(w, r)
	if answerForm == nil {
		return
	}
	if r.Method == http.MethodGet {
		questionHandler.tmpl.ExecuteTemplate(w, "question.detail.layout", answerForm)
		return
	}
	if util.IsParsableFormPost(w, r, questionHandler.csrfSigningKey) {
		//Check if the form is valid
		answerForm.FormInput.ValidateRequiredFields(answerKey)
		if !answerForm.FormInput.IsValid() {
			questionHandler.tmpl.ExecuteTemplate(w, "question.detail.layout", answerForm)
			return
		}
		//Save the answer
		currentSession, _ := r.Context().Value(ctxUserSessionKey).(*entity.Session)
		answer := entity.Answer{
			UserID:     currentSession.UUID,
			Message:    r.FormValue(answerKey),
			QuestionID: answerForm.Question.ID,
		}
		_, errs := questionHandler.answerService.StoreAnswer(&answer)
		//Reload question to show change
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		answerForm := questionHandler.getQuestion(w, r)
		if answerForm == nil {
			return
		}
		questionHandler.tmpl.ExecuteTemplate(w, "question.detail.layout", answerForm)
	}
}

func (questionHandler *QuestionHandler) NewQuestion(w http.ResponseWriter, r *http.Request) {
	CSFRToken, err := token.NewCSRFToken(questionHandler.csrfSigningKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	categories, errs := questionHandler.categoryService.Catagories()
	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	newQuestionForm := NewQuestionForm{

		FormInput: form.Input{
			CSFR:    CSFRToken,
			Values:  r.PostForm,
			VErrors: form.VaildationErros{},
		},
		Categories: categories,
	}

	if r.Method == http.MethodGet {
		questionHandler.tmpl.ExecuteTemplate(w, "add_question.layout", newQuestionForm)
		return
	}
	if util.IsParsableFormPost(w, r, questionHandler.csrfSigningKey) {
		newQuestionForm.FormInput.ValidateRequiredFields(questionTitleKey, questionDescriptionKey, categoryFormKey)
		if !newQuestionForm.FormInput.IsValid() {
			questionHandler.tmpl.ExecuteTemplate(w, "add_question.layout", newQuestionForm)
			return
		}
		categoryIdString := r.FormValue(categoryFormKey)
		categoryId, err := strconv.Atoi(categoryIdString)
		if err != nil {
			newQuestionForm.FormInput.VErrors.Add(categoryFormKey, "Invalid category id")
			questionHandler.tmpl.ExecuteTemplate(w, "add_question.layout", newQuestionForm)
			return
		}
		multiPartFile, fileHeader, err := r.FormFile(questionImagekey)
		img := ""
		if err == nil {
			defer multiPartFile.Close()
			writeFile(&multiPartFile, fileHeader.Filename)
			img = fileHeader.Filename
		}

		currentSession, _ := r.Context().Value(ctxUserSessionKey).(*entity.Session)
		question := entity.Question{
			Title:       r.FormValue(questionTitleKey),
			Description: r.FormValue(questionDescriptionKey),
			Image:       img,
			UserID:      currentSession.UUID,
			CategoryID:  uint(categoryId),
			Answers:     nil,
		}

		savedQuestion, errs := questionHandler.questionService.StoreQuestion(&question)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		url := fmt.Sprintf("/question?id=%d", savedQuestion.ID)
		http.Redirect(w, r, url, http.StatusSeeOther)
		return
	}
}
func (questionHandler *QuestionHandler) getQuestion(w http.ResponseWriter, r *http.Request) *QuestionForm {
	questionIdString := r.URL.Query().Get(questionKey)
	questionId, err := strconv.Atoi(questionIdString)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return nil
	}

	question, _ := questionHandler.questionService.Question(uint(questionId))
	CSFRToken, err := token.NewCSRFToken(questionHandler.csrfSigningKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return nil
	}
	return &QuestionForm{
		Question: *question,
		FormInput: form.Input{
			Values: r.PostForm,
			CSFR:   CSFRToken,
		},
	}
}

func writeFile(mf *multipart.File, fname string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	path := filepath.Join(wd, "ui", "assets", "img", fname)
	image, err := os.Create(path)
	if err != nil {
		return err
	}
	defer image.Close()
	io.Copy(image, *mf)
	return nil
}
