package main

import (
	"fmt"
	"strconv"
)

//A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
//Find the largest palindrome made from the product of two 3-digit numbers.

//Function to check if product is a palindrome.
func isPalindrome(x int) bool {
	//Convert int to string
	numString := strconv.Itoa(x)
	//create a slice of runes in the converted string
	numRune := []rune(numString)
	var pdrome []int32

	for i := 1; i <= len(numRune); i++ {
		pdrome = append(pdrome, numRune[len(numRune)-i])
	}

	for i := range numRune {
		if numRune[i] != pdrome[i] {
			return false
		}
	}
	return true
}

func highestNum(pdromes []int) int {
	topNum := 0
	for i := range pdromes {
		if pdromes[i] > topNum {
			topNum = pdromes[i]
		}
		continue
	}
	return topNum
}

func main() {
	var highest []int
	for bignum := 999; bignum >= 100; bignum-- {
		for littlenum := 999; littlenum >= 100; littlenum-- {
			num := bignum * (littlenum)
			if isPalindrome(num) == true {
				highest = append(highest, num)

			}
		}
	}
	fmt.Println(highestNum(highest))
}
