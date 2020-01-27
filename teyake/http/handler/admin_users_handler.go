package handler

import (
	"github.com/jinzhu/gorm"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
	"teyake/entity"
	"teyake/form"
	"teyake/user"
	"teyake/util/hash"
	"teyake/util/token"
)

type AdminUsersHandler struct {
	tmpl        *template.Template
	userService user.UserService
	roleService user.RoleService
	csrfSignKey []byte
}

func NewAdminUsersHandler(
	t *template.Template,
	userService user.UserService,
	roleService user.RoleService,
	csrfSigningKey []byte,
) *AdminUsersHandler {
	return &AdminUsersHandler{
		tmpl:        t,
		userService: userService,
		roleService: roleService,
		csrfSignKey: csrfSigningKey,
	}
}

func (auh *AdminUsersHandler) AdminUsers(w http.ResponseWriter, r *http.Request) {
	users, errs := auh.userService.Users()
	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	token, err := token.NewCSRFToken(auh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	data := struct {
		Users []entity.User
		CSRF  string
	}{
		Users: users,
		CSRF:  token,
	}
	auh.tmpl.ExecuteTemplate(w, "admin.users.layout", data)
}

func (auh *AdminUsersHandler) AdminUsersUpdate(w http.ResponseWriter, r *http.Request) {
	token, err := token.NewCSRFToken(auh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		user, errs := auh.userService.User(uint(id))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		roles, errs := auh.roleService.Roles()
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		role, errs := auh.roleService.Role(user.RoleID)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		values := url.Values{}
		values.Add("userid", idRaw)
		values.Add("fullname", user.FullName)
		values.Add("email", user.Email)
		values.Add("role", string(user.RoleID))
		values.Add("rolename", role.Name)

		userData := struct {
			Values  url.Values
			VErrors form.VaildationErros
			Roles   []entity.Role
			User    *entity.User
			CSRF    string
		}{
			Values:  values,
			VErrors: form.VaildationErros{},
			Roles:   roles,
			User:    user,
			CSRF:    token,
		}
		auh.tmpl.ExecuteTemplate(w, "admin.users.update.layout", userData)
	}
	if r.Method == http.MethodPost {

		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		updateUserForm := form.Input{Values: r.PostForm, VErrors: form.VaildationErros{}}
		updateUserForm.ValidateRequiredFields("fullname", "email", "role")
		updateUserForm.ValidateEmail("email")
		updateUserForm.CSFR = token
		if !updateUserForm.IsValid() {
			auh.tmpl.ExecuteTemplate(w, "admin.users.update.layout", updateUserForm)
			return
		}
		userID := r.FormValue("userid")
		uid, err := strconv.Atoi(userID)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		user, errs := auh.userService.User(uint(uid))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		emailExists := auh.userService.EmailExists(r.FormValue("email"))
		if (user.Email != r.FormValue("email")) && emailExists {
			updateUserForm.VErrors.Add("email", "Email Already exists")
			auh.tmpl.ExecuteTemplate(w, "admin.users.update.layout", updateUserForm)
			return
		}
		roleId, err := strconv.Atoi(r.FormValue("role"))
		if err != nil {
			updateUserForm.VErrors.Add("role", "Could not retrieve role id")
			auh.tmpl.ExecuteTemplate(w, "admin.users.update.layout", updateUserForm)
			return
		}
		usr := &entity.User{
			Model:    gorm.Model{ID: user.ID},
			FullName: r.FormValue("fullname"),
			Email:    r.FormValue("email"),
			Password: user.Password,
			RoleID:   uint(roleId),
		}
		_, errs = auh.userService.UpdateUser(usr)

		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/users", http.StatusSeeOther)

	}

}

func (auh *AdminUsersHandler) AdminUsersNew(w http.ResponseWriter, r *http.Request) {
	token, err := token.NewCSRFToken(auh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		roles, errs := auh.roleService.Roles()
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		newUserForm := struct {
			Values  url.Values
			VErrors form.VaildationErros
			Roles   []entity.Role
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			Roles:   roles,
			CSRF:    token,
		}
		auh.tmpl.ExecuteTemplate(w, "admin.users.new.layout", newUserForm)
		return
	}
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Validate the form contents
		newUserForm := form.Input{Values: r.PostForm, VErrors: form.VaildationErros{}}
		newUserForm.ValidateRequiredFields("fullname", "email", "password")
		newUserForm.ValidateEmail("email")
		newUserForm.ValidatePassword("password")
		newUserForm.CSFR = token
		// If there are any errors, redisplay the signup form.
		if !newUserForm.IsValid() {
			auh.tmpl.ExecuteTemplate(w, "admin.users.new.layout", newUserForm)
			return
		}

		if auh.userService.EmailExists(r.FormValue(emailKey)) {
			newUserForm.VErrors.Add(emailKey, "This email is already in use!")
			auh.tmpl.ExecuteTemplate(w, "admin.users.new.layout", newUserForm)
			return
		}

		hashedPassword, err := hash.HashPassword(r.FormValue(passwordKey))
		if err != nil {
			newUserForm.VErrors.Add("password", "Password Could not be stored")
			auh.tmpl.ExecuteTemplate(w, "admin.users.new.layout", newUserForm)
			return
		}

		roleId, err := strconv.Atoi(r.FormValue("role"))

		if err != nil {
			newUserForm.VErrors.Add("role", "could not retrieve role id")
			auh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", newUserForm)
			return
		}
		///Get the data from the form and construct user object
		user := entity.User{
			FullName: r.FormValue(fullnameKey),
			Email:    r.FormValue(emailKey),
			Password: string(hashedPassword),
			RoleID:   uint(roleId),
		}
		// Save the user to the database
		_, ers := auh.userService.StoreUser(&user)
		if len(ers) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
	}
}

func (auh *AdminUsersHandler) AdminUsersDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		_, errs := auh.userService.DeleteUser(uint(id))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}

	http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}
