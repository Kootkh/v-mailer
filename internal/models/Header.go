package models

// Header - Описание структуры для объекта заголовка "Header" ...
type Header struct {
	name  string
	value string
}

func NewHeader() *Header {
	header := new(Header)
	return header
}

type Headers map[string]string

func NewHeaders() *Headers {
	return &Headers{}
}

type HeadersList map[string]string

func NewHeadersList() *HeadersList {
	return &HeadersList{}
}

func (hl *HeadersList) Add(name string, value string) {
	(*hl)[name] = value
}
