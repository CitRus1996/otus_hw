package hw03frequencyanalysis

import (
	"fmt"
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
		if _, ok := dictionary[words[i]]; !ok {
			dictionary[words[i]] = 1
		} else {
			dictionary[words[i]]++
		}
	}

	countedWords := make([]Word, 0)
	for k, v := range dictionary {
		if len(k) > 0 {
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
	fmt.Println(countedWords)
	for i := 0; i < 10; i++ {
		top10 = append(top10, countedWords[i].text)
	}
	return top10
}
