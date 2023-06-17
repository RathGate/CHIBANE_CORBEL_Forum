package data

import (
	"fmt"
	"strings"
)

func WriteAllTopicsRequest(t TopicFilters) string {
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

func WriteShortTopicRequest(t TopicFilters) string {
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
