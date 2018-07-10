package handlers

import (
	"log"
	"go-training/go-http/models"
	"net/http"
	"html/template"
)

type Person struct {
	UserName string
}

type Handler struct {
	Logger *log.Logger
	DB     *models.DB
}

func New(logger *log.Logger, db *models.DB) *Handler {
	return &Handler{Logger: logger, DB: db}
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	h.Logger.Println(r.Method, r.URL.Path)
	var body struct {
		UserName string
	}

	switch r.Method {
	case http.MethodGet:
		t, err := template.ParseFiles("views/index.html")
		if err != nil {
			h.Logger.Fatal(err)
			http.NotFound(w, r)
		}
		cookie, _ := r.Cookie(AuthKey)
		if cookie != nil {
			body.UserName = cookie.Value
		}
		t.Execute(w, body)
	default:
		http.NotFound(w, r)
	}
}
