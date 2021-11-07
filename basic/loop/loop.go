package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func basicLoop() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func convertToBin(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	// introduce variable: win + shift + l
	open, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(open)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func foreverLoop() {
	for {
		fmt.Println("ok")
	}
}

func main() {
	basicLoop()

	fmt.Println(convertToBin(5))
	fmt.Println(strconv.FormatInt(5, 2))

	fmt.Println(convertToBin(19))
	fmt.Println(strconv.FormatInt(19, 2))

	fmt.Println(convertToBin(0))
	fmt.Println(strconv.FormatInt(0, 2))

	printFile("basic/branch/abc.txt")

	foreverLoop()
}
