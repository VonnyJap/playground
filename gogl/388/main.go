package main

import (
	"fmt"
	"strings"
)

func main() {

	// dir := "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"
	dir := "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"

	fmt.Println(lengthLongestPath(dir))
}

// first define the tree thing
func lengthLongestPath(input string) int {

	var max = 0
	var splits = strings.Split(input, "\n")
	fmt.Printf("splits: %#v\n", splits)

	var stack = make([]int, len(splits)+1)

	for _, name := range splits {
		var level = strings.LastIndex(name, "\t") + 1
		var curLen = len(name) - level
		var preLen = 0
		if level > 0 {
			preLen = stack[level-1]
		}
		if strings.Contains(name, ".") {
			if preLen+curLen > max {
				max = preLen + curLen
			}
			continue
		}
		stack[level] = preLen + curLen + 1 // the "/" - hence +1
	}

	return max
}
