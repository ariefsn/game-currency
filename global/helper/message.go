package helper

type Message struct {
	Duplicate            string
	DuplicateId          string
	MethodNotAllowed     string
	RouteNotFound        string
	DataNotFound         string
	EmptyResponse        string
	FromCurrencyNotFound string
	ToCurrencyNotFound   string
}

func NewMessage() *Message {
	m := new(Message)

	m.Duplicate = "duplicate data"
	m.DuplicateId = "duplicate id"
	m.MethodNotAllowed = "method not allowed"
	m.RouteNotFound = "route not found"
	m.DataNotFound = "data not found"
	m.EmptyResponse = "service response is nil"
	m.FromCurrencyNotFound = "from currency not found"
	m.ToCurrencyNotFound = "to currency not found"

	return m
}
