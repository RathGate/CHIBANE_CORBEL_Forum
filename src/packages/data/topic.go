package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"forum/packages/utils"
	"net/http"
	"strings"
)

func CheckReadPermission(temptopic TempTopic, roleID int) bool {
	return temptopic.Permissions.MinReadRole >= int64(roleID) && temptopic.CatPermissions.MinReadRole >= int64(roleID)
}
func CheckWritePermission(temptopic TempTopic, roleID int) bool {
	fmt.Println(temptopic.Permissions.MinWriteRole, temptopic.CatPermissions.MinWriteRole, roleID)
	return temptopic.Permissions.MinWriteRole >= int64(roleID) && temptopic.CatPermissions.MinWriteRole >= int64(roleID)
}

func GetBaseTopicData(dba utils.DB_Access, id int) (tempTopic TempTopic, err error) {
	db, err := sql.Open("mysql", dba.ToString())
	if err != nil {
		return tempTopic, err
	}
	defer db.Close()

	err = db.QueryRow(`SELECT t.id, t.min_read_role, t.min_write_role, c.min_read_role, c.min_write_role FROM topics AS t JOIN categories AS c ON t.category_id = c.id WHERE t.id = ?`, id).Scan(&tempTopic.ID, &tempTopic.Permissions.MinReadRole,
		&tempTopic.Permissions.MinWriteRole, &tempTopic.CatPermissions.MinReadRole, &tempTopic.CatPermissions.MinWriteRole)
	if err != nil {
		return tempTopic, err
	}

	return tempTopic, nil
}

func QuerySingleTopicData(dba utils.DB_Access, topicID int, userID int) (data TempTopic, err error) {
	var tempTags sql.NullString
	var tempUser TempUser

	db, err := sql.Open("mysql", dba.ToString())
	if err != nil {
		return data, err
	}
	defer db.Close()

	// ?Get the base data of the topic (without the data of the first post)
	row := db.QueryRow(`SELECT t.id, c.name AS "category", t.title, GROUP_CONCAT(DISTINCT tags.name SEPARATOR ";") as "tags", t.is_closed, t.is_pinned, t.is_archived, t.min_read_role, t.min_write_role, c.min_read_role, c.min_write_role,
	(SELECT COUNT(p.id) from posts as p WHERE p.id != tfp.post_id AND p.topic_id = t.id) as "answer_count"
	FROM topics AS t
	JOIN topic_first_posts AS tfp ON tfp.topic_id = t.id
	JOIN categories AS c ON t.category_id = c.id
	JOIN posts AS p ON p.id = tfp.post_id
	LEFT JOIN topic_tags AS tag ON tfp.topic_id = tag.topic_id
	LEFT JOIN tags ON tag.tag_id = tags.id
	WHERE t.id = ?
    GROUP BY t.id`, topicID)
	row.Scan(&data.ID, &data.Category, &data.Title, &tempTags, &data.State.IsClosed, &data.State.IsPinned, &data.State.IsArchived,
		&data.Permissions.MinReadRole, &data.Permissions.MinWriteRole, &data.CatPermissions.MinReadRole, &data.CatPermissions.MinWriteRole, &data.AnswerCount)

	if tempTags.Valid {
		data.Tags = strings.Split(tempTags.String, ";")
	}

	// ?Get the data of the first post
	row = db.QueryRow(`SELECT p.id AS "post_id", u.id AS "user_id", u.username AS "username", u.role_id, r.name AS "role_name", p.content, p.creation_date, p.modification_date,
	(SELECT COUNT(pr.post_id) from post_reactions as pr where pr.post_id = p.id and pr.reaction_id = 1) AS "likes", 
	(SELECT COUNT(pr.post_id) from post_reactions as pr where pr.post_id = p.id and pr.reaction_id = 2) AS "dislikes",
	(SELECT pr.reaction_id from post_reactions as pr WHERE pr.post_id = p.id AND pr.user_id = ?) as "current_user_reaction"
	FROM posts AS p
	JOIN topic_first_posts AS tfp ON tfp.topic_id = p.topic_id
	JOIN users AS u on p.user_id = u.id
	JOIN roles AS r ON u.role_id = r.id
	WHERE p.topic_id = ?`, userID, topicID)
	row.Scan(&data.FirstPost.ID, &tempUser.ID, &tempUser.Username, &tempUser.RoleID, &tempUser.Role,
		&data.FirstPost.Content, &data.FirstPost.Timeline.CreationDate, &data.FirstPost.Timeline.ModificationDate, &data.FirstPost.Reactions.Likes, &data.FirstPost.Reactions.Dislikes, &data.FirstPost.CurrentUserReaction)
	data.FirstPost.User = tempUser.GetValidValues()
	data.FirstPost.Reactions.Score = int(data.FirstPost.Reactions.Likes) - int(data.FirstPost.Reactions.Dislikes)
	// ?Gets all the answers to the topic
	rows, err := db.Query(`SELECT p.id, u.id AS "user_id", u.username AS "username", r.id AS "role_id", r.name AS "role_name", p.content, p.creation_date, p.modification_date,
	(SELECT COUNT(pr.post_id) FROM post_reactions AS pr where pr.post_id = p.id and pr.reaction_id = 1) AS "likes", 
	(SELECT COUNT(pr.post_id) FROM post_reactions AS pr where pr.post_id = p.id and pr.reaction_id = 2) AS "dislikes",
	(SELECT pr.reaction_id FROM post_reactions AS pr WHERE pr.post_id = p.id AND pr.user_id = ?) AS "current_user_reaction"
	FROM posts AS p 
	LEFT JOIN users AS u ON p.user_id = u.id
	JOIN roles AS r ON r.id = u.role_id
	JOIN topic_first_posts AS tfp ON tfp.topic_id = p.topic_id
	WHERE p.id != tfp.post_id AND p.topic_id = ?`, userID, topicID)

	if err != nil {
		if err == sql.ErrNoRows {
			return data, nil
		}
		return TempTopic{}, err
	}

	for rows.Next() {
		tempPost := new(TempPost)
		tempUser = TempUser{}
		rows.Scan(&tempPost.ID, &tempUser.ID, &tempUser.Username, &tempUser.RoleID, &tempUser.Role,
			&tempPost.Content, &tempPost.Timeline.CreationDate, &tempPost.Timeline.ModificationDate, &tempPost.Reactions.Likes, &tempPost.Reactions.Dislikes, &tempPost.CurrentUserReaction)
		tempPost.User = tempUser.GetValidValues()
		tempPost.Reactions.Score = int(tempPost.Reactions.Likes) - int(data.FirstPost.Reactions.Dislikes)

		data.Answers = append(data.Answers, *tempPost)
	}
	return data, nil
}

