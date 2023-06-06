package main

import (
	"database/sql"
	"fmt"
	"forum/packages/credentials"
	"forum/packages/utils"
	"html/template"
	"net/http"
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
	tmpl = template.Must(tmpl.ParseFiles("templates/views/index.html", "templates/components/register_form.html", "templates/components/login_form.html"))

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
		/*	birthdate := r.FormValue("birthdate")*/

		if len(password) < 8 {
			http.Error(w, "Password must be at least 8 characters long", http.StatusBadRequest)
			return
		}

		if !credentials.ContainsLetter(password) || !credentials.ContainsDigit(password) || !credentials.ContainsSpecialChar(password) {
			http.Error(w, "Password must contain at least one letter, one digit, and one special character", http.StatusBadRequest)
			return
		}

		if !utils.IsValidEmail(email) {
			http.Error(w, "Invalid email address", http.StatusBadRequest)
			return
		}

		//file, _, err := r.FormFile("profile_picture")
		//if err != nil && err != http.ErrMissingFile {
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}
		//defer func() {
		//	if file != nil {
		//		file.Close()
		//	}
		//}()

		/*	dateofbirth, err := time.Parse("2006-01-02", birthdate)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}*/

		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum?parseTime=true")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		var stmt *sql.Stmt
		//if file != nil {
		//	stmt, err = db.Prepare("INSERT INTO users (username, password, email, profile_picture, /*birthdate,*/ role_id, isActive) VALUES (?, ?, ?, ?, /*?,*/ ?, ?)")
		//} else {
		stmt, err = db.Prepare("INSERT INTO users (username, password, email, /*birthdate,*/ role_id, isActive) VALUES (?, ?, ?, /*?,*/ ?, ?)")
		//}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		var result sql.Result
		//if file != nil {
		//	result, err = stmt.Exec(username, password, email, nil /*dateofbirth,*/, 3, 1)
		//} else {
		result, err = stmt.Exec(username, password, email /*dateofbirth,*/, 3, 1)
		//}
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

func successHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/results/success.html"))
	tmpl.Execute(w, nil)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/results/error.html"))
	tmpl.Execute(w, nil)
}
