package main

import (
	"fmt"
	"strconv"
)

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		nt := "odd"
		if number%2 == 0 {
			nt = "even"
		}
		fmt.Println(strconv.Itoa(number) + " is " + nt)
	}
}
