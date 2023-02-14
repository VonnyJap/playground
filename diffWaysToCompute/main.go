package main

import "fmt"

func main() {
	fmt.Println(diffWaysToCompute("2-1-1"))
	fmt.Println(diffWaysToCompute("2*3-4*5"))

}

// 2*3-4*5
func diffWaysToCompute(expression string) []string {

	if len(expression) == 0 {
		return []string{}
	}
	if len(expression) == 1 {
		return []string{expression}
	}
	result := []string{}
	ctr := 2
	for ctr < len(expression) {
		prefix := expression[:ctr]
		suffixes := diffWaysToCompute(expression[ctr:])
		for _, s := range suffixes {
			result = append(result, fmt.Sprintf("(%s)%s(%s)", prefix[0:len(prefix)-1], string(prefix[len(prefix)-1]), s))
		}
		ctr += 2
	}

	return result
}
