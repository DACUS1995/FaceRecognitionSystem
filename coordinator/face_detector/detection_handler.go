package face_detector

type Handler interface {
	Handle(*HandlerParameter)
}

type HandlerParameter struct {
	boundingBoxes []int32
	embeddings    []float32
}

type DatabaseSearcher struct {
}

type DatabaseRecord struct {
}

func NewDatatabaseSeacher() Handler {
	return &DatabaseSearcher{}
}

func (handler *DatabaseSearcher) Handle(params *HandlerParameter) {

}
