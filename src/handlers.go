package main

import (
	"database/sql"
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
	w.WriteHeader(http.StatusNotFound)

	tmpl := generateTemplate("404.html", []string{"templates/404.html"})
	tmpl.Execute(w, nil)
}
func notAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)

	tmpl := generateTemplate("403.html", []string{"templates/403.html"})
	tmpl.Execute(w, nil)
}

/* indexHandler handles the index page, parses most of the templates and executes them */
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tData := getSession(r)
	tData.PageTitle = "Home"
	tData.Categories, _ = data.GetCategories(DATABASE_ACCESS, tData.User.ID)

	tData.TopTrainers, _ = data.QueryTopTrainers(DATABASE_ACCESS, tData.User.ID)

	tmpl := generateTemplate("base.html", []string{"templates/base.html", "templates/components/mobile-menus.html", "templates/views/index.html", "templates/components/header.html", "templates/components/topic_list.html", "templates/components/pagination.html", "templates/components/column_nav.html", "templates/components/popup_register.html", "templates/components/popup_login.html", "templates/components/column_ads.html", "templates/components/footer.html", "templates/components/cat_display.html", "templates/components/latest_news.html", "templates/components/new_topic.html"})

	err := tmpl.Execute(w, tData)
	fmt.Println(err)
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
			setSession(r, &w, lastInsertedID)
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
	tData.Categories, _ = data.GetCategories(DATABASE_ACCESS, tData.User.ID)
	tData.TopTrainers, _ = data.QueryTopTrainers(DATABASE_ACCESS, tData.User.ID)
	filters := data.RetrieveFilters(r, tData.User.IsAuthenticated)

	if filters.CategoryID != 0 {
		cat, err := data.GetCategoryData(DATABASE_ACCESS, filters.CategoryID)
		if (err == sql.ErrNoRows && filters.CategoryID != 0) || cat.MinReadRole < int64(tData.User.RoleID) {
			http.Redirect(w, r, "/topics", http.StatusSeeOther)
		}
	}

	filters.UserID = tData.User.ID

	temp, err := data.TempQuery(DATABASE_ACCESS, filters, tData.User.RoleID)
	if err != nil {
		log.Fatal(err)
	}
	tData.Topics = temp.Topics
	tData.Filters = temp.Filters
	if r.Method == "POST" {
		r.ParseForm()
		tmpl := generateTemplate("", []string{"templates/components/topic_list.html", "templates/components/pagination.html", "templates/components/noresult.html"})
		err := tmpl.ExecuteTemplate(w, "topic_list", tData)
		fmt.Println(err)
		return
	}

	tmpl := generateTemplate("base.html", []string{"templates/base.html", "templates/views/topics.html", "templates/components/header.html", "templates/components/topic_list.html", "templates/components/pagination.html",
		"templates/components/column_nav.html", "templates/components/popup_register.html", "templates/components/popup_login.html", "templates/components/new_topic.html", "templates/components/column_ads.html",
		"templates/components/footer.html", "templates/components/mobile-menus.html", "templates/components/noresult.html"})
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
	if topic, err := data.GetBaseTopicData(DATABASE_ACCESS, topicID); err != nil {
		notFoundHandler(w, r)
		return
	} else if !data.CheckReadPermission(topic, tData.User.RoleID) {
		notAllowedHandler(w, r)
		return
	}

	// Reload template if user clicks on another page
	if r.Method == "POST" {
		r.ParseForm()
		content := r.FormValue("content")
		data := data.AddAnswerToTopic(DATABASE_ACCESS, topicID, tData.User.ID, tData.User.RoleID, content)
		jsonValues, _ := json.Marshal(data)
		w.Write(jsonValues)
		return
	}

	// Loads categories for left nav
	tData.Categories, _ = data.GetCategories(DATABASE_ACCESS, tData.User.ID)

	tData.TopTrainers, _ = data.QueryTopTrainers(DATABASE_ACCESS, tData.User.ID)

	tData.Topic, err = data.QuerySingleTopicData(DATABASE_ACCESS, topicID, tData.User.ID)
	if err != nil {
		log.Fatal(err)
	}

	tmpl := generateTemplate("base.html", []string{"templates/base.html", "templates/views/topic_view.html", "templates/components/header.html", "templates/components/topic_list.html", "templates/components/pagination.html", "templates/components/column_nav.html", "templates/components/popup_register.html", "templates/components/popup_login.html", "templates/components/new_topic.html", "templates/components/column_ads.html", "templates/components/footer.html", "templates/components/mobile-menus.html"})
	tmpl.Execute(w, tData)
}

