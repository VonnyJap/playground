package main

import "fmt"

func main() {
	regions := [][]string{
		{"Earth", "North America", "South America"},
		{"North America", "United States", "Canada"},
		{"United States", "New York", "Boston"},
		{"Canada", "Ontario", "Quebec"},
		{"South America", "Brazil"},
	}
	fmt.Println(findSmallestRegion(regions, "Quebec", "New York"))
	fmt.Println(findSmallestRegion(regions, "Canada", "South America"))
	fmt.Println(findSmallestRegion(regions, "Canada", "Quebec"))

}

func findSmallestRegion(regions [][]string, region1 string, region2 string) string {

	mapping := map[string]string{}

	for _, region := range regions {
		parent := region[0]
		for _, r := range region[1:] {
			mapping[r] = parent
		}
	}

	// fmt.Println(mapping)

	parent1 := findParents(mapping, region1)
	parent2 := findParents(mapping, region2)
	fmt.Println(parent1)
	fmt.Println(parent2)

	i := 0
	var intersection string
	for {
		if len(parent1) <= i || len(parent2) <= i {
			break
		}
		if parent1[i] != parent2[i] {
			break
		}
		intersection = parent1[i]
		i++
	}

	return intersection

}

func findParents(mapping map[string]string, region string) []string {

	if _, ok := mapping[region]; !ok {
		return []string{region}
	}

	next := mapping[region]

	return append(findParents(mapping, next), region)
}
