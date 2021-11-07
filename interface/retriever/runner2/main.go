package main

import (
	"fmt"
	"imooc.com/weihc/learngo/interface/retriever/mock"
)

var url = "https://www.imooc.com"

type Retriever interface {
	Get(string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

type Poster interface {
	Post(string, map[string]string) string
}

func post(p Poster) string {
	return p.Post(url, map[string]string{"name": "golang"})
}

/*
接口还可以组合接口
*/

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(rp RetrieverPoster) {
	fmt.Println(rp.Post(url, map[string]string{"content": "session Post()"}))
	fmt.Println(rp.Get(url))
}

func main() {
	mr := mock.Retriever{}
	session(mr)

	mr2 := &mock.Retriever2{}
	session(mr2)
}
