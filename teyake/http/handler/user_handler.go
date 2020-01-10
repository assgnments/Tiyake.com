package handler

import (
	"html/template"
	"net/http"
	"teyake/entity"
	"teyake/form"
	"teyake/user"
	"teyake/util"
)

const fullnameKey = "fullname"
const passwordKey = "password"
const emailKey = "email"

type UserHandler struct {
	tmpl        *template.Template
	userService user.UserService
}

func NewUserHandler(tmpl *template.Template, userService user.UserService) *UserHandler {
	return &UserHandler{
		tmpl:        tmpl,
		userService: userService,
	}
}

func (userHandler *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	//If it's requesting the login page
	if r.Method == http.MethodGet {
		userHandler.tmpl.ExecuteTemplate(w, "login.layout", form.Input{})
		return
	}
	if r.Method == http.MethodPost {
		//Validate form data
		loginForm := form.Input{Values: r.PostForm, VErrors: form.VaildationErros{}}
		loginForm.ValidateRequiredFields(emailKey, passwordKey)
		email := r.FormValue(emailKey)
		password := r.FormValue(passwordKey)
		user, errs := userHandler.userService.UserByEmail(email)
		if len(errs) > 0 || user.Password != password {
			loginForm.VErrors.Add("generic", "Your email address or password is incorrect")
			userHandler.tmpl.ExecuteTemplate(w, "login.layout", loginForm)
			return
		}
		userHandler.tmpl.ExecuteTemplate(w, "signup.layout", loginForm)
	}
}
func (userHandler *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		userHandler.tmpl.ExecuteTemplate(w, "signup.layout", form.Input{})
		return
	}
	if r.Method == http.MethodPost && util.ParseForm(w, r) {
		///Validate the form data
		signUpForm := form.Input{Values: r.PostForm, VErrors: form.VaildationErros{}}
		signUpForm.ValidateRequiredFields(fullnameKey, emailKey, passwordKey)
		signUpForm.ValidateEmail(emailKey)
		signUpForm.ValidatePassword(passwordKey)

		if !signUpForm.IsValid() {
			userHandler.tmpl.ExecuteTemplate(w, "signup.layout", signUpForm)
			return
		}
		if userHandler.userService.EmailExists(r.FormValue(emailKey)) {
			signUpForm.VErrors.Add(emailKey, "This email is already in use!")
			userHandler.tmpl.ExecuteTemplate(w, "signup.layout", signUpForm)
			return
		}

		///Get the data from the form and construct user object
		user := entity.User{
			FullName: r.FormValue(fullnameKey),
			Email:    r.FormValue(emailKey),
			Password: r.FormValue(passwordKey),
		}
		_, errs := userHandler.userService.StoreUser(&user)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
func (userHandler *UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	userHandler.tmpl.ExecuteTemplate(w, "index.layout", nil)
}
