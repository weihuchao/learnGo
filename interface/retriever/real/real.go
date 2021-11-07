package real

import (
	"io"
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

//func (r Retriever) Get(url string) string {
func (r *Retriever) Get(url string) string {
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

	bytes, err := httputil.DumpResponse(res, true)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
