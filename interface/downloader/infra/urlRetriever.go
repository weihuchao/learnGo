package infra

import (
	"io"
	"io/ioutil"
	"net/http"
)

type Retriever struct{}

func (r Retriever) Get(url string) string {
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
