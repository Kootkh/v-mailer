package types

import (
	"strings"
)

type Attachments struct {
	Attachment []Attachment `validate:"omitempty,dive,unique"`
}

func NewAttachments() *Attachments {
	return &Attachments{
		Attachment: []Attachment{},
	}
}

func (a *Attachments) Set(at *Attachment) {
	a.Attachment = append(a.Attachment, *at)
}

func (a *Attachments) Add(attachment Attachment) {
	a.Attachment = append(a.Attachment, attachment)
}

func (a *Attachments) Get(i int) Attachment {
	return a.Attachment[i]
}

func (a *Attachments) Len() int {
	return len(a.Attachment)
}

type Attachment struct {
	FilePath       string `validate:"omitempty,filepath"`
	AttachmentName string `validate:"omitempty,unique=FilePath"`
	MimeType       string `validate:"omitempty"`
	EncodingType   string `validate:"omitempty"`
	Name           string `validate:"omitempty"`
}

func NewAttachment() *Attachment {
	return &Attachment{}
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
