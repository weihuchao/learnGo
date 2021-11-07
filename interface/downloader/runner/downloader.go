package main

import (
	"fmt"
	"imooc.com/weihc/learngo/interface/downloader/infra"
	"imooc.com/weihc/learngo/interface/downloader/testor"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
)

func retrieve(url string) string {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	bytes, _ := ioutil.ReadAll(res.Body)
	return string(bytes)
}

func getRetriever() retriever {
	if rand.Int() == 1 {
		return infra.Retriever{}
	} else {
		return testor.Retriever{}
	}
}

/*
引入原因:
	解耦, 本来我们只需要一个东西, 它能够Get返回一个东西即可;
	之所以不关心这个东西是什么的原因是, 有可能这是实际开发的infra.Retriever{}, 也有可能是测试用的testor.Retriever{}.

接口的含义:
	定义一个类型, 这个类型满足它体内要求的方法就可以;
	它相当于是一类事物的统称, 有点像是基类, 其余实现是继承了它的子类;

*/
type retriever interface {
	Get(string) string
}

func main() {
	//fmt.Println(retrieve("https://www.imooc.com"))

	//fmt.Println(infra.Retriever{}.Get("https://www.imooc.com"))

	//var retriever infra.Retriever = getRetriever()
	//fmt.Println(retriever.Get("https://www.imooc.com"))

	var r retriever = getRetriever()
	fmt.Println(r.Get("https://www.imooc.com"))
}
