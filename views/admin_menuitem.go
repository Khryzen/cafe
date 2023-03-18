package views

import (
	"net/http"

	"github.com/mbdeguzman/cafe/models"
	"github.com/uadmin/uadmin"
)

func AdminMenuItemHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	session := uadmin.IsAuthenticated(r)
	if session == nil {
		http.Redirect(w, r, "/console/login", http.StatusSeeOther)
	}

	cafe := models.Cafe{}
	uadmin.Get(&cafe, "active = ?", true)

	menuitem := []models.MenuItems{}
	uadmin.All(&menuitem)
	for _, item := range menuitem {
		uadmin.Preload(&item)
	}

	context["Cafe"] = cafe
	context["MenuItem"] = menuitem
	context["Title"] = "Console - Cafe"
	return context
}
