package main

import (
	"fmt"
	"io/ioutil"
)

const (
	filename     = "abc.txt"
	trueFilename = "basic/branch/abc.txt"
)

func readFile() {
	contents, err := ioutil.ReadFile(trueFilename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	// contents, err只在if内部生效
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

}

func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	case "%", "y":
		result = a % b
	default:
		panic("unsupported operator: " + op)
	}
	return result
}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
		fallthrough
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		g = ""
	}
	return g
}

func main() {
	readFile()
	fmt.Println(eval(3, 4, "*"))
	fmt.Println(eval(3, 4, "%"))
	fmt.Println(eval(3, 4, "y"))

	fmt.Println(0, grade(0))
	fmt.Println(59, grade(59))
	fmt.Println(61, grade(61))
	fmt.Println(82, grade(82))
	fmt.Println(92, grade(92))

}
