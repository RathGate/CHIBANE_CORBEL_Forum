package main

import (
	"database/sql"
	"fmt"
	"forum/packages/credentials"
	"forum/packages/dbData"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
)

func generateTemplate(templateName string, filepaths []string) *template.Template {
	tmpl, err := template.New(templateName).Funcs(template.FuncMap{
		"getTimeSincePosted": dbData.GetTimeSincePosted,
	}).ParseFiles(filepaths...)
	if err != nil {
		log.Fatal(err)
	}
	return tmpl
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tmpl := generateTemplate("index.html", []string{"templates/views/index.html"})
	tmpl.Execute(w, userData)
}

/* indexHandler handles the index page and executes most of the templates */
func indexHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := dbGetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Categories []dbData.Category
	}{
		Categories: categories,
	}

	tmpl := generateTemplate("index.html", []string{"templates/views/index.html", "templates/components/cat_navigation.html", "templates/components/register_form.html", "templates/components/login_form.html", "templates/components/navbar.html"})

	err = tmpl.Execute(w, data)
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

		// TODO: Handle verification of form fields and error messages properly

		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")

		// TODO: Add password confirmation field to registration form, check if passwords match

		if len(password) < 8 {
			http.Error(w, "Password must be at least 8 characters long", http.StatusBadRequest)
			return
		}

		if !credentials.ContainsLetter(password) || !credentials.ContainsDigit(password) || !credentials.ContainsSpecialChar(password) {
			http.Error(w, "Password must contain at least one letter, one digit, and one special character", http.StatusBadRequest)
			return
		}

		// TODO: Check if username is already taken and handle different possible scenarios, same for email

		if !credentials.IsValidEmail(email) {
			http.Error(w, "Invalid email address", http.StatusBadRequest)
			return
		}

		if !credentials.IsValidUsername(username) {
			http.Error(w, "Invalid username", http.StatusBadRequest)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Server error, unable to create your account.", 500)
			return
		}

		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum?parseTime=true")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		stmt, err := db.Prepare("INSERT INTO users (username, password, email, role_id) VALUES (?, ?, ?, ?)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		result, err := stmt.Exec(username, hashedPassword, email, 4)
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

	tmpl := generateTemplate("", []string{"templates/components/register_form.html"})
	tmpl.Execute(w, nil)
}

/* loginHandler handles the login form and redirects to the profile page */
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum?parseTime=true")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		var databaseUsername string
		var databasePassword string
		err = db.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&databaseUsername, &databasePassword)
		if err != nil {
			http.Redirect(w, r, "/error", http.StatusFound)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
		if err != nil {
			http.Redirect(w, r, "/error", http.StatusFound)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/profile?username=%s", username), http.StatusFound)
		return
	}

	tmpl := generateTemplate("", []string{"templates/components/login_form.html"})
	tmpl.Execute(w, nil)
}

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
		Categories []dbData.Category
	}{
		ID:         id,
		Name:       name,
		Categories: categories,
	}

	tmpl := generateTemplate("categories", []string{"templates/components/cat_navigation.html"})
	err = tmpl.ExecuteTemplate(w, "categories", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//TODO: Link categories and topics

/* topicsHandler handles the topics page */
func topicsHandler(w http.ResponseWriter, r *http.Request) {
	filters := dbData.RetrieveFilters(r)
	temp, err := dbData.GetTopics(filters)
	if err != nil {
		fmt.Println("Error in handlers.go")
		log.Fatal(err)
	}
	userData.Topics = temp.Topics
	userData.Filters = temp.Filters

	if r.Method == "POST" {
		tmpl := generateTemplate("", []string{"templates/components/topics-ctn.html"})
		tmpl.ExecuteTemplate(w, "topics-ctn", userData)
		return
	}

	tmpl := generateTemplate("topics.html", []string{"templates/views/topics.html", "templates/components/navbar.html", "templates/components/topics-ctn.html"})
	tmpl.Execute(w, userData)
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

	tmpl := generateTemplate("profile.html", []string{"templates/views/profile.html"})
	tmpl.Execute(w, data)
}

/* Temporary redirections for testing purposes */
func successHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := generateTemplate("", []string{"templates/results/success.html"})
	tmpl.Execute(w, nil)
}

/* Temporary redirections for testing purposes */
func errorHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := generateTemplate("", []string{"templates/results/error.html"})
	tmpl.Execute(w, nil)
}
