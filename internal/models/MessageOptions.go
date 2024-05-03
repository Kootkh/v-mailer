package models

type MessageOptions struct {
	Charset     string `yaml:"charset" env-default:"utf-8" env-description:"sets message charset" validate:"omitempty"`
	From        string `yaml:"from" env-description:"sets sender email" validate:"required,email"`
	Subject     string `yaml:"subject" env-description:"sets message subject" validate:"required"`
	MessageBody string `env-description:"sets message body" validate:"required"`
	ContentType string `env-description:"sets message content type" validate:"omitempty,oneof=text html"`
}

func NewMessageOptions() *MessageOptions {
	return &MessageOptions{}
}
