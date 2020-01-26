package handler

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"html/template"
	"net/http"
	"strings"
	"teyake/entity"
	"teyake/form"
	"teyake/user"
	"teyake/util"
	"teyake/util/hash"
	"teyake/util/permission"
	"teyake/util/session"
	"teyake/util/token"
)

const fullnameKey = "fullname"
const passwordKey = "password"
const emailKey = "email"

const ctxUserSessionKey = "signed_in_user_session"

type UserHandler struct {
	tmpl           *template.Template
	userService    user.UserService
	sessionService user.SessionService
	roleService    user.RoleService
	csrfSignKey    []byte
}

func NewUserHandler(
	t *template.Template,
	userService user.UserService,
	sessionService user.SessionService,
	roleService user.RoleService,
	csKey []byte,
) *UserHandler {
	return &UserHandler{
		tmpl:           t,
		userService:    userService,
		sessionService: sessionService,
		roleService:    roleService,
		csrfSignKey:    csKey,
	}
}

// Authenticated checks if a user is authenticated to access a given route
func (userHandler *UserHandler) Authenticated(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		//Check if the user is logged in if not redirect the user to login page
		activeSession := userHandler.IsLoggedIn(r)
		if activeSession == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		//Provide context for the next handlers
		ctx := context.WithValue(r.Context(), ctxUserSessionKey, activeSession)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func (userHandler *UserHandler) getSigningKey(token *jwt.Token) (interface{}, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		sessionId := claims["sessionId"].(string)
		session, err := userHandler.sessionService.Session(sessionId)
		if len(err) > 0 {
			return nil, err[0]
		}
		return session.SigningKey, nil
	}
	return nil, nil
}

//Returns the user session id if it's logged in or nil
func (userHandler *UserHandler) IsLoggedIn(r *http.Request) *entity.Session {
	signedStringCookie, err := r.Cookie(session.SessionKey)
	if err != nil {
		return nil
	}

	sessionId := token.GetSessionIdFromToken(signedStringCookie.Value, userHandler.getSigningKey)
	if sessionId == "" {
		return nil
	}

	activeSession, errs := userHandler.sessionService.Session(sessionId)
	if len(errs) > 0 {
		return nil
	}

	return activeSession
}

