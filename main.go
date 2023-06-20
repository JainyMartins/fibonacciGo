package main

import (
	"fmt"

	"github.com/JainyMartins/fibonacciGo/fibonacci"
)

func main(){
	for i := 0; i < 30; i++ {
		fmt.Println("(", i, "):", fibonacci.Fibonacci(i), "\t")
	}
}