func QueryNewTopic(categoryID int, title string, userID int, content string, tagID int) (int64, error) {
	var topicID int64
	var postID int64

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum?parseTime=true")
	if err != nil {
		return 0, err
	}
	defer db.Close()

	// Inserts new topic into topics table
	result, err := db.Exec("INSERT INTO topics (category_id, title) VALUES (?, ?)", categoryID, title)
	if err != nil {
		return 0, err
	}

	// Gets the ID of the newly created topic
	topicID, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Inserts new post into posts table
	result, err = db.Exec("INSERT INTO posts (topic_id, user_id, content) VALUES (?, ?, ?)", topicID, userID, content)
	if err != nil {
		return 0, err
	}

	// Gets the ID of the newly created post
	postID, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Links the topic and its first post
	_, err = db.Exec("INSERT INTO topic_first_posts (topic_id, post_id) VALUES (?, ?)", topicID, postID)
	if err != nil {
		return 0, err
	}

	// Add tags to the topic
	//if tagID > 0 {
	//	_, err = db.Exec("INSERT INTO topic_tags (topic_id, tag_id) VALUES (?, ?)", topicID, tagID)
	//	if err != nil {
	//		return 0, err
	//	}
	//}

	return topicID, nil
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "   ")
	return string(s)
}

func GetAllowedRoles(p TempPermissions) map[string][]int {
	result := map[string][]int{
		"read":  {},
		"write": {},
	}
	for i := 2; i <= int(p.MinReadRole) && i <= 4; i++ {
		result["read"] = append(result["read"], i)
	}
	for i := 2; i <= int(p.MinWriteRole) && i <= 4; i++ {
		result["write"] = append(result["write"], i)
	}
	return result
}

type AnswerValidation struct {
	Status int    `json:"status"`
	Error  string `json:"error_msg"`
}

func AddAnswerToTopic(dba utils.DB_Access, topicID, userID, roleID int, content string) AnswerValidation {
	topic, err := GetBaseTopicData(dba, topicID)
	if err != nil {
		if err != sql.ErrNoRows {
			return AnswerValidation{
				Status: http.StatusBadRequest,
				Error:  "An error occurred : the topic might have been deleted",
			}
		}
		return AnswerValidation{
			Status: http.StatusInternalServerError,
			Error:  "An error occurred : try again later",
		}
	}
	if !CheckWritePermission(topic, roleID) {
		return AnswerValidation{
			Status: http.StatusBadRequest,
			Error:  "An error occurred : write permissions missing",
		}
	}
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum?parseTime=true")
	if err != nil {
		return AnswerValidation{
			Status: http.StatusBadRequest,
			Error:  "An error occurred : try again later",
		}
	}
	defer db.Close()

	// Inserts new topic into topics table
	result, err := db.Exec("INSERT INTO posts (topic_id, user_id, content) VALUES (?, ?, ?)", topic.ID, userID, content)
	if err != nil {
		fmt.Println(err)
		return AnswerValidation{
			Status: http.StatusInternalServerError,
			Error:  "An error occurred : try again later",
		}
	}

	// Gets the ID of the newly created topic
	if _, err = result.LastInsertId(); err != nil {
		return AnswerValidation{
			Status: http.StatusInternalServerError,
			Error:  "An error occurred : try again later",
		}
	}

	return AnswerValidation{
		Status: http.StatusOK,
	}

}
