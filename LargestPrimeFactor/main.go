package main

import (
	"fmt"
	"math/big"
)

// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?
var primeNums []int64
var primeFactors []int64
var count int64

func prime(num int64) []int64 {
	i := big.NewInt(num)
	if i.ProbablyPrime(1) == true {
		primeNums = append(primeNums, num)
	}
	return primeNums
}

func areFactors(primes []int64) []int64 {
	for _, i := range primes {
		if 600851475143%i == 0 {
			primeFactors = append(primeFactors, i)
		}
	}
	return primeFactors
}
func main() {
	for count = 0; count <= 600851475143; count++ {
		prime(count)
	}

	areFactors(primeNums)
	fmt.Println(primeFactors[len(primeFactors)-1])

}
