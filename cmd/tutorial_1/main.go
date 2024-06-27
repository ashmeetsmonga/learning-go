package main

import "fmt"

func main() {
	var num int = 193
	var denom int = 4

	var result, remainder int = division(num, denom)
	fmt.Printf("The result when %v is divided by %v is %v and remainder is %v", num, denom, result, remainder)
}

func division(num int, denom int) (int, int) {
	var result int = num / denom
	var remainder int = num % denom
	return result, remainder
}
