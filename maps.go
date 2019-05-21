package main

import (
	"strings"
)

func WordCount(s string) map[string]int {
	split := strings.Fields(s)
	wordMap := make(map[string]int)
	for _, v := range split {
		count, ok := wordMap[v]
		if !ok {
			wordMap[v] = 1
		} else {
			wordMap[v] = count + 1
		}
	}
	return wordMap
}

/*func main() {
	wc.Test(WordCount)
}*/
