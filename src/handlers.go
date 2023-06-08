package main

import (
	"database/sql"
	"fmt"
	"forum/packages/credentials"
	"forum/packages/structs"
	"html/template"
	"net/http"
)

/* notFoundHandler handles the 404 page */
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	// Generates and executes templates:
	tmpl := template.Must(template.ParseFiles("templates/views/index.html"))
	tmpl.Execute(w, nil)
}

/* indexHandler handles the index page, parses most of the templates and executes them */
func indexHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := dbGetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Categories []structs.Category
	}{
		Categories: categories,
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
	if r.Method == "POST" {

		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")

		if len(password) < 8 {
			http.Error(w, "Password must be at least 8 characters long", http.StatusBadRequest)
			return
		}

		if !credentials.ContainsLetter(password) || !credentials.ContainsDigit(password) || !credentials.ContainsSpecialChar(password) {
			http.Error(w, "Password must contain at least one letter, one digit, and one special character", http.StatusBadRequest)
			return
		}

		if !credentials.IsValidEmail(email) {
			http.Error(w, "Invalid email address", http.StatusBadRequest)
			return
		}

		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum?parseTime=true")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		var stmt *sql.Stmt
		stmt, err = db.Prepare("INSERT INTO users (username, password, email, role_id, isActive) VALUES (?, ?, ?, ?, ?)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		var result sql.Result
		result, err = stmt.Exec(username, password, email, 3, 1)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if rowsAffected > 0 {
			http.Redirect(w, r, "/success", http.StatusFound)
		} else {

			http.Redirect(w, r, "/error", http.StatusFound)
		}
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/components/register_form.html"))
	tmpl.Execute(w, nil)
}

/* loginHandler handles the login form and redirects to the profile page */
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		user, err := credentials.ValidateUser(username, password)
		if err != nil {
			http.Redirect(w, r, "/error", http.StatusFound)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/profile?id=%d&username=%s", user.ID, user.Username), http.StatusFound)
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
		Categories []structs.Category
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
