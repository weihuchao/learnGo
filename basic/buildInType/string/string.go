package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/*
字符串的长度是字节数, 等同于[]byte(s)的长度
中文字符是3个字节
英文字符是1个字节

[]rune(s)是字符数, 经过解码之后的, 还可以使用utf8.RuneCountInString(s)获取
 */
func printChinese() {
	s := "Just学习!"
	// 获得的是字节数, 中文字符是3个字节, 所以是11
	fmt.Println(len(s))
	fmt.Println(len([]byte(s)))

	for _, b := range []byte(s) {
		// %X: Integer base 16, with upper-case letters for A-F
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	// 4A 75 73 74 E5 AD A6 E4 B9 A0 21

	for i, ch := range s {
		// %d: Integer base 10
		fmt.Printf("(%d-%X) ", i, ch)
	}
	fmt.Println()
	// (0-4A) (1-75) (2-73) (3-74) (4-5B66) (7-4E60) (10-21)

	// 获得的是字符(rune)数量, 这才是预期的7
	fmt.Println(utf8.RuneCountInString(s))
	fmt.Println(len([]rune(s)))

	bytes := []byte(s)
	for len(bytes) > 0 {
		// utf8.DecodeRune()会解码出第一个UTF8的字符返回并返回它的长度
		decodeRune, size := utf8.DecodeRune(bytes)
		// %c: Integer the character represented by the corresponding Unicode code point
		fmt.Printf("%c ", decodeRune)
		bytes = bytes[size:]
	}
	fmt.Println()

	// 直接解码成rune
	for i, ch := range []rune(s) {
		fmt.Printf("(%d-%c) ", i, ch)
	}
	fmt.Println()

}

func lengthOfMaxSubString(s string) int {
	lastOccurred := make(map[rune]int)
	start, maxLen := 0, 0

	//for i, ch := range []byte(s) {
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLen
}

func stringOperations() {
	// https://pkg.go.dev/strings

	fmt.Println(strings.Contains("abcb", "bc"))
	fmt.Println(strings.Count("abcbc", "bc"))
	// 根据空格分割字符串
	fmt.Println(strings.Fields("张三疯"))
	for i, s := range strings.Fields("Just 学习!") {
		fmt.Println(i, s)
	}
	// 转成大写
	fmt.Println(strings.ToUpper("abcbc"))
	// 移除两边的指定字符
	fmt.Println(strings.Trim("!!Hello world!!", "!"))
	fmt.Println(strings.TrimLeft("!!Hello world!!", "!"))
	fmt.Println(strings.TrimRight("!!Hello world!!", "!"))
}

func main() {
	printChinese()

	fmt.Println("abcba", lengthOfMaxSubString("abcba"))
	fmt.Println("aaaaa", lengthOfMaxSubString("aaaaa"))
	fmt.Println("一二三二一", lengthOfMaxSubString("一二三二一"))
	fmt.Println("", lengthOfMaxSubString(""))

	stringOperations()
}
