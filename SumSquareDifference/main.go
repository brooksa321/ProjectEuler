package main

import (
	"fmt"
)

func sumOfsquares(n int) int {
	var sumTotal int
	for x := 0; x <= n; x++ {
		sumTotal += x * x
	}
	return sumTotal
}

func squareOfsums(n int) int {
	var squareTotal int
	for x := 0; x <= n; x++ {
		squareTotal += x
	}
	squareTotal = squareTotal * squareTotal
	return squareTotal
}

func main() {
	fmt.Println(squareOfsums(100) - sumOfsquares(100))
}
