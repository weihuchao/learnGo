package main

import (
	"fmt"
	"imooc.com/weihc/learngo/interface/commonInterface"
)

type Retriever interface {
	Get(string) string
}

type Poster interface {
	Post(string, map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func main() {
	var r RetrieverPoster
	r = &commonInterface.Retriever{Content: "commonInterface.Retriever"}
	fmt.Println(r)
	//&{commonInterface.Retriever}

	r = &commonInterface.Retriever2{Content: "commonInterface.Retriever2"}
	fmt.Println(r)
	//Retriever2 string info: commonInterface.Retriever2
}
