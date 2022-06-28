package hw03frequencyanalysis

import (
	"regexp"
	"sort"
)

type Word struct {
	text  string
	count int
}

var re = regexp.MustCompile(`\s`)

func Top10(text string) []string {
	top10 := make([]string, 0)
	// Place your code here.
	if len(text) == 0 {
		return top10
	}
	words := re.Split(text, -1)
	dictionary := make(map[string]int)
	for i := range words {
		dictionary[words[i]]++
	}

	countedWords := make([]Word, 0, len(dictionary))
	for k, v := range dictionary {
		if k != "" {
			countedWords = append(countedWords, Word{text: k, count: v})
		}
	}
	sort.Slice(countedWords, func(i, j int) bool {
		if countedWords[i].count > countedWords[j].count {
			return true
		} else if countedWords[i].count == countedWords[j].count {
			return countedWords[i].text < countedWords[j].text
		}
		return false
	})
	for i := 0; i < 10; i++ {
		top10 = append(top10, countedWords[i].text)
	}
	return top10
}
