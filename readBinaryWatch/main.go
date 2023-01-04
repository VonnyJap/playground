package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println(readBinaryWatch(1))
	// fmt.Println(readBinaryWatch(2))
	// fmt.Println(readBinaryWatch(9))

	fmt.Println(1 ^ 3)
}

func readBinaryWatch(turnedOn int) []string {

	if turnedOn == 0 {
		return []string{}
	}

	watches := []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"}
	result := []string{}

	for _, w := range watches {
		addList := readBinaryWatch(turnedOn - 1)
		if len(addList) == 0 {
			result = append(result, w)
		}
		items := strings.Split(w, ":")
		hourInt, _ := strconv.Atoi(items[0])
		minInt, _ := strconv.Atoi(items[1])
		for _, add := range addList {
			if add == w {
				continue
			}
			arr := strings.Split(add, ":")
			hourAddInt, _ := strconv.Atoi(arr[0])
			minAddInt, _ := strconv.Atoi(arr[1])
			carry := 0
			totalMin := minAddInt + minInt
			if totalMin >= 60 {
				carry = 1
				totalMin = totalMin - 60
			}
			totalHour := hourAddInt + hourInt + carry
			if totalHour > 12 {
				continue
			}
			if totalMin >= 10 {
				result = append(result, fmt.Sprintf("%d:%d", totalHour, totalMin))
				continue
			}
			result = append(result, fmt.Sprintf("%d:0%d", totalHour, totalMin))
		}

	}

	return result
}

// type Watch struct {
// 	Num  int
// 	Hour bool
// }
