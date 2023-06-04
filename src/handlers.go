package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"time"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	// *Generates and executes templates:
	tmpl := template.Must(template.ParseFiles("templates/views/index.html"))
	tmpl.Execute(w, nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Generates and executes templates:
	tmpl := template.New("index.html")

	// Parse the templates
	tmpl = template.Must(tmpl.ParseFiles("templates/views/index.html", "templates/components/register_form.html"))

	// Execute the template
	tmpl.ExecuteTemplate(w, "index.html", nil)
}
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
		birthdate := r.FormValue("birthdate")

		file, _, err := r.FormFile("profile_picture")
		if err != nil && err != http.ErrMissingFile {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer func() {
			if file != nil {
				file.Close()
			}
		}()

		dateofbirth, err := time.Parse("2006-01-02", birthdate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum?parseTime=true")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		var stmt *sql.Stmt
		if file != nil {
			stmt, err = db.Prepare("INSERT INTO users (username, password, email, profile_picture, birthdate, role_id, isActive) VALUES (?, ?, ?, ?, ?, ?, ?)")
		} else {
			stmt, err = db.Prepare("INSERT INTO users (username, password, email, birthdate, role_id, isActive) VALUES (?, ?, ?, ?, ?, ?)")
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		var result sql.Result
		if file != nil {
			result, err = stmt.Exec(username, password, email, nil, dateofbirth, 3, 1)
		} else {
			result, err = stmt.Exec(username, password, email, dateofbirth, 3, 1)
		}
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

	tmpl := template.Must(template.ParseFiles("templates/views/register_form.html"))
	tmpl.Execute(w, nil)
}

func successHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/error_handling/success.html"))
	tmpl.Execute(w, nil)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/error_handling/error.html"))
	tmpl.Execute(w, nil)
}
