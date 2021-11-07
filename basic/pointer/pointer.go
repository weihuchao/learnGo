package main

import "fmt"

func basic() {
	var a int = 2
	var p *int = &a
	*p = 3
	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(*p)
}

/*
值传递, 引用传递
Go语言只有值传递一种方式, 值传递的时候会把值拷贝一份给新的变量
*/
func swap(a, b *int) {
	*a, *b = *b, *a
}

func swapRet(a, b int) (int, int) {
	return b, a
}

func main() {
	basic()

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)

	a, b = swapRet(a, b)
	fmt.Println(a, b)
}
