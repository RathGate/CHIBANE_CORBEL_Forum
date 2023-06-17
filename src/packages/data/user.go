package data

type Data struct {
	UserID     int
	PageTitle  string
	Topics     []Topic
	Filters    TopicFilters
	Categories []Category
	User       ShortUser
}

type ShortUser struct {
	IsAuthenticated bool   `json:"is_authenticated"`
	ID              int    `json:"id"`
	Username        string `json:"username"`
}
