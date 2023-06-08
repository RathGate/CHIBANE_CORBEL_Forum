package user

import "forum/packages/dbData"

type Data struct {
	IsLoggedIn bool
	PageTitle  string
	Topics     []dbData.Topic
	Filters    dbData.TopicFilters
}
