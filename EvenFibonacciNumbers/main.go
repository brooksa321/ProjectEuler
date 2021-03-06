package main

import (
	"fmt"
)

//Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
//1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
//By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.

func main() {
	previous := 1
	current := 2
	evenSum := 2
	var fibNum int
	for {

		if fibNum > 4000000 {
			fmt.Println(evenSum)
			break
		}

		fibNum = previous + current

		if fibNum%2 == 0 {
			evenSum += fibNum
		}
		previous = current
		current = fibNum
	}
}
