package main

import "fmt"

/*
函数体:
	局部变量 i
	自由变量 sum
 */

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	innerA := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0+1+...+%d=%d\n", i, innerA(i))
	}
}
