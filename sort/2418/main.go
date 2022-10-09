package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sortPeople([]string{"Alice", "Bob", "Bob"}, []int{155, 185, 150}))
}

type People struct {
	Name   string
	Height int
}

func sortPeople(names []string, heights []int) []string {

	var people = []People{}

	for i, name := range names {
		people = append(people, People{Name: name, Height: heights[i]})
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Height > people[j].Height
	})

	for i, p := range people {
		names[i] = p.Name
	}

	return names
}
