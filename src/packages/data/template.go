package data

type TemplateData struct {
	PageTitle  string
	User       TemplateUser
	Topics     []Topic
	Filters    TopicFilters
	Categories []Category
	// Ici
}

type TemplateUser struct {
	IsAuthenticated bool   `json:"is_authenticated"`
	ID              int    `json:"id"`
	Username        string `json:"username"`
}
