package engine

type Requset struct {
	Url string
	ParserFunc func([]byte) ParserResult
}
type ParserResult struct {
	Requsets []Requset
	Items []interface{}
}

func NewParser([]byte) ParserResult {
	return ParserResult{}
}