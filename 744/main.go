package main

import "fmt"

func main() {

	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a')))
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c')))
	fmt.Println(string(nextGreatestLetter([]byte{'x', 'x', 'y', 'y'}, 'z')))

}

func nextGreatestLetter(letters []byte, target byte) byte {

	for _, b := range letters {
		if b > target {
			return b
		}
	}

	return letters[0]
}
