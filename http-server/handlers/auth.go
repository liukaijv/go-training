package handlers

import (
	"net/http"
	"html/template"
	"go-training/go-http/models"
	"time"
)

const AuthKey = "username"

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	h.Logger.Println(r.Method, r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		t, err := template.ParseFiles("views/register.html")
		if err != nil {
			h.Logger.Fatal(err)
			http.NotFound(w, r)
		}
		t.Execute(w, nil)
	case http.MethodPost:
		r.ParseForm()
		email, password := r.Form.Get("email"), r.Form.Get("password")
		r.Context()
		h.DB.Create(&models.User{Email: email, Password: password})
	default:
		http.NotFound(w, r)
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	h.Logger.Println(r.Method, r.URL.Path)

	var body struct {
		ErrorMsg string
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		username, password := r.Form.Get("username"), r.Form.Get("password")

		if username != "" && password != "" {
			expires := time.Now().AddDate(1, 0, 0)
			cookie := &http.Cookie{Name: AuthKey, Value: username, Expires: expires}
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/index", http.StatusFound)
			return
		} else {
			body.ErrorMsg = "Username and Password field must required"
		}
	}

	t, err := template.ParseFiles("views/login.html")
	if err != nil {
		h.Logger.Fatal(err)
	}
	t.Execute(w, body)

}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	h.Logger.Println(r.Method, r.URL.Path)
	cookie, _ := r.Cookie(AuthKey)
	if cookie != nil {
		cookie.Expires = time.Now().AddDate(-1, 0, 0)
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/login", http.StatusFound)
}

func (h *Handler) Authorized(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie(AuthKey)
		if err != nil || cookie.Value == "" {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	}
}
