package main

import (
	"errors"
	"fmt"
)

func main() {
	var num int = 193
	var denom int = 0

	var result, remainder, err = division(num, denom)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("The result when %v is divided by %v is %v and remainder is %v", num, denom, result, remainder)
	}

}

func division(num int, denom int) (int, int, error) {
	var err error
	if denom == 0 {
		err = errors.New("cannot divide by 0")
		return 0, 0, err
	}
	var result int = num / denom
	var remainder int = num % denom
	return result, remainder, err
}
