package handler

import (
	"github.com/jinzhu/gorm"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
	"teyake/category"
	"teyake/entity"
	"teyake/form"
	"teyake/util/token"
)

type AdminCategoryHandler struct {
	tmpl            *template.Template
	categoryService category.CategoryService
	csrfSigningKey  []byte
}

func NewAdminCategoryHandler(
	tmpl *template.Template,
	categoryService category.CategoryService,
	csrfSigningKey []byte) *AdminCategoryHandler {
	return &AdminCategoryHandler{
		tmpl:            tmpl,
		categoryService: categoryService,
		csrfSigningKey:  csrfSigningKey,
	}
}

func (ach *AdminCategoryHandler) AdminCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	categoryList := []entity.Category{}

	categories, errs := ach.categoryService.Catagories()
	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	categoryList = categories

	data := struct {
		Categories []entity.Category
	}{
		Categories: categoryList,
	}
	ach.tmpl.ExecuteTemplate(w, "admin.categories.layout", data)
}

func (ach *AdminCategoryHandler) AdminCategoriesNew(w http.ResponseWriter, r *http.Request) {
	token, err := token.NewCSRFToken(ach.csrfSigningKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {

		newCategoryForm := struct {
			Values  url.Values
			VErrors form.VaildationErros

			CSRF string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		ach.tmpl.ExecuteTemplate(w, "admin.categories.new.layout", newCategoryForm)
		return
	}
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		updateCategoryForm := form.Input{Values: r.PostForm, VErrors: form.VaildationErros{}}
		updateCategoryForm.ValidateRequiredFields("category")
		updateCategoryForm.CSFR = token

		if !updateCategoryForm.IsValid() {
			ach.tmpl.ExecuteTemplate(w, "admin.categories.new.layout", updateCategoryForm)
			return
		}

		///Get the data from the form and construct user object
		category := entity.Category{
			Name: r.FormValue("category"),
		}
		// Save the user to the database
		_, ers := ach.categoryService.StoreCategory(&category)
		if len(ers) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)
	}
}

func (ach *AdminCategoryHandler) AdminCategoriesUpdate(w http.ResponseWriter, r *http.Request) {
	token, err := token.NewCSRFToken(ach.csrfSigningKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		catid, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		category, errs := ach.categoryService.Category(uint(catid))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		values := url.Values{}
		values.Add("id", idRaw)
		values.Add("category", category.Name)

		newCategoryForm := struct {
			Values  url.Values
			VErrors form.VaildationErros
			CSRF    string
		}{
			Values:  values,
			VErrors: form.VaildationErros{},
			CSRF:    token,
		}
		ach.tmpl.ExecuteTemplate(w, "admin.categories.update.layout", newCategoryForm)
		return
	}
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		updateCategoryForm := form.Input{Values: r.PostForm, VErrors: form.VaildationErros{}}
		updateCategoryForm.ValidateRequiredFields("category")
		updateCategoryForm.CSFR = token

		if !updateCategoryForm.IsValid() {
			ach.tmpl.ExecuteTemplate(w, "admin.categories.update.layout", updateCategoryForm)
			return
		}
		ID := r.FormValue("id")
		catid, err := strconv.Atoi(ID)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		cat, errs := ach.categoryService.Category(uint(catid))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		///Get the data from the form and construct user object
		category := entity.Category{
			Model: gorm.Model{ID: cat.ID},
			Name:  r.FormValue("category"),
		}
		// Save the user to the database
		_, ers := ach.categoryService.UpdateCategory(&category)
		if len(ers) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)
	}
}
func (ach *AdminCategoryHandler) AdminCategoriesDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		_, errs := ach.categoryService.DeleteCategory(uint(id))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
	http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)
}
