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
	"strconv"

	"github.com/gorilla/mux"
)

func generateTemplate(templateName string, filepaths []string) *template.Template {
	tmpl, err := template.New(templateName).Funcs(template.FuncMap{
		"getTimeSincePosted":  utils.GetTimeSincePosted,
		"getPagesArr":         utils.GetPagesArr,
		"GetPaginationValues": utils.GetPaginationValues,
		"getAllowedRoles":     data.GetAllowedRoles,
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
	tData.Categories, _ = data.GetCategories(DATABASE_ACCESS)
	tData.TopTrainers, _ = data.QueryTopTrainers(DATABASE_ACCESS, tData.User.ID)

	tmpl := generateTemplate("base.html", []string{"templates/base.html", "templates/views/index.html", "templates/components/header.html", "templates/components/topic_list.html", "templates/components/pagination.html", "templates/components/column_nav.html", "templates/components/popup_register.html", "templates/components/popup_login.html", "templates/components/column_ads.html", "templates/components/footer.html", "templates/components/cat_display.html", "templates/components/latest_news.html", "templates/components/new_topic.html"})

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

		formValidation, lastInsertedID := credentials.RegisterNewUser(DATABASE_ACCESS, username, password, email)
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
	tData.PageTitle = "Topics"
	tData.Categories, _ = data.GetCategories(DATABASE_ACCESS)
	tData.TopTrainers, _ = data.QueryTopTrainers(DATABASE_ACCESS, tData.User.ID)

	filters := data.RetrieveFilters(r)
	filters.UserID = tData.User.ID

	temp, err := data.TempQuery(DATABASE_ACCESS, filters)
	if err != nil {
		log.Fatal(err)
	}
	tData.Topics = temp.Topics
	tData.Filters = temp.Filters
	if r.Method == "POST" {
		r.ParseForm()
		tmpl := generateTemplate("", []string{"templates/components/topic_list.html", "templates/components/pagination.html"})
		tmpl.ExecuteTemplate(w, "topic_list", tData)
		return
	}

	tmpl := generateTemplate("base.html", []string{"templates/base.html", "templates/views/topics.html", "templates/components/header.html", "templates/components/topic_list.html", "templates/components/pagination.html", "templates/components/column_nav.html", "templates/components/popup_register.html", "templates/components/popup_login.html", "templates/components/new_topic.html", "templates/components/column_ads.html", "templates/components/footer.html"})
	tmpl.Execute(w, tData)

}

/* loginHandler handles the login form and redirects to the profile page */
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var formValidation credentials.FormValidation
		var userID int
		username := r.FormValue("username")
		password := r.FormValue("password")

		formValidation, userID = credentials.CheckUserCredentials(DATABASE_ACCESS, username, password)
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

func topicHandler(w http.ResponseWriter, r *http.Request) {
	tData := getSession(r)
	tData.PageTitle = "Topic"
	id := mux.Vars(r)["id"]

	// Checks if [id] parameter is a valid parameter
	topicID, err := strconv.Atoi(id)
	if err != nil {
		notFoundHandler(w, r)
		return
	}
	// Checks if a topic with this id exists
	if exists, err := data.TopicExists(DATABASE_ACCESS, topicID); !exists || err != nil {
		notFoundHandler(w, r)
	}

	// Reload template if user clicks on another page
	if r.Method == "POST" {
		// TODO
		fmt.Println("Soon")
		// Return
	}

	// Loads categories for left nav
	tData.Categories, _ = data.GetCategories(DATABASE_ACCESS)
	tData.TopTrainers, _ = data.QueryTopTrainers(DATABASE_ACCESS, tData.User.ID)

	tData.Topic, err = data.QuerySingleTopicData(DATABASE_ACCESS, topicID, tData.User.ID)
	if err != nil {
		log.Fatal(err)
	}

	tmpl := generateTemplate("base.html", []string{"templates/base.html", "templates/views/topic_view.html", "templates/components/header.html", "templates/components/topic_list.html", "templates/components/pagination.html", "templates/components/column_nav.html", "templates/components/popup_register.html", "templates/components/popup_login.html", "templates/components/new_topic.html", "templates/components/column_ads.html", "templates/components/footer.html"})
	tmpl.Execute(w, tData)
}

func newTopicHandler(w http.ResponseWriter, r *http.Request) {
	tData := getSession(r)
	tData.PageTitle = "New Topic"

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Extracts form values
		categoryID, err := strconv.Atoi(r.FormValue("category"))
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}
		title := r.FormValue("title")
		content := r.FormValue("content")
		userID := tData.User.ID

		// Creates the new topic in the database
		topicID, err := data.QueryNewTopic(categoryID, title, userID, content, 0)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirects to the newly created topic
		http.Redirect(w, r, "/topic/"+strconv.FormatInt(topicID, 10), http.StatusSeeOther)
		return
	}

	// Loads categories for select input
	tData.Categories, _ = data.GetCategories()

	tmpl := generateTemplate("base.html", []string{"templates/base.html", "templates/views/new_topic.html", "templates/components/header.html", "templates/components/topic_list.html", "templates/components/pagination.html", "templates/components/column_nav.html", "templates/components/popup_register.html", "templates/components/popup_login.html", "templates/components/column_ads.html", "templates/components/footer.html"})
	tmpl.Execute(w, tData)
}
