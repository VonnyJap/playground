package main

import "fmt"

func main() {
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 4}))
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 5, 6, 7}))
	fmt.Println(numberOfArithmeticSlices([]int{1}))
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 4, 5}))
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 4, 7, 10}))

}

// [1,2,3,4,7,10]
func numberOfArithmeticSlices(nums []int) int {

	total := 0
	for size := 3; size <= len(nums); size++ {

		temp := []int{}
		i := 0
		diff := 0
		for i < len(nums) {
			// pop one out
			if len(temp) == size {
				temp = temp[1:size]
			}
			temp = append(temp, nums[i])
			if len(temp) == 1 {
				i++
				continue
			}
			if len(temp) == 2 {
				diff = temp[1] - temp[0]
			}
			// when we reach here - compare between the last val vs the val before
			// if not differ by 1 then make this thing leaving only 1 value - last value
			if temp[len(temp)-1]-temp[len(temp)-2] != diff {
				temp = temp[len(temp)-2:]
				diff = temp[1] - temp[0]
				i++
				continue
			}
			if len(temp) == size {
				total++
			}
			i++
		}
	}

	return total
}

// make window size as the outer loop - from 3 to whatever len of nums - name it size
// maintain an arr each time the size of arr equal to 3 then +1
