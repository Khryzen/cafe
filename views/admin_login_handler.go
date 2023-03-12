package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	context := map[string]interface{}{}
	user := uadmin.User{}
	username := r.FormValue("username")
	password := r.FormValue("password")

	if r.Method == "POST" {
		sess := uadmin.IsAuthenticated(r)
		if sess != nil {
			http.Redirect(w, r, "/console/index", http.StatusSeeOther)
			return
		}

		uadmin.Get(&user, "username = ?", username)
		session := user.Login(password, "")

		if session != nil {
			http.SetCookie(w, &http.Cookie{
				Path:  "/",
				Name:  "session",
				Value: session.Key,
			})
			http.Redirect(w, r, "/console/index", http.StatusSeeOther)
		}
	}
	uadmin.RenderHTML(w, r, "templates/console/login.html", context)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session := uadmin.Session{}
	key := uadmin.GetUserFromRequest(r).GetActiveSession().Key
	uadmin.DeleteList(&session, "`key` = ?", key)

	// Expire session cookie on logout
	sessionCookie := &http.Cookie{
		Name:   "session",
		Path:   "/",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, sessionCookie)
	http.Redirect(w, r, "/console/login", http.StatusSeeOther)
}
