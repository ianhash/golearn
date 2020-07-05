package exercise

import (
	"strings"
)

// WordCount 统计数字
func WordCount(s string) map[string]int {
	wordCounts := make(map[string]int)
	words := strings.Fields(s)
	for _, w := range words {
		var _, has = wordCounts[w]
		if has {
			wordCounts[w]++
		} else {
			wordCounts[w] = 1
		}
	}
	return wordCounts
}
