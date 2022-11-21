package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

type Users struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
}

var (
	templates = template.Must(template.ParseGlob("templates/*.html"))
	users     = []*Users{}
)

//Root slash html
func RootSlash(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	err := templates.ExecuteTemplate(w, "index", nil)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Ha habido un error a la hora de cargar los templates. Lo estamos solucionando.")
	}
}

//Function to get users
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if len(users) == 0 {
		w.WriteHeader(http.StatusOK)
		templates.ExecuteTemplate(w, "noUsers", nil)
	} else {
		templates.ExecuteTemplate(w, "getUsers", users)
	}
}

//Function to get ONE user by the id
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		fmt.Fprintf(w, "You didn't send an int on the url param.")
	}

	if len(users) == 0 {
		w.WriteHeader(http.StatusOK)
		templates.ExecuteTemplate(w, "noUsers", nil)
	} else {
		for _, v := range users {
			if v.ID == idInt {
				w.WriteHeader(http.StatusOK)
				templates.ExecuteTemplate(w, "getUser", v)
				break
			}
		}
	}

}

//Function to Post users
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	user := Users{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "You send something wrong in the body.")
	}

	w.WriteHeader(http.StatusCreated)
	users = append(users, &user)
	fmt.Fprintf(w, "The user was created correctly.")
}

//Function to delete users
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		fmt.Fprintf(w, "You didn't send an int on the url param.")
	}

	for i, v := range users {
		if v.ID == idInt {
			users = append(users[:i], users[i+1:]...)

			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "User deleted correctly.")
		}
	}
}

//Function to update user
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	user := &Users{}

	if err != nil {
		fmt.Fprintf(w, "You didn't send an int on the url param.")
	}

	for _, v := range users {
		if v.ID == idInt {
			err := json.NewDecoder(r.Body).Decode(user)
			w.WriteHeader(http.StatusOK)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "You send something wrong in the body.")
				break
			}

			v.Nombre = user.Nombre
			v.Apellido = user.Apellido
			v.Edad = user.Edad
			fmt.Fprintf(w, "User updated correctly.")
			break
		}
	}

}
