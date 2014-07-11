package main

import "fmt"

func MyCount(input string, cnt int) {
	fmt.Println("Input: ",
		input, ": called(",
		cnt, ") times")
}

func PrintWithCount() func(string) {
	cnt := 0;
	return func(msg string) {
		MyCount(msg, cnt)
		cnt += 1
	}	
}

func main() {
	printWithCount :=  PrintWithCount()
	printWithCount("hello")
	printWithCount("world")
}