func newTopicHandler(w http.ResponseWriter, r *http.Request) {
	tData := getSession(r)
	tData.PageTitle = "New Topic"
	tData.Categories, _ = data.GetCategories(DATABASE_ACCESS, tData.User.ID)
	tData.TopTrainers, _ = data.QueryTopTrainers(DATABASE_ACCESS, tData.User.ID)

	if r.Method == "POST" {
		r.ParseForm()
		tempCat := r.FormValue("category_id")
		categoryID, err := strconv.Atoi(tempCat)
		if err != nil {
			jsonValues, _ := json.Marshal(data.AnswerValidation{
				Status: http.StatusBadRequest,
				Error:  "You must choose a valid category",
			})
			w.Write(jsonValues)
			return
		}
		title := r.FormValue("title")
		content := r.FormValue("content")
		jsonValues, _ := json.Marshal(data.CreateNewTopic(DATABASE_ACCESS, tData.User.ID, tData.User.RoleID, categoryID, title, content))

		w.Write(jsonValues)
		return
	}

	if !tData.User.IsAuthenticated {
		notAllowedHandler(w, r)
		return
	}

	tmpl := generateTemplate("base.html", []string{"templates/base.html", "templates/views/new_topic.html", "templates/components/header.html", "templates/components/topic_list.html", "templates/components/pagination.html", "templates/components/column_nav.html", "templates/components/popup_register.html", "templates/components/popup_login.html", "templates/components/column_ads.html", "templates/components/footer.html", "templates/components/mobile-menus.html", "templates/components/create_topic.html"})
	err := tmpl.Execute(w, tData)
	fmt.Println(err)
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	tData := getSession(r)
	tData.PageTitle = "New Topic"

	if tData.User.RoleID < 0 || tData.User.RoleID > 2 {
		notAllowedHandler(w, r)
		return
	}

	tmpl := generateTemplate("base.html", []string{"templates/base.html", "templates/views/topic_view.html", "templates/components/header.html", "templates/components/topic_list.html", "templates/components/pagination.html", "templates/components/column_nav.html", "templates/components/popup_register.html", "templates/components/popup_login.html", "templates/components/column_ads.html", "templates/components/footer.html", "templates/components/mobile-menus.html"})
	tmpl.Execute(w, tData)
}
func editProfileHandler(w http.ResponseWriter, r *http.Request) {
	tData := getSession(r)
	tData.PageTitle = "New Topic"

	tmpl := generateTemplate("base.html", []string{"templates/base.html", "templates/views/topic_view.html", "templates/components/header.html", "templates/components/topic_list.html", "templates/components/pagination.html", "templates/components/column_nav.html", "templates/components/popup_register.html", "templates/components/popup_login.html", "templates/components/column_ads.html", "templates/components/footer.html", "templates/components/mobile-menus.html"})
	tmpl.Execute(w, tData)
}
func userHandler(w http.ResponseWriter, r *http.Request) {
	tData := getSession(r)
	tData.PageTitle = "New Topic"

	tmpl := generateTemplate("base.html", []string{"templates/base.html", "templates/views/topic_view.html", "templates/components/header.html", "templates/components/topic_list.html", "templates/components/pagination.html", "templates/components/column_nav.html", "templates/components/popup_register.html", "templates/components/popup_login.html", "templates/components/column_ads.html", "templates/components/footer.html", "templates/components/mobile-menus.html"})
	tmpl.Execute(w, tData)
}

func privacyHandler(w http.ResponseWriter, r *http.Request) {
	tData := getSession(r)
	tData.Categories, _ = data.GetCategories(DATABASE_ACCESS, tData.User.ID)

	tData.TopTrainers, _ = data.QueryTopTrainers(DATABASE_ACCESS, tData.User.ID)
	tData.PageTitle = "Privacy Policy"

	tmpl := generateTemplate("base.html", []string{"templates/base.html", "templates/components/header.html", "templates/views/privacy.html", "templates/components/pagination.html", "templates/components/column_nav.html", "templates/components/popup_register.html", "templates/components/popup_login.html", "templates/components/column_ads.html", "templates/components/footer.html", "templates/components/mobile-menus.html"})
	tmpl.Execute(w, tData)
}

func tosHandler(w http.ResponseWriter, r *http.Request) {
	tData := getSession(r)
	tData.Categories, _ = data.GetCategories(DATABASE_ACCESS, tData.User.ID)

	tData.TopTrainers, _ = data.QueryTopTrainers(DATABASE_ACCESS, tData.User.ID)
	tData.PageTitle = "Forum Guidelines"

	tmpl := generateTemplate("base.html", []string{"templates/base.html", "templates/components/header.html", "templates/views/tos.html", "templates/components/pagination.html", "templates/components/column_nav.html", "templates/components/popup_register.html", "templates/components/popup_login.html", "templates/components/column_ads.html", "templates/components/footer.html", "templates/components/mobile-menus.html"})
	tmpl.Execute(w, tData)
}
