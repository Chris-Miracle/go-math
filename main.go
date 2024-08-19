package main

import (
	calc "Dev/Learnings/Go/go-book/math"
	"fmt"
)

func main() {
	xs := []float64{10, 4, 0.5, 1, 2, 23, 3, 4}
	fmt.Println(calc.Average(xs))
	fmt.Println(calc.Min(xs))
	fmt.Println(calc.Max(xs))
	fmt.Println(calc.Fib(35))
}
