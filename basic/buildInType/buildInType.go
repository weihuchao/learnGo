package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/*
bool string
(u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
	(u)int:		位数根据操作系统来
	uintptr:	指针, 也是位数根据操作系统来
string, byte, rune
	string: 	字符串, 用""定义或者``定义, 后者可以保留换行
	byte:		8位(1字节), 用''定义
	rune:		相当于char, 32位(4字节), 字符类型, 字符串强转成它的时候会自动UTF8解码
	中文字符: 	24位(3字节)
float32, float64
complex64, complex128
*/

/*
最大整数的定义
const (
	MaxInt8   = 1<<7 - 1
	MinInt8   = -1 << 7
	MaxInt16  = 1<<15 - 1
	MinInt16  = -1 << 15
	MaxInt32  = 1<<31 - 1
	MinInt32  = -1 << 31
	MaxInt64  = 1<<63 - 1
	MinInt64  = -1 << 63
	MaxUint8  = 1<<8 - 1
	MaxUint16 = 1<<16 - 1
	MaxUint32 = 1<<32 - 1
	MaxUint64 = 1<<64 - 1
)
最大浮点数的定义
const (
	MaxFloat32             = 3.40282346638528859811704183484516925440e+38  // 2**127 * (2**24 - 1) / 2**23
	SmallestNonzeroFloat32 = 1.401298464324817070923729583289916131280e-45 // 1 / 2**(127 - 1 + 23)

	MaxFloat64             = 1.797693134862315708145274237317043567981e+308 // 2**1023 * (2**53 - 1) / 2**52
	SmallestNonzeroFloat64 = 4.940656458412465441765687928682213723651e-324 // 1 / 2**(1023 - 1 + 52)
)
 */
func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))

	// 注意1i这样的写法
	e := cmplx.Pow(math.E, 1i*math.Pi) + 1
	fmt.Println(e)

	e2 := cmplx.Exp(1i*math.Pi) + 1
	fmt.Println(e2)

	fmt.Printf("%.3f\n", e)
	fmt.Printf("%.3f\n", e2)

}

/*
类型转换是强制的
*/

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)

	// 精度丢失问题
	// https://juejin.cn/post/7010416601711427615
	var d float64 = 3.999999
	fmt.Println(int(d))
}

func main() {
	euler()
	triangle()
}
