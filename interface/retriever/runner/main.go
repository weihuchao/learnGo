package main

import (
	"fmt"
	"imooc.com/weihc/learngo/interface/retriever/mock"
	"imooc.com/weihc/learngo/interface/retriever/real"
	"time"
)

type Retriever interface {
	Get(string) string
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

func inspect(r Retriever) {
	fmt.Printf("%T, %v", r, r)
	fmt.Println()
}

func main() {
	var r Retriever
	r = mock.Retriever{Content: "mock content"}
	inspect(r)
	fmt.Println("Download---->", download(r))

	//rr := real.Retriever{}
	//fmt.Println(download(rr)[:100])

	// 可以定义一个指针的方法, 这样就需要取地址之后才能给Retriever接口
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	fmt.Println("Download---->", download(r)[:100])

	// 此时不能r.TimeOut, 需要转化类型, 方法为 varName.(ToType)
	if realRetriever, ok := r.(*real.Retriever); ok {
		// 这样才能.TimeOut
		fmt.Println(realRetriever.TimeOut)
	}
}
