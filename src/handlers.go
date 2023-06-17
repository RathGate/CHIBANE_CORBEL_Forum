package main

import (
	"encoding/json"
	"fmt"
	"forum/packages/credentials"
	"forum/packages/data"
	"html/template"
	"log"
	"net/http"
)

func generateTemplate(templateName string, filepaths []string) *template.Template {
	tmpl, err := template.New(templateName).Funcs(template.FuncMap{
		"getTimeSincePosted": data.GetTimeSincePosted,
		"getPagesArr":        data.GetPagesArr,
		"getPagesValues":     data.GetPagesValues,
	}).ParseFiles(filepaths...)
	// Error check:
	if err != nil {
		log.Fatal(err)
	}
	return tmpl
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	// *Generates and executes templates:
	tmpl := generateTemplate("index.html", []string{"templates/views/index.html"})
	tmpl.Execute(w, userData)

}

/* indexHandler handles the index page, parses most of the templates and executes them */
func indexHandler(w http.ResponseWriter, r *http.Request) {
	currentUser := getSession(r)

	categories, err := dbGetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		User       data.ShortUser
		Categories []data.Category
	}{
		Categories: categories,
		User:       currentUser,
	}

	// Generates template:
	tmpl := template.New("index.html")

	// Parse the templates:
	tmpl = template.Must(tmpl.ParseFiles("templates/views/index.html", "templates/components/cat_navigation.html", "templates/components/register_form.html", "templates/components/login_form.html", "templates/components/navbar.html"))

	// Execute the templates
	err = tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

/* registerHandler handles the registration form and redirects to the (temporary) success page */
func registerHandler(w http.ResponseWriter, r *http.Request) {
	_ = getSession(r)
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")

		formValidation, lastInsertedID := credentials.RegisterNewUser(username, password, email)
		if lastInsertedID > 0 {
			err = setSession(r, &w, lastInsertedID)
			fmt.Println(err)
		}

		jsonValues, _ := json.Marshal(formValidation)
		w.Write(jsonValues)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/components/register_form.html"))
	tmpl.Execute(w, nil)
}

func topicsHandler(w http.ResponseWriter, r *http.Request) {
	userData.User = getSession(r)
	categories, err := dbGetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userData.Categories = categories

	filters := data.RetrieveFilters(r)
	filters.UserID = userData.UserID
	temp, err := data.GetTopics(filters)
	if err != nil {
		fmt.Println("Error in handlers.go")
		log.Fatal(err)
	}
	userData.Topics = temp.Topics
	userData.Filters = temp.Filters

	if r.Method == "POST" {
		r.ParseForm()
		tmpl := generateTemplate("", []string{"templates/components/topics-ctn.html", "templates/components/pagination.html"})
		tmpl.ExecuteTemplate(w, "topics-ctn", userData)
		return
	}
	tmpl := generateTemplate("topics.html", []string{"templates/views/topics.html", "templates/components/navbar.html", "templates/components/topics-ctn.html", "templates/components/pagination.html", "templates/components/cat_navigation.html", "templates/components/register_form.html", "templates/components/login_form.html"})
	tmpl.Execute(w, userData)
}

/* loginHandler handles the login form and redirects to the profile page */
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var formValidation credentials.FormValidation
		var userID int
		username := r.FormValue("username")
		password := r.FormValue("password")

		formValidation, userID = credentials.CheckUserCredentials(username, password)
		if userID > 0 {
			_ = setSession(r, &w, userID)
		}
		jsonValues, _ := json.Marshal(formValidation)
		w.Write(jsonValues)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/components/login_form.html"))
	tmpl.Execute(w, nil)
}

/* profileHandler handles the profile page */
func profileHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	username := r.URL.Query().Get("username")

	data := struct {
		ID       string
		Username string
	}{
		ID:       id,
		Username: username,
	}

	tmpl := template.Must(template.ParseFiles("templates/views/profile.html"))
	tmpl.Execute(w, data)
}

/* Temporary redirections for testing purposes */
func successHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/results/success.html"))
	tmpl.Execute(w, nil)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/results/error.html"))
	tmpl.Execute(w, nil)
}

/* Temporary redirections for testing purposes */

/* catHandler handles the category navigation */
func catHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	categories, err := dbGetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := struct {
		ID         string
		Name       string
		Categories []data.Category
	}{
		ID:         id,
		Name:       name,
		Categories: categories,
	}

	tmpl := template.Must(template.New("categories").ParseFiles("templates/components/cat_navigation.html"))
	err = tmpl.ExecuteTemplate(w, "categories", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(r, &w)
	http.Redirect(w, r, "/topics", http.StatusSeeOther)
}
