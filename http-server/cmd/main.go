package main

import (
	"fmt"
	"http-server/conf"
	"http-server/handlers"
	"http-server/models"
	"log"
	"net/http"
	"os"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

func main() {

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	db, err := models.NewDB()
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	db.InitSchema()

	h := handlers.New(logger, db)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/test", SayHello)

	http.HandleFunc("/", h.Authorized(h.Index))
	http.HandleFunc("/register", h.Register)
	http.HandleFunc("/login", h.Login)
	http.HandleFunc("/logout", h.Logout)

	err = http.ListenAndServe(":"+conf.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
