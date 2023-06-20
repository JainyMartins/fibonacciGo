package tests_test

import (
	"testing"

	"github.com/JainyMartins/fibonacciGo/fibonacci"
	"github.com/stretchr/testify/assert"
)

func TestFibonacciWith0(t *testing.T){
	result := fibonacci.Fibonacci(0)
	assert.Equal(t, 0, result)
}

func TestFibonacciWith1(t *testing.T){
	result := fibonacci.Fibonacci(1)
	assert.Equal(t, 1, result)
}

func TestFibonacciWithMoreThen1(t *testing.T){
	result := fibonacci.Fibonacci(10)
	assert.Equal(t, 55, result)
}