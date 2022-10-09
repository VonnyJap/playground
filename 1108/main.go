package main

import "strings"

func main() {}

func defangIPaddr(address string) string {

	arr := strings.Split(address, ".")

	return strings.Join(arr, "[.]")
}
