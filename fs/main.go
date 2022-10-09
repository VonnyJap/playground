package main

import (
	"fmt"
	"strings"
)

func main() {
	fs := Constructor()
	fsPtr := &fs
	fmt.Println(fsPtr.CreatePath("/leet", 1))
	fmt.Println(fsPtr.CreatePath("/leet/code", 2))
	fmt.Println(fsPtr.Get("/leet"))
	fmt.Println(fsPtr.Get("/leet/code"))
	fmt.Println(fsPtr.CreatePath("/c/d", 1))
	fmt.Println(fsPtr.Get("/c"))
}

// type Node struct {
// 	Name  string
// 	Value int
// 	Next  *Node
// }

// we want to make it like a graph
type FileSystem struct {
	paths map[string]int
}

func Constructor() FileSystem {
	return FileSystem{map[string]int{}}
}

func (this *FileSystem) CreatePath(path string, value int) bool {
	if len(path) < 2 {
		return false
	}
	if string(path[0]) != "/" {
		return false
	}
	var paths []string
	for _, p := range strings.Split(path, "/") {
		if p != "" {
			paths = append(paths, p)
		}
	}
	if len(paths) == 1 {
		if _, ok := this.paths[paths[0]]; ok {
			return false
		}
		this.paths[paths[0]] = value
		return true
	}

	for _, p := range paths[0 : len(paths)-1] {
		if _, ok := this.paths[p]; !ok {
			return false
		}
	}
	if _, ok := this.paths[paths[len(paths)-1]]; ok {
		return false
	}
	this.paths[paths[len(paths)-1]] = value
	return true
}

func (this *FileSystem) Get(path string) int {
	var paths []string
	for _, p := range strings.Split(path, "/") {
		if p != "" {
			paths = append(paths, p)
		}
	}
	for _, p := range paths {
		if _, ok := this.paths[p]; !ok {
			return -1
		}
	}
	return this.paths[paths[len(paths)-1]]
}

/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.CreatePath(path,value);
 * param_2 := obj.Get(path);
 */
