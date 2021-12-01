package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {
	maxNumber := 100000000
	if len(os.Args) > 1 {
		if num, err := strconv.Atoi(os.Args[1]); err == nil {
			maxNumber = num
		}
	}

	sum1, elapsed1 := slow(maxNumber)
	log.Printf("The slow() function took %v for make a sum of all numbers from 1 to %d,"+
		" and the sum was %d", elapsed1, maxNumber, sum1)

	sum2, elapsed2 := fast(maxNumber)
	log.Printf("The fast() function took %v for make a sum of all numbers from 1 to %d,"+
		" and the sum was %d", elapsed2, maxNumber, sum2)

}

func slow(n int) (int, time.Duration) {
	start := time.Now()
	sum := 0
	for i := 0; i <= n; i++ { // Loop through all the numbers one by one
		sum += i
	}
	elapsed := time.Since(start)
	return sum, elapsed
}

func fast(n int) (int, time.Duration) {
	start := time.Now()
	sum := int(math.Pow(float64(n), 2))/2 + n/2 // Gauss summary. No loop needed.
	elapsed := time.Since(start)
	return sum, elapsed
}
