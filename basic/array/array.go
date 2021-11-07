package main

import "fmt"

func basic() {
	// 注意方括号需要有内容, 没有内容的就变成了切片
	var arr1 [5]int
	arr2 := [3]int{1, 3, 4}
	arr3 := [...]int{2, 4, 6, 8, 10}
	// 从前往后看: 4个后面的东西, 后面的这个东西是5个int
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for i := 0; i < len(arr2); i++ {
		fmt.Print(arr2[i])
	}
	fmt.Println()

	// range可以返回一个值, 一个值的时候返回的是索引
	// 返回两个值的时候, 第一个是索引, 第二个是值
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	fmt.Println()
}

/*
数组是值类型
*/
func printArr(a [5]int) {
	a[0] = 100
	for i, v := range a {
		fmt.Println(i, v)
	}
}

// 采用指针的方式修改
func printArrPoint(a *[5]int) {
	// Go的指针非常好用, 可以直接用不需要(*a)[0] = 100
	a[0] = 100
	for i, v := range a {
		fmt.Println(i, v)
	}
}

// 传入slice进行修改
func printArrSlice(a []int) {
	a[0] = 200
	for i, v := range a {
		fmt.Println(i, v)
	}
}

func main() {
	basic()

	arr := [...]int{2, 4, 6, 8, 10}
	printArr(arr)
	fmt.Println(arr)

	// [3]int和[5]int不是同一个类型, 这样写会报错
	// cannot use arr2 (type [3]int) as type [5]int in argument to printArr
	//arr2 := [3]int{1, 3, 4}
	//printArr(arr2)

	printArrPoint(&arr)
	fmt.Println(arr)

	printArrSlice(arr[:])
	fmt.Println(arr)
}
