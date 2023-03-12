package views

import (
	"net/http"

	"github.com/mbdeguzman/cafe/models"
	"github.com/uadmin/uadmin"
)

func AdminIndexHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	session := uadmin.IsAuthenticated(r)
	if session == nil {
		http.Redirect(w, r, "/console/login", http.StatusSeeOther)
	}

	orders := []models.Orders{}
	uadmin.Filter(&orders, "status in (?,?,?)", 1, 2, 3)

	context["Title"] = "Console - Index"
	return context
}
