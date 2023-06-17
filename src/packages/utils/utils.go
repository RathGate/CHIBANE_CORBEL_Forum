package utils

import (
	"fmt"
	"strconv"
	"time"
)

func GetDeltaValues(a, b time.Time) []int {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year := int(y2 - y1)
	month := int(M2 - M1)
	day := int(d2 - d1)
	hour := int(h2 - h1)
	min := int(m2 - m1)
	sec := int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return []int{year, month, day, hour, min, sec}
}

func TimeToUTC(baseTime *time.Time) {
	_, offset := time.Now().Zone()
	*baseTime = baseTime.Add(-time.Duration(offset) * time.Second)
}

func GetIntFromString(str string) int {
	if value, err := strconv.Atoi(str); err != nil || value < 1 {
		return 0
	} else {
		return value
	}
}

// Gets a readable version of the time elapsed since a given datetime.
// Format example: "8mo ago", "24sec ago"...
func GetTimeSincePosted(t time.Time) string {
	timeValues := GetDeltaValues(t, time.Now())
	timeNames := []string{"yr", "mo", "d", "hr", "min", "sec"}
	for i, value := range timeValues {
		if value > 0 {
			return fmt.Sprintf("%02d%s ago", value, timeNames[i])
		}
	}
	return "now"
}

func GetPagesArr(currPage, totalPages int) []int {
	result := []int{1}
	// No page or no result somehow
	if currPage == 0 || totalPages == 0 {
		return nil
	}
	// Number of total pages less or equal to 6
	if totalPages <= 7 {
		for i := 2; i <= totalPages-1; i++ {
			result = append(result, i)
		}
	} else if currPage <= 5 {
		for i := 2; i <= 5; i++ {
			result = append(result, i)
		}
		result = append(result, -1)
	} else if currPage <= totalPages-4 {
		result = append(result, -1)
		for i := totalPages - 4; i <= totalPages-1; i++ {
			result = append(result, i)
		}
	} else {
		result = append(result, -1)
		for i := currPage - 1; i <= currPage+1; i++ {
			result = append(result, i)
		}
		result = append(result, -1)
	}
	if totalPages > 1 {
		result = append(result, totalPages)
	}
	return result
}

func GetPaginationValues(currentPage, resultCount, limit int) map[string]int {
	result := map[string]int{"minBound": 0, "maxBound": 0}
	if resultCount == 0 {
		return result
	}
	result["minBound"] = 1 + (currentPage-1)*limit
	result["maxBound"] = result["minBound"] + 9
	if result["maxBound"] > resultCount {
		result["maxBound"] = resultCount
	}
	return result
}
