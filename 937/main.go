package main

import (
	"fmt"
	"sort"
	"strings"
)

type Log struct {
	Digit   int
	Prefix  string
	Content string
}

func main() {
	logs := []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}
	fmt.Println(reorderLogFiles(logs))

	logs = []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"}
	fmt.Println(reorderLogFiles(logs))

}

// first create the Log struct for each log string
// compare digit - return the smaller one
// compare content - if they are not the same then return smaller one
// what about the content for digit?

// else just do digits in separate files and then loop again later?

func reorderLogFiles(logs []string) []string {

	letters := []Log{}
	digits := []string{}

	for _, log := range logs {
		logSlice := strings.Split(log, " ")
		if IsDigitsOnly(logSlice[1]) {
			digits = append(digits, log)
			continue
		}
		letters = append(letters, Log{
			Prefix:  logSlice[0],
			Content: strings.Join(logSlice[1:], " "),
		})
	}

	sort.Slice(letters, func(i, j int) bool {
		if letters[i].Content != letters[j].Content {
			return letters[i].Content < letters[j].Content
		}
		return letters[i].Prefix > letters[j].Prefix
	})

	logs = []string{}
	for _, letter := range letters {
		logs = append(logs, fmt.Sprintf("%s %s", letter.Prefix, letter.Content))
	}
	return append(logs, digits...)
}

// func reorderLogFiles(logs []string) []string {

// 	logSlice := []Log{}

// 	for _, log := range logs {
// 		split := strings.Split(log, " ")
// 		digit := 0
// 		if IsDigitsOnly(split[1]) {
// 			digit = 1
// 		}
// 		logSlice = append(logSlice, Log{
// 			Digit:   digit,
// 			Prefix:  split[0],
// 			Content: strings.Join(split[1:], " "),
// 		})
// 	}

// 	sort.Slice(logSlice, func(i, j int) bool {
// 		if logSlice[i].Digit != logSlice[j].Digit {
// 			return logSlice[i].Digit < logSlice[j].Digit
// 		}
// 		if logSlice[i].Digit == 1 && logSlice[j].Digit == 1 {
// 			return false
// 		}
// 		if logSlice[i].Content != logSlice[j].Content {
// 			return logSlice[i].Content < logSlice[j].Content
// 		}
// 		return logSlice[i].Prefix < logSlice[j].Prefix
// 	})

// 	logs = []string{}
// 	for _, letter := range logSlice {
// 		logs = append(logs, fmt.Sprintf("%s %s", letter.Prefix, letter.Content))
// 	}
// 	return logs
// }

func IsDigitsOnly(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
