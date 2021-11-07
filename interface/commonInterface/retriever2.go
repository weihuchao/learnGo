package commonInterface

type Retriever2 struct {
	Content string
}

func (r *Retriever2) Get(s string) string {
	return s + " -> " + r.Content
}

func (r *Retriever2) Post(url string, form map[string]string) string {
	r.Content = form["content"]
	return "ok"
}

func (r *Retriever2) String() string {
	return "Retriever2 string info: " + r.Content
}
