package models

// Body - Описание структуры для объекта тела сообщения "Body" ...
type Body struct {
	File         string
	Message      string
	Content      string
	MimeType     string
	EncodingType string
	Disposition  string
	CharacterSet string
}

func NewBody() *Body {
	body := new(Body)
	return body
}
