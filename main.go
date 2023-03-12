package main

import (
	"net/http"

	"github.com/mbdeguzman/cafe/models"
	"github.com/mbdeguzman/cafe/views"
	"github.com/uadmin/uadmin"
)

func main() {
	uadmin.Register(
		models.Address{},
		models.Cafe{},
		models.Category{},
		models.Client{},
		models.MenuItems{},
		models.MenuItemMod{},
		models.MenuItemSize{},
		models.OrderLine{},
		models.Orders{},
	)

	http.HandleFunc("/console/", uadmin.Handler(views.AdminHandler))
	http.HandleFunc("/console/login", uadmin.Handler(views.AdminLoginHandler))

	uadmin.RootURL = "/admin"
	uadmin.Port = 8080
	uadmin.StartServer()
}
