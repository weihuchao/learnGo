package mock

type Retriever struct {
	Content string
}

func (r Retriever) Get(s string) string {
	return s + " -> " + r.Content
}

func (r Retriever) Post(url string, form map[string]string) string {
	r.Content = form["content"]
	return "ok"
}
