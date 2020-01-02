package util

import "net/http"

func ParseForm(w http.ResponseWriter,r *http.Request) bool{
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return false
	}
	return true
}
