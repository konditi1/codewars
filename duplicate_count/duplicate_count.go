package kata

import "strings"

func Duplicate_count(s1 string) int {
	// convert all caps to lower
	s2 := strings.ToLower(s1)
	count := 0

	sMap := make(map[string]int)

	for _, val := range s2 {
		sMap[string(val)]++
	}

	for _, value := range sMap {
		if value > 1 {
			count++
		}
	}
	return count
}
