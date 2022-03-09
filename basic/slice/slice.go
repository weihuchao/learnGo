package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
	fmt.Println("update slice: ", s)
}

// slice是对数组的一个视图, 是引用类型
func basic() {
	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr: ", arr)

	fmt.Println("arr[2:6]: ", arr[2:6])
	fmt.Println("arr[2:]: ", arr[2:])
	fmt.Println("arr[:6]: ", arr[:6])
	fmt.Println("arr[:]: ", arr[:])

	s1 := arr[2:6]
	updateSlice(s1)
	fmt.Println("arr: ", arr)
}

/*
可以对slice再次slice, 只要在满足cap的情况下可以往后扩展
但是不能:
	向前扩展
	也不能越界取值
 */
func reSlice() {
	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr: ", arr)

	s := arr[:5]
	fmt.Println("arr[:5]", s)

	s = s[2:]
	fmt.Println("s[2:]", s)
	fmt.Println("arr[:5][2:]", arr[:5][2:])
}

/*
slice的实现, 组成部分:
	ptr
	len
	cap
*/
func extendSlice() {
	var arr = [...]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println("arr: ", arr)

	fmt.Println("重复切片:")
	s1 := arr[2:6]
	fmt.Println("arr[2:6]", s1)

	s2 := s1[3:5]
	fmt.Println("arr[2:6][3:5]", s2)

	// panic: runtime error: slice bounds out of range [:6] with capacity 5
	//fmt.Println("arr[2:6][3:6]", arr[2:6][3:6])

	fmt.Println("切片可扩展范围:")
	fmt.Printf("arr[3:4]: %v, len: %d, cap: %d \n", arr[3:4], len(arr[3:4]), cap(arr[3:4]))
	fmt.Printf("arr[2:6]: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("arr[2:6][3:5]: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))
}

/*
append函数的功效
	修改数组对应索引位置的值;
	传入的slice不会变化;
	返回新的slice, 这个新的slice实际上就是对array的向后view;
	之所以有函数返回值, 是因为slice经过append之后里面的ptr/len/cap都可能发生变化
	如果数组长度不足, 会拷贝出相同内容的数组, 但是cap更大, 再更改新数组的具体位置的值
		后续仍然是不改变源slice
		新返回一个slice, 这个slice指向新的array
 */
func appendSlice() {
	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr: ", arr)

	s1 := arr[2:6]
	s2 := s1[3:5]
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)

	fmt.Println("s1=arr[2:6] ", s1)
	fmt.Println("s2=arr[2:6][3:5] ", s2)
	fmt.Println("s3:=append(s2, 10) ", s3)
	fmt.Println("s4:=append(s3, 11) ", s4)
	fmt.Println("s5:=append(s4, 12) ", s5)
	fmt.Println("arr: ", arr)

	// s4和s5已经超出了源arr的范围了, 因此实际上s4和s5已经指向了和arr相同值但是不同的array上了
	s3[2] = 100
	fmt.Println("s3[2] = 100")
	fmt.Println("arr, s3, s4, s5: ", arr, s3, s4, s5)

	s4[3] = 500
	fmt.Println("s4[3] = 500")
	fmt.Println("arr, s3, s4, s5: ", arr, s3, s4, s5)

}

func showSliceInfo(s []int, showValue bool) {
	if showValue {
		fmt.Printf("%v, len: %d, cap: %d \n", s, len(s), cap(s))
	} else {
		fmt.Printf("len: %d, cap: %d \n", len(s), cap(s))
	}

}

/*
slice的创建:
	1. 不赋初值的时候为nil, 可以操作
	2. 赋初值
	3. make, make可以申明len和cap

slice的cap是按照2倍的方式增加的 1,2,4,8,16,32,64,128...

可以用append的方式来移除slice的某个值, 但是append会修改底层数组
*/
func useSlice() {
	// 空slice和cap的递增
	var s1 []int
	for i := 0; i < 100; i++ {
		showSliceInfo(s1, false)
		s1 = append(s1, i)
	}

	// 在没有append的时候, len=cap
	s2 := []int{2, 4, 6, 8}
	showSliceInfo(s2, true)
	// append两个元素之后, len=len+2, cap=cap*2
	s2 = append(s2, 10, 11)
	showSliceInfo(s2, true)

	s3 := make([]int, 16)
	s4 := make([]int, 10, 32)
	showSliceInfo(s3, true)
	showSliceInfo(s4, true)

	// copy是把s2对应的每个位子拷贝到s3上, 对s2的修改会体现到s3上
	// copy是浅拷贝
	fmt.Println("copy(s3, s2) >:")
	copy(s3, s2)
	showSliceInfo(s3, true)

	fmt.Println("删除s3[3] >:")
	// 删除s3[3]
	// 使用...来展开序列
	s3 = append(s3[:3], s3[4:]...)
	showSliceInfo(s3, true)

	fmt.Println("考察值传递问题 >:")
	// 考察值传递问题, append会修改底层的数组
	s6 := s3[:6]
	showSliceInfo(s6, true)
	// 继续删除s3[3]
	s3 = append(s3[:3], s3[4:]...)
	showSliceInfo(s3, true)
	showSliceInfo(s6, true)

}

func main() {
	//basic()
	//reSlice()
	//extendSlice()
	//appendSlice()
	useSlice()
}

