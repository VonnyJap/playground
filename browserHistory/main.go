package main

import "fmt"

func main() {
	history := Constructor("leetcode.com")
	(&history).Visit("google.com")
	(&history).Visit("facebook.com")
	(&history).Visit("youtube.com")
	fmt.Println((&history).Back(1))
	fmt.Println((&history).Back(1))
	fmt.Println((&history).Forward(1))
	(&history).Visit("linkedin.com")
	fmt.Println((&history).Forward(2))
	fmt.Println((&history).Back(2))
	fmt.Println((&history).Back(7))

}

type BrowserHistory struct {
	sites    []string
	position int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		sites:    []string{homepage},
		position: 0,
	}

}

func (this *BrowserHistory) Visit(url string) {
	this.sites = append(this.sites[:this.position+1], url)
	this.position = len(this.sites) - 1
}

func (this *BrowserHistory) Back(steps int) string {
	x := this.position - steps
	if x >= 0 {
		this.position = x
	} else {
		this.position = 0
	}
	return this.sites[this.position]
}

func (this *BrowserHistory) Forward(steps int) string {
	x := this.position + steps
	if x < len(this.sites) {
		this.position = x
	} else {
		this.position = len(this.sites) - 1
	}
	return this.sites[this.position]
}
