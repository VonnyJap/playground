package main

import "fmt"

func main() {
	fmt.Println(destCity([][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}))
	fmt.Println(destCity([][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}))
}

func destCity(paths [][]string) string {
	var dict = map[string][]string{}

	for _, path := range paths {
		if _, ok := dict[path[0]]; !ok {
			dict[path[0]] = []string{path[1]}
		} else {
			dict[path[0]] = append(dict[path[0]], path[1])
		}
		if _, ok := dict[path[1]]; !ok {
			dict[path[1]] = []string{}
		}
	}

	var last string
	for city, targets := range dict {
		if len(targets) == 0 {
			last = city
			break
		}
	}
	return last
}
