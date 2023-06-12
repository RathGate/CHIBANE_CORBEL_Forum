package main

import (
	"database/sql"
	"fmt"
	"forum/packages/credentials"
	"forum/packages/dbData"
	"html/template"
	"log"
	"net/http"
)

func generateTemplate(templateName string, filepaths []string) *template.Template {
	tmpl, err := template.New(templateName).Funcs(template.FuncMap{
		"getTimeSincePosted": dbData.GetTimeSincePosted,
		"getPagesArr":        dbData.GetPagesArr,
		"getPagesValues":     dbData.GetPagesValues,
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
	// TEMPORARY:
	http.Redirect(w, r, "/topics", http.StatusSeeOther)

	// categories, err := dbGetCategories()
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// data := struct {
	// 	Categories []dbData.Category
	// }{
	// 	Categories: categories,
	// }

	// // Generates template:
	// tmpl := template.New("index.html")

	// // Parse the templates:
	// tmpl = template.Must(tmpl.ParseFiles("templates/views/index.html", "templates/components/cat_navigation.html", "templates/components/register_form.html", "templates/components/login_form.html", "templates/components/navbar.html"))

	// // Execute the templates
	// err = tmpl.ExecuteTemplate(w, "index.html", data)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
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
		stmt, err = db.Prepare("INSERT INTO users (username, password, email, role_id) VALUES (?, ?, ?, ?)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		var result sql.Result
		result, err = stmt.Exec(username, password, email, 3)
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

func topicsHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := dbGetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userData.Categories = categories

	filters := dbData.RetrieveFilters(r)
	temp, err := dbData.GetTopics(filters)
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
		Categories []dbData.Category
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
