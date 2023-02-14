package main

import "strings"

func main() {

}

/**
Given a csv file, we want to read the values one-by-one and print it to the output. For example:

Input:

hello.csv:
Hello,world,good,morning

Output:

Hello
world
good
morning

However, the file has been encrypted! So you will need to read the content of the file via the provided class below.

type FileReader struct {
    …
}

func (f* FileReader) openFile(filename string) {
}

func (f* FileReader) closeFIle() {
}

func (f *FileReader) read(charCount int) string {
    // secret implementation that decrypts the file content
   // and return a string with length charCount or less in case you read end of file
   // it will return empty string
}

FileReader f;
f.openFile(“hello.csv”);
f.read(5) // returns “Hello”
f.closeFile();

// please implement this
// returns the next value if available, returns empty string otherwise
Func (f* FileReader) nextValue() string {
}
*/

type FileReader struct {
	Remaining string
}

func (f *FileReader) openFile(filename string) {
}

func (f *FileReader) closeFIle() {
}

func (f *FileReader) read(charCount int) string {
	// secret implementation that decrypts the file content
	// and return a string with length charCount or less in case you read end of file
	// it will return empty string
}

func (f *FileReader) nextValue() string {
	// Hello,world,good,morning
	// f.read() => what will be the param to give
	// an input count => EOF
	// steps:
	// 1. read 1 by 1 in a loop
	// 2. when comma is found then break, or when the read actually return empty
	// 3. then return that string when bumping into comma

	// var word string
	// for {

	// }

	// strings.Index("chicken", "ken")

	// check remaining if comma can be found
	index := strings.Index(f.Remaining, ",")
	if index > 0 {
		word := f.Remaining[0:index]
		f.Remaining = f.Remaining[index+1:]
		return word
	}
	// "home,wel"
	// ",,,"

	chunk := 1000
	for {
		response := f.read(chunk) // =>  ",,,"
		f.Remaining += response
		if strings.Index(response, ",") >= 0 || response == "" {
			break
		}
	}

	// Hello,world,good,morning

	index = strings.Index(f.Remaining, ",") // =>  ",,," => 0
	word := f.Remaining[0:index]
	f.Remaining = f.Remaining[index+1:] // => ",,"

	return word
}

// f.openFile(“hello.csv”);
// f.read(5)
// f.closeFile();
