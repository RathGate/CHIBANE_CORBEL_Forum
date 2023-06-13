package dbData

type Data struct {
	IsLoggedIn bool
	Username   string
	UserID     int
	PageTitle  string
	Topics     []Topic
	Filters    TopicFilters
	Categories []Category
}
