package dbData

type Data struct {
	IsLoggedIn bool
	PageTitle  string
	Topics     []Topic
	Filters    TopicFilters
	Categories []Category
}
