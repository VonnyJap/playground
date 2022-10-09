package main

import "fmt"

func main() {}

var solution = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	return func(buf []byte, n int) int {

		var total int = 0
		buf = make([]byte, n)

		for n > 0 {
			buf4 := make([]byte, 4)
			num := read4(buf4)
			for i := 0; i < num; i++ {
				buf[total+i] = buf4[i]
			}
			total = total + num
			if num < 4 {
				break
			}
			n = n - total
		}
		fmt.Println(string(buf))
		fmt.Println(total)

		return total
	}
}

/**
 * The read4 API is already defined for you.
 *
 *     read4 := func(buf4 []byte) int
 *
 * // Below is an example of how the read4 API can be called.
 * file := File("abcdefghijk") // File is "abcdefghijk", initially file pointer (fp) points to 'a'
 * buf4 := make([]byte, 4) // Create buffer with enough space to store characters
 * read4(buf4) // read4 returns 4. Now buf = ['a','b','c','d'], fp points to 'e'
 * read4(buf4) // read4 returns 4. Now buf = ['e','f','g','h'], fp points to 'i'
 * read4(buf4) // read4 returns 3. Now buf = ['i','j','k',...], fp points to end of file
 */
