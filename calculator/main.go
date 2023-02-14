package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println(calculate("3+21*2/2"))
	// fmt.Println(calculate("3+2 * 2 "))
	// fmt.Println(calculate(" 3+5 / 2 "))
	// fmt.Println(calculate("0*0"))
	fmt.Println(calculate("1-1+1"))
}

func calculate(s string) int {

	if len(s) == 0 {
		return 0
	}
	s = strings.ReplaceAll(s, " ", "")
	fmt.Println(s)
	expList := []string{}

	for len(s) > 0 {
		arr := []int{}
		plus := strings.Index(s, "+")
		if plus >= 0 {
			arr = append(arr, plus)
		}
		minus := strings.Index(s, "-")
		if minus >= 0 {
			arr = append(arr, minus)
		}
		multiply := strings.Index(s, "*")
		if multiply >= 0 {
			arr = append(arr, multiply)
		}
		divide := strings.Index(s, "/")
		if divide >= 0 {
			arr = append(arr, divide)
		}
		if len(arr) == 0 {
			expList = append(expList, s)
			break
		}
		index := min(arr)
		expList = append(expList, s[:index])
		expList = append(expList, s[index:index+1])
		s = s[index+1:]
	}

	return calculateUtil(expList)
}

func calculateUtil(expList []string) int {

	fmt.Println("expList", expList)
	if len(expList) == 0 {
		return 0
	}
	first, _ := strconv.Atoi(expList[0])
	if len(expList) == 1 {
		return first
	}
	if expList[1] == "+" {
		return first + calculateUtil(expList[2:])
	}
	if expList[1] == "-" {
		return first - calculateUtil(expList[2:])
	}
	second, _ := strconv.Atoi(expList[2])
	result := first * second
	if expList[1] == "/" {
		result = first / second
	}
	return calculateUtil(append([]string{fmt.Sprintf("%d", result)}, expList[3:]...))
}

func min(arr []int) int {

	min := arr[0]

	for _, num := range arr[1:] {
		if min > num {
			min = num
		}
	}

	return min
}
