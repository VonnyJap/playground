package main

func main() {}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	result := 0
	index := 0
	switch ruleKey {
	case "color":
		index = 1
	case "name":
		index = 2
	}

	for _, item := range items {
		if item[index] == ruleValue {
			result++
		}
	}

	return result
}

// 1. init a result=0
// 2. based on ruleKey => decide the index
// 3. in each loop check the items[i][index]==ruleValue => if so ++
