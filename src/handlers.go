package main

import (
	"encoding/json"
	"fmt"
	"forum/packages/credentials"
	"forum/packages/data"
	"forum/packages/utils"
	"html/template"
	"log"
	"net/http"
)

func generateTemplate(templateName string, filepaths []string) *template.Template {
	tmpl, err := template.New(templateName).Funcs(template.FuncMap{
		"getTimeSincePosted":  utils.GetTimeSincePosted,
		"getPagesArr":         utils.GetPagesArr,
		"GetPaginationValues": utils.GetPaginationValues,
	}).ParseFiles(filepaths...)
	// Error check:
	if err != nil {
		log.Fatal(err)
	}
	return tmpl
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	tData := getSession(r)
	tData.PageTitle = "404 Not Found"
	w.WriteHeader(http.StatusNotFound)

	tmpl := generateTemplate("base.html", []string{"templates/base.html", "templates/views/404.html"})
	tmpl.Execute(w, tData)
}

/* indexHandler handles the index page, parses most of the templates and executes them */
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tData := getSession(r)
	tData.PageTitle = "Home"

	categories, err := data.GetCategories()
	if err != nil {
		log.Fatal(err)
	}
	tData.Categories = categories

	tmpl := generateTemplate("base.html", []string{"templates/base.html", "templates/views/index.html", "templates/components/cat_navigation.html", "templates/components/popup_register.html", "templates/components/popup_login.html", "templates/components/header.html"})
	tmpl.Execute(w, tData)
}

/* registerHandler handles the registration form and redirects to the (temporary) success page */
func registerHandler(w http.ResponseWriter, r *http.Request) {
	tData := getSession(r)
	tData.PageTitle = "Register"

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

	tmpl := template.Must(template.ParseFiles("templates/components/popup_register.html"))
	tmpl.Execute(w, nil)
}

func topicsHandler(w http.ResponseWriter, r *http.Request) {
	tData := getSession(r)
	tData.PageTitle = "Register"

	categories, err := data.GetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tData.Categories = categories

	filters := data.RetrieveFilters(r)
	filters.UserID = tData.User.ID

	temp, err := data.GetTopicListData(filters)
	if err != nil {
		fmt.Println("Error in handlers.go")
		log.Fatal(err)
	}
	tData.Topics = temp.Topics
	tData.Filters = temp.Filters

	if r.Method == "POST" {
		r.ParseForm()
		tmpl := generateTemplate("", []string{"templates/components/topics-ctn.html", "templates/components/pagination.html"})
		tmpl.ExecuteTemplate(w, "topics-ctn", tData)
		return
	}
	tmpl := generateTemplate("topics.html", []string{"templates/views/topics.html", "templates/components/header.html", "templates/components/topics-ctn.html", "templates/components/pagination.html", "templates/components/cat_navigation.html", "templates/components/popup_register.html", "templates/components/popup_login.html"})
	tmpl.Execute(w, tData)
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

	tmpl := template.Must(template.ParseFiles("templates/components/popup_login.html"))
	tmpl.Execute(w, nil)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(r, &w)
	http.Redirect(w, r, "/topics", http.StatusSeeOther)
}
