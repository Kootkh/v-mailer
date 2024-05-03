package models

import "strings"

// Attachment - Описание структуры для объекта файла вложений "Attachment" ...
type Attachment struct {
	FilePath       string `validate:"omitempty,filepath"`
	AttachmentName string `validate:"omitempty,unique=FilePath"`
	MimeType       string `validate:"omitempty"`
	EncodingType   string `validate:"omitempty"`
	Inline         bool   `validate:"omitempty, boolean"`
	Name           string `validate:"omitempty"`
}

func NewAttachment() *Attachment {
	attachment := new(Attachment)
	attachment.EncodingType = "base64"
	return attachment
}

type AttachmentsList []string

func NewAttachmentsList() *AttachmentsList {
	return &AttachmentsList{}
}

// String returns the string representation of AttachmentsList
func (al *AttachmentsList) String() string {
	return strings.Join(*al, ";")
}

func (al *AttachmentsList) Len() int {
	return len(*al)
}

func (al *AttachmentsList) Get(i int) string {
	return (*al)[i]
}

func (al *AttachmentsList) Add(value string) {
	*al = append(*al, value)
}

// Set implements the flag.Value interface for AttachmentsList
func (al *AttachmentsList) Set(value string) error {
	arr := strings.Split(value, ";")

	for _, v := range arr {
		*al = append(*al, strings.TrimSpace(v))
	}

	return nil
}
