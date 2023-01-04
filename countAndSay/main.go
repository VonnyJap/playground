package main

import "fmt"

func main() {
	fmt.Println(countAndSay(2))
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay(5))

}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	prev := countAndSay(n - 1)

	var result string
	i := 1
	local := 1

	for i < len(prev) {
		if prev[i] != prev[i-1] {
			result = fmt.Sprintf("%s%d%s", result, local, string(prev[i-1]))
			local = 1
			i++
			continue
		}
		local++
		i++
	}

	result = fmt.Sprintf("%s%d%s", result, local, string(prev[i-1]))
	return result
}

// base case n == 1
// for n > 1
// 1. get result from n-1
// 2. based on the result from 1 -> create map
// 3. create empty string
// 4. loop again and find the corresponding count for each char
// 5. return result
