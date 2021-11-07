package main

import (
	"fmt"
	"math"
)

// 包变量
var aa = 3
var bb = 4
var cc = true

var (
	dd = 5
	ee = 6
	ff = "ff"
)

func variableZeroValue() {
	var a int
	var s string
	// https://pkg.go.dev/fmt
	fmt.Printf("%d %s\n", a, s)
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "abc"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "abc"
	b = 5
	fmt.Println(a, b, c, s)
}

func consts() {
	const filename string = "abc.txt"
	const (
		a, b = 3, 4
	)
	// 这里常量相当于是文本替换, 不用强转float64
	c := int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)

}

func enums() {
	//const (
	//	cpp = 0
	//	java = 1
	//	python = 2
	//	golang = 3
	//)
	const (
		cpp = iota
		java
		python
		golang
	)
	// 0 1 2 3
	fmt.Println(cpp, java, python, golang)

	const (
		zhang3 = iota
		li4 = "li4"
		wang5, zhao6 = iota, iota
		qian7 = iota
	)
	// 0 li4 2 2 3
	// iota是按行递增的
	fmt.Println(zhang3, li4, wang5, zhao6, qian7)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	// 1 1024 1048576 1073741824 1099511627776 1125899906842624
	fmt.Println(b, kb, mb, gb, tb, pb)

}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, cc, dd, ee, ff)

	consts()
	enums()
}