// Authorized checks if a user has proper authority to access the given route
func (userHandler *UserHandler) Authorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		///Get the user for the current active session
		activeSession := r.Context().Value(ctxUserSessionKey).(*entity.Session)
		user, errs := userHandler.userService.User(activeSession.UUID)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		///Get the role of the user
		role, errs := userHandler.roleService.Role(user.RoleID)

		//Check if the user role is authorized to access the specific path and method requested
		if len(errs) > 0 || !permission.HasPermission(role.Name, r.URL.Path, r.Method) {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		//Check the validity of signed token inside the form if the form is post
		if r.Method == http.MethodPost {
			if !token.ISValidCSRF(r.FormValue(util.CSFRKey), userHandler.csrfSignKey) {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func (userHandler *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	//If it's requesting the login page return CSFR Signed token with the form
	CSFRToken, err := token.NewCSRFToken(userHandler.csrfSignKey)
	if r.Method == http.MethodGet {

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		userHandler.tmpl.ExecuteTemplate(w, "login.layout", form.Input{
			CSFR: CSFRToken,
		})
		return
	}
	//Only reply to forms that have that are parsable and have valid csfrToken
	if util.IsParsableFormPost(w, r, userHandler.csrfSignKey) {

		//Validate form data
		loginForm := form.Input{Values: r.PostForm, VErrors: form.VaildationErros{}}
		loginForm.ValidateRequiredFields(emailKey, passwordKey)
		loginForm.CSFR = CSFRToken
		email := r.FormValue(emailKey)
		password := r.FormValue(passwordKey)
		user, errs := userHandler.userService.UserByEmail(email)

		///Check form validity and user password
		if len(errs) > 0 || !hash.ArePasswordsSame(user.Password, password) {
			loginForm.VErrors.Add("generic", "Your email address or password is incorrect")
			userHandler.tmpl.ExecuteTemplate(w, "login.layout", loginForm)
			return
		}

		//At this point user is successfully logged in so creating a session
		newSession, errs := userHandler.sessionService.StoreSession(session.NewSession(user.ID))
		claims := token.NewClaims(newSession.SessionId, newSession.Expires)
		if len(errs) > 0 {
			loginForm.VErrors.Add("generic", "Failed to create session")
			userHandler.tmpl.ExecuteTemplate(w, "login.layout", loginForm)
			return
		}
		//Save session Id in cookies
		session.SetCookies(claims, newSession.Expires, newSession.SigningKey, w)

		//Check if user is an admin
		roles, errs := userHandler.userService.UserRoles(user)
		if userHandler.checkAdmin(roles) {
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
		}

		////Finally open the home page for the user
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
func (userHandler *UserHandler) checkAdmin(roles []entity.Role) bool {
	for _, role := range roles {
		if strings.ToUpper(role.Name) == strings.ToUpper("Admin") {
			return true
		}
	}
	return false
}

// Logout logout requests
func (userHandler *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	//Remove cookies
	session.RemoveCookies(w)
	//Delete session from the database
	currentSession, _ := r.Context().Value(ctxUserSessionKey).(*entity.Session)
	userHandler.sessionService.DeleteSession(currentSession.SessionId)
	//Redirect to login page
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
func (userHandler *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	CSFRToken, err := token.NewCSRFToken(userHandler.csrfSignKey)
	if r.Method == http.MethodGet {

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		userHandler.tmpl.ExecuteTemplate(w, "signup.layout", form.Input{CSFR: CSFRToken})

		return
	}
	//Only reply to forms that have that are parsable and have valid csfrToken
	if util.IsParsableFormPost(w, r, userHandler.csrfSignKey) {
		///Validate the form data
		signUpForm := form.Input{Values: r.PostForm, VErrors: form.VaildationErros{}}
		signUpForm.ValidateRequiredFields(fullnameKey, emailKey, passwordKey)
		signUpForm.ValidateEmail(emailKey)
		signUpForm.ValidatePassword(passwordKey)
		signUpForm.CSFR = CSFRToken
		if !signUpForm.IsValid() {
			userHandler.tmpl.ExecuteTemplate(w, "signup.layout", signUpForm)
			return
		}
		if userHandler.userService.EmailExists(r.FormValue(emailKey)) {
			signUpForm.VErrors.Add(emailKey, "This email is already in use!")
			userHandler.tmpl.ExecuteTemplate(w, "signup.layout", signUpForm)
			return
		}
		//Create password hash
		hashedPassword, err := hash.HashPassword(r.FormValue(passwordKey))
		if err != nil {
			signUpForm.VErrors.Add("password", "Password Could not be stored")
			userHandler.tmpl.ExecuteTemplate(w, "signup.layout", signUpForm)
			return
		}
		//Create a user role for the User
		role, errs := userHandler.roleService.RoleByName("USER")

		if len(errs) > 0 {
			signUpForm.VErrors.Add("generic", "Role couldn't be assigned to user")
			userHandler.tmpl.ExecuteTemplate(w, "signup.layout", signUpForm)
			return
		}
		///Get the data from the form and construct user object
		user := entity.User{
			FullName: r.FormValue(fullnameKey),
			Email:    r.FormValue(emailKey),
			Password: string(hashedPassword),
			RoleID:   role.ID,
		}
		// Save the user to the database
		_, ers := userHandler.userService.StoreUser(&user)
		if len(ers) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func (userHandler *UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	userHandler.tmpl.ExecuteTemplate(w, "index.layout", nil)
}

func (userHandler *UserHandler) Admin(w http.ResponseWriter, r *http.Request) {

	userHandler.tmpl.ExecuteTemplate(w, "admin.index.layout", nil)
}
