package util

import (
	"net/http"
	"teyake/util/token"
)
const CSFRKey = "csfrKey"
func ParseForm(w http.ResponseWriter,r *http.Request) bool{
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return false
	}
	return true
}
func  IsParsableFormPost(w http.ResponseWriter, r *http.Request,csrfSignKey []byte,) bool {
	return r.Method == http.MethodPost &&
		ParseForm(w, r) &&
		token.ISValidCSRF(r.FormValue(CSFRKey), csrfSignKey)
}
