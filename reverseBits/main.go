package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := int64(17)
	fmt.Println(strconv.FormatInt(1, 2))

	fmt.Println(strconv.FormatInt(n, 2))
	// fmt.Println(strconv.FormatInt(1, 2))

	// fmt.Println(strconv.FormatInt(n&1, 2))
	fmt.Println(strconv.FormatInt(int64(n>>1), 2))
	fmt.Println(strconv.FormatInt(int64((n&1)>>1), 2))

	fmt.Println(strconv.FormatInt(int64(n<<1), 2))
	fmt.Println(strconv.FormatInt(int64((n&1)<<1), 2))

}

func reverseBits(num uint32) uint32 {
	ret := uint32(0)
	power := uint32(31)
	for num != 0 {
		fmt.Println("boom: ", num&1)
		ret += (num & 1) << power
		// << will add all trailing zero
		// & will check the last bit and push it to the most front
		num = num >> 1 // will remove one number at the back
		fmt.Printf("bang: %v\n", uint32(num))
		power -= 1
	}
	return ret
}
