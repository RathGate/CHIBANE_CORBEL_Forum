package dbData

import (
	"database/sql"
	"fmt"
	"forum/packages/utils"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

type Topic struct {
	ID                  int64          `json:"id"`
	Title               string         `json:"title"`
	Category            string         `json:"category"`
	Content             string         `json:"content"`
	Username            sql.NullString `json:"username"`
	CreationDate        time.Time      `json:"creation_date"`
	Answers             int64          `json:"answers"`
	Score               int64          `json:"score"`
	Tags                []string       `json:"tags"`
	CurrentUserReaction sql.NullInt64  `json:"current_user_reaction"`
}

type TopicFilters struct {
	OrderBy     string `json:"orderBy"`
	TimePeriod  int    `json:"timePeriod"`
	CurrentPage int    `json:"currentPage"`
	Limit       int    `json:"limit"`
	ApplyLimit  bool
	Results     struct {
		PageCount   int `json:"pageCount"`
		ResultCount int `json:"resultCount"`
	} `json:"result"`
}

var DefaultTopicFilters = TopicFilters{
	OrderBy:     "score",
	TimePeriod:  -1,
	CurrentPage: 1,
	Limit:       10,
	ApplyLimit:  true,
}

func RetrieveFilters(r *http.Request) (result TopicFilters) {
	if r.Method == "POST" {
		result.OrderBy = r.FormValue("order")
		result.TimePeriod = getIntFromString(r.FormValue("timePeriod"))
		result.CurrentPage = getIntFromString(r.FormValue("currentPage"))
		result.Limit = getIntFromString(r.FormValue("limit"))
	} else {
		result.OrderBy = r.URL.Query().Get("order")
		result.TimePeriod = getIntFromString(r.URL.Query().Get("date"))
		result.CurrentPage = getIntFromString(r.URL.Query().Get("page"))
		result.Limit = getIntFromString(r.URL.Query().Get("results"))
	}
	result.ApplyLimit = true

	result.CorrectFilters()
	return result
}

func (t *TopicFilters) CorrectFilters() {
	if !slices.Contains([]int{1, 7, 15, 30}, t.TimePeriod) {
		t.TimePeriod = DefaultTopicFilters.TimePeriod
	}
	if !slices.Contains([]int{5, 10, 25, 50}, t.Limit) {
		t.Limit = DefaultTopicFilters.Limit
	}
	if !slices.Contains([]string{"score", "oldest", "newest"}, t.OrderBy) {
		t.OrderBy = DefaultTopicFilters.OrderBy
	}
	if t.CurrentPage < 1 {
		t.CurrentPage = DefaultTopicFilters.CurrentPage
	}
	t.ApplyLimit = true
}

func getIntFromString(str string) int {
	if value, err := strconv.Atoi(str); err != nil || value < 1 {
		return 0
	} else {
		return value
	}
}

type TopicData struct {
	Topics  []Topic
	Filters TopicFilters
}

func GetTimeSincePosted(topic Topic) string {
	timeValues := utils.GetDeltaValues(topic.CreationDate, time.Now())
	timeNames := []string{"yr", "mo", "d", "hr", "min", "sec"}
	for i, value := range timeValues {
		if value > 0 {
			return fmt.Sprintf("%02d%s ago", value, timeNames[i])
		}
	}
	return "now"
}

func GetTopics(filters TopicFilters) (TopicData, error) {
	// Tags are stored in a "tag1;tag2;tag3" way.
	// Require to be splitted when received from the database.
	var tempTags sql.NullString
	data := TopicData{Filters: filters}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum?parseTime=true")
	if err != nil {
		return data, err
	}
	defer db.Close()

	// ? Gets the total number of rows
	row := db.QueryRow(WriteShortTopicRequest(data.Filters))

	if err := row.Scan(&data.Filters.Results.ResultCount); err != nil {
		return data, err
	}
	data.Filters.Results.PageCount = int(math.Ceil(float64(data.Filters.Results.ResultCount) / float64(data.Filters.Limit)))

	if data.Filters.CurrentPage > data.Filters.Results.PageCount {
		data.Filters.CurrentPage = data.Filters.Results.PageCount
	}
	if data.Filters.CurrentPage < 1 {
		data.Filters.CurrentPage = 1
	}
	fmt.Println(WriteAllTopicsRequest(data.Filters))
	rows, _ := db.Query(WriteAllTopicsRequest(data.Filters))
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		tempTopic := new(Topic)
		err := rows.Scan(&tempTopic.ID, &tempTopic.Category, &tempTopic.Title, &tempTopic.Content, &tempTopic.Username, &tempTags, &tempTopic.CreationDate, &tempTopic.Answers, &tempTopic.Score, &tempTopic.CurrentUserReaction)
		if err != nil {
			return data, err
		}

		// Splits tags if any in result row:
		if tempTags.Valid {
			tempTopic.Tags = strings.Split(tempTags.String, ";")
		}
		// Converts corrupted date to UTC:
		utils.TimeToUTC(&tempTopic.CreationDate)

		data.Topics = append(data.Topics, *tempTopic)
	}
	return data, nil
}
