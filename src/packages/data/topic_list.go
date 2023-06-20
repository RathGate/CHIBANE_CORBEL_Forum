package data

import (
	"database/sql"
	"fmt"
	"forum/packages/utils"
	"math"
	"net/http"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

// *STRUCTURES
type Topic struct {
	ID                  int64          `json:"id"`
	Title               string         `json:"title"`
	Category            string         `json:"category"`
	Content             string         `json:"content"`
	Username            sql.NullString `json:"username"`
	CreationDate        time.Time      `json:"creation_date"`
	Answers             int64          `json:"answer_count"`
	Score               int64          `json:"score"`
	Tags                []string       `json:"tags"`
	CurrentUserReaction sql.NullInt64  `json:"current_user_reaction"`
}

type TempTopic struct {
	ID          int64           `json:"id"`
	Title       string          `json:"title"`
	Category    string          `json:"category"`
	FirstPost   TempPost        `json:"first_post"`
	AnswerCount int64           `json:"answer_count"`
	Tags        []string        `json:"tags"`
	Answers     []TempPost      `json:"answers"`
	State       TempState       `json:"topic_state"`
	Permissions TempPermissions `json:"permissions"`
}
type TempPost struct {
	ID                  int64         `json:"id"`
	User                BaseUser      `json:"original_poster"`
	Content             string        `json:"content"`
	Timeline            Timeline      `json:"timeline"`
	Reactions           Reactions     `json:"reactions"`
	CurrentUserReaction sql.NullInt64 `json:"current_user_reaction"`
}

type Timeline struct {
	CreationDate     sql.NullTime
	ModificationDate sql.NullTime
}

type Reactions struct {
	Score    int   `json:"score"`
	Likes    int64 `json:"likes"`
	Dislikes int64 `json:"dislikes"`
}

type TempUser struct {
	ID       sql.NullInt64
	Username sql.NullString
	RoleID   sql.NullInt64
	Role     sql.NullString
}
type BaseUser struct {
	IsAuthenticated bool
	ID              int    `json:"id"`
	Username        string `json:"username"`
	RoleID          int    `json:"role_id"`
	Role            string `json:"role"`
	IsDeleted       bool   `json:"is_deleted"`
}

func (temp *TempUser) GetValidValues() (user BaseUser) {
	if !temp.ID.Valid || !temp.Username.Valid || !temp.RoleID.Valid {
		return BaseUser{IsDeleted: true}
	}
	user = BaseUser{
		ID:       int(temp.ID.Int64),
		Username: temp.Username.String,
		RoleID:   int(temp.RoleID.Int64),
		Role:     temp.Role.String,
	}

	return user
}

type TempState struct {
	IsClosed   int64 `json:"is_closed"`
	IsArchived int64 `json:"is_archived"`
	IsPinned   int64 `json:"is_pinned"`
}

type TempPermissions struct {
	MinReadRole  int64 `json:"min_read_role"`
	MinWriteRole int64 `json:"min_write_role"`
}

type Post struct {
	ID int64 `json:"id"`
}
type TopicFilters struct {
	OrderBy     string `json:"orderBy"`
	TimePeriod  int    `json:"timePeriod"`
	CurrentPage int    `json:"currentPage"`
	Limit       int    `json:"limit"`
	CategoryID  int    `json:"category_id"`
	ApplyLimit  bool
	UserID      int `json:"user_id"`
	Results     struct {
		PageCount   int `json:"pageCount"`
		ResultCount int `json:"resultCount"`
	} `json:"result"`
}

type Tag struct {
	Name  string
	Color string
}

type TopicData struct {
	Topics  []Topic
	Filters TopicFilters
}

// *DEFAULT FILTER VALUES
var DefaultTopicFilters = TopicFilters{
	OrderBy:     "newest",
	TimePeriod:  -1,
	CurrentPage: 1,
	Limit:       10,
	ApplyLimit:  true,
}

// *FUNCTIONS AND METHODS

// Retrieves the filters from the URL if GET request, and from form if POST request.
// Note: if a specific filter has no value, DefaultTopicFilters values are used.
func RetrieveFilters(r *http.Request) (result TopicFilters) {
	var tempDate string
	var tempCat string
	if r.Method == "POST" {
		result.OrderBy = r.FormValue("order")
		result.CurrentPage = utils.GetIntFromString(r.FormValue("page"))
		result.Limit = utils.GetIntFromString(r.FormValue("limit"))
		tempDate = r.FormValue("timePeriod")
		tempCat = r.FormValue("category")
	} else {
		result.OrderBy = r.URL.Query().Get("order")
		result.CurrentPage = utils.GetIntFromString(r.URL.Query().Get("page"))
		result.Limit = utils.GetIntFromString(r.URL.Query().Get("results"))
		tempDate = r.URL.Query().Get("date")
		tempCat = r.URL.Query().Get("category")
	}
	if tempDate == "all" {
		result.TimePeriod = -1
	} else {
		result.TimePeriod = utils.GetIntFromString(tempDate)
	}
	if tempCat == "all" {
		result.CategoryID = -1
	} else {
		result.CategoryID = utils.GetIntFromString(tempCat)
	}
	result.ApplyLimit = true
	result.CorrectFilters()
	return result
}

// Checks if all filter values are correct.
// If not correct or not specified, uses DefaultTopicFilters Values instead.
func (t *TopicFilters) CorrectFilters() {
	if !slices.Contains([]int{-1, 1, 7, 15, 30}, t.TimePeriod) {
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

// Retrieves all base data from topics given specific filters, and their total count.
func GetTopicListData(filters TopicFilters) (TopicData, error) {
	var tempTags sql.NullString
	data := TopicData{Filters: filters}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum?parseTime=true")
	if err != nil {
		return data, err
	}
	defer db.Close()

	// ?Gets the total number of rows, regardless of limit and offset filters
	row := db.QueryRow(QueryTopicCount(data.Filters))

	if err := row.Scan(&data.Filters.Results.ResultCount); err != nil {
		return data, err
	}

	// Corrects the original filter data if they are not relevant compared
	// to results (for example, user requested page 20 but there are only
	// 2 pages of results).
	data.Filters.Results.PageCount = int(math.Ceil(float64(data.Filters.Results.ResultCount) / float64(data.Filters.Limit)))

	if data.Filters.CurrentPage > data.Filters.Results.PageCount {
		data.Filters.CurrentPage = data.Filters.Results.PageCount
	}
	if data.Filters.CurrentPage < 1 {
		data.Filters.CurrentPage = 1
	}

	// ?Gets the base data for all topics with all filters applied, including limit and offset
	rows, _ := db.Query(QueryTopicsData(data.Filters))
	if err != nil {
		return data, err
	}

	// Iterates over rows to store data
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

// *QUERIES TO DATABASE

// Queries topic list base data from data base.
func QueryTopicsData(t TopicFilters) string {
	var orderByValues = map[string]string{
		"score":  "ORDER BY score DESC",
		"newest": "ORDER BY p.creation_date DESC",
		"oldest": "ORDER BY p.creation_date",
	}
	var stringBuilder = []string{`SELECT t.id, c.name, t.title, p.content, u.username, GROUP_CONCAT(DISTINCT tags.name SEPARATOR ";") as "tags", p.creation_date,`,
		`(SELECT COUNT(p.id) from posts as p WHERE p.id != tfp.post_id AND p.topic_id = t.id) as "answers",`,
		`(SELECT COUNT(pr.post_id) from post_reactions as pr where pr.post_id = p.id and pr.reaction_id = 1) -`,
		`(SELECT COUNT(pr.post_id) from post_reactions as pr where pr.post_id = p.id and pr.reaction_id = 2) as "score",`,
		fmt.Sprintf(`(SELECT pr.reaction_id from post_reactions as pr WHERE pr.post_id = p.id AND pr.user_id = %d) as "current_user_reaction"`, t.UserID),
		`FROM topic_first_posts AS tfp`,
		`JOIN topics AS t ON tfp.topic_id = t.id`,
		`JOIN posts AS p ON tfp.topic_id = p.topic_id`,
		`LEFT JOIN post_reactions AS pr ON pr.post_id = p.id`,
		`LEFT JOIN topic_tags AS tag ON t.id = tag.topic_id`,
		`LEFT JOIN tags ON tag.tag_id = tags.id`,
		`LEFT JOIN users AS u ON p.user_id = u.id`,
		`JOIN categories AS c ON c.id = t.category_id`}

	// WHERE p.creation_date >= DATE_SUB(SYSDATE(), INTERVAL %d DAY)
	if t.TimePeriod > 0 {
		stringBuilder = append(stringBuilder, fmt.Sprintf("WHERE p.creation_date >= DATE_SUB(SYSDATE(), INTERVAL %d DAY)", t.TimePeriod))
		if t.CategoryID > 0 {
			stringBuilder = append(stringBuilder, fmt.Sprintf("AND t.category_id = %d", t.CategoryID))
		}
	} else {
		if t.CategoryID > 0 {
			stringBuilder = append(stringBuilder, fmt.Sprintf("WHERE t.category_id = %d", t.CategoryID))
		}
	}
	// GROUP BY p.topic_id
	stringBuilder = append(stringBuilder, "GROUP BY p.topic_id")
	// ORDER BY %s DESC/ASC
	stringBuilder = append(stringBuilder, orderByValues[t.OrderBy])

	// LIMIT %d OFFSET %d
	if t.ApplyLimit {
		stringBuilder = append(stringBuilder, fmt.Sprintf("LIMIT %d OFFSET %d", t.Limit, t.Limit*(t.CurrentPage-1)))
	}

	return strings.Join(stringBuilder, "\n")
}

// Queries total topic count from database
func QueryTopicCount(t TopicFilters) string {
	var stringBuilder = []string{`SELECT COUNT(*) FROM topics AS t`,
		`JOIN topic_first_posts AS tfp ON t.id = tfp.topic_id`,
		`JOIN posts AS p ON p.id = tfp.post_id`}
	if t.TimePeriod > 0 {
		stringBuilder = append(stringBuilder, fmt.Sprintf("WHERE p.creation_date >= DATE_SUB(SYSDATE(), INTERVAL %d DAY)", t.TimePeriod))
		if t.CategoryID > 0 {
			stringBuilder = append(stringBuilder, fmt.Sprintf("AND t.category_id = %d", t.CategoryID))
		}
	} else {
		if t.CategoryID > 0 {
			stringBuilder = append(stringBuilder, fmt.Sprintf("WHERE t.category_id = %d", t.CategoryID))
		}
	}
	return strings.Join(stringBuilder, "\n")
}
