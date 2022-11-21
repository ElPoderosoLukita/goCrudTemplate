package main

import (
	"net/http"
	"text/template"

	"github.com/ElPoderosoLukita/goCRUD2/handlers"
)

var (
	templates = template.Must(template.ParseGlob("templates/*.html"))
)

func main() {
	staticFiles := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", staticFiles))

	http.HandleFunc("/", handlers.RootSlash)
	http.HandleFunc("/get/users", handlers.GetUsersHandler)
	http.HandleFunc("/get/user", handlers.GetUserHandler)
	http.HandleFunc("/post/user", handlers.PostUserHandler)
	http.HandleFunc("/delete/user", handlers.DeleteUserHandler)
	http.HandleFunc("/update/user", handlers.UpdateUserHandler)

	http.ListenAndServe(":8081", nil)
}
