package views

import (
	"net/http"
	"strings"

	"github.com/mbdeguzman/cafe/models"
	"github.com/uadmin/uadmin"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	context := map[string]interface{}{}

	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/console/")
	page := strings.TrimSuffix(r.URL.Path, "/")

	switch page {
	case "index":
		context = AdminIndexHandler(w, r)
	case "cafe":
		context = AdminCafeHandler(w, r)
	case "category":
		context = AdminCategoryHandler(w, r)
	default:
		page = "index"
		context = AdminIndexHandler(w, r)
	}
	RenderInterface(w, r, page, context)
}
func RenderInterface(w http.ResponseWriter, r *http.Request, page string, context map[string]interface{}) {
	templateList := []string{}
	templateList = append(templateList, "./templates/console/base.html")

	path := "./templates/console/" + page + ".html"
	templateList = append(templateList, path)

	cafe := models.Cafe{}
	if uadmin.Count(&cafe, "active = ?", true) > 0 {
		uadmin.Get(&cafe, "active = ?", true)
		context["Title"] = cafe.BusinessName
	}
	uadmin.RenderMultiHTML(w, r, templateList, context)
}
