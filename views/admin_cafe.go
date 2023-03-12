package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func AdminCafeHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	session := uadmin.IsAuthenticated(r)
	if session == nil {
		http.Redirect(w, r, "/console/login", http.StatusSeeOther)
	}

	context["Title"] = "Console - Cafe"
	return context
}
