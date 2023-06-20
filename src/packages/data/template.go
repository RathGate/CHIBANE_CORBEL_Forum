package data

type TemplateData struct {
	PageTitle  string
	User       BaseUser
	Topics     []Topic
	Filters    TopicFilters
	Categories []Category
	// Ici
	Topic       TempTopic
	TopTrainers [6]TopTrainer
}

type TemplateUser struct {
	IsAuthenticated bool   `json:"is_authenticated"`
	ID              int    `json:"id"`
	Username        string `json:"username"`
}
