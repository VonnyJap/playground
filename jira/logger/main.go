package main

func main() {}

type Logger struct {
	Collection map[string]int
}

func Constructor() Logger {
	return Logger{map[string]int{}}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	next, ok := this.Collection[message]
	if !ok {
		this.Collection[message] = timestamp + 10
		return true
	}
	if timestamp >= next {
		this.Collection[message] = timestamp + 10
		return true
	}
	return false
}
