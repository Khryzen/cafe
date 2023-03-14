package views

import (
	"net/http"

	"github.com/mbdeguzman/cafe/models"
	"github.com/uadmin/uadmin"
)

func AdminCafeHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	session := uadmin.IsAuthenticated(r)
	if session == nil {
		http.Redirect(w, r, "/console/login", http.StatusSeeOther)
	}

	cafe := models.Cafe{}
	uadmin.Get(&cafe, "active = ?", true)

	context["Cafe"] = cafe
	context["Title"] = "Console - Cafe"
	return context
}
