package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func eval(a, b int, op string) (int, error) {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		q, _ := div(a, b)
		result = q
	case "%", "y":
		result = a % b
	default:
		return 0, fmt.Errorf("unsupported operator: %s", op)
	}
	return result, nil
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("apply function %s with args: (%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for _, num := range numbers {
		s += num
	}
	return s
}

func main() {
	q, r := div(10, 3)
	fmt.Println(q, r)

	if ret, err := eval(10, 3, "x"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ret)
	}

	fmt.Println(apply(pow, 3, 4))
	fmt.Println(apply(func(i int, i2 int) int {
		return i + i2
	}, 3, 4))

	fmt.Println(sum(4, 5, 6, 7, 8))
}
