package testor

type Retriever struct{}

func (r Retriever) Get(url string) string {
	return "This is testor.Retriever.Get()`s content."
}
