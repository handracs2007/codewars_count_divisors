package main

import "fmt"

func Divisors(n int) int {
	count := 1

	for i := 2; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(Divisors(500000))
}
