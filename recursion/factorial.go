package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 10; i++ {
		ans := factorial(i)
		fmt.Printf("\nfactorial of %v is %v", i, ans)
		fmt.Printf("\nfactorial of %v is %v (recursion)", i, recursiveFactorial(i))
	}
}

func factorial(num int) int {
	factorial := 1
	for i := 1; i <= num; i++ {
		factorial = factorial * i
	}
	return factorial
}

func recursiveFactorial(num int) int {
	if num == 1 {
		return 1
	}

	return num * recursiveFactorial(num-1)
}
