package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(subdomainVisits([]string{"9001 discuss.leetcode.com"}))
	fmt.Println(subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}

func subdomainVisits(cpdomains []string) []string {

	dict := map[string]int{}

	for _, cpdomain := range cpdomains {
		total := cpdomainsUtil(cpdomain)
		for _, item := range total {
			breaking := strings.Split(item, " ")
			count, _ := strconv.Atoi(breaking[0])
			if _, ok := dict[breaking[1]]; ok {
				dict[breaking[1]] += count
				continue
			}
			dict[breaking[1]] = count
		}
	}

	result := []string{}

	for i, d := range dict {
		result = append(result, fmt.Sprintf("%d %s", d, i))
	}

	return result
}

func cpdomainsUtil(cpdomain string) []string {

	visits := strings.Split(cpdomain, " ")
	index := strings.Index(visits[1], ".")
	if index == -1 {
		return []string{cpdomain}
	}

	return append([]string{cpdomain}, cpdomainsUtil(fmt.Sprintf("%s %s", visits[0], visits[1][index+1:]))...)
}
