package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Recipients - Описание структуры для объекта списка адресов "Recipients" ...
type Recipient struct {
	Address string `validate:"required,email"`
	Name    string `validate:"omitempty"`
	Mode    string `yaml:"send-mode" validate:"omitempty,oneof=to cc bcc"`
}

func NewRecipient() *Recipient {
	recipient := new(Recipient)
	return recipient
}

func (r *Recipient) Append(address, name, mode string) error {
	r.Address = address
	r.Name = name
	r.Mode = mode
	return nil
}

type RecipientsList []string

func NewRecipientsList() RecipientsList {
	return RecipientsList{}
}

func (rl RecipientsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(rl)
}

func (rl RecipientsList) MarshalText() ([]byte, error) {
	return json.Marshal(rl)
}

func (rl RecipientsList) UnmarshalText(b []byte) error {
	for _, str := range strings.Split(string(b), ",") {
		rl = append(rl, str)
	}
	return nil
}

func (rl RecipientsList) Strings() []string {
	return rl
}

func (rl RecipientsList) Add(value string) {
	rl = append(rl, strings.TrimSpace(value))
}

func (rl RecipientsList) String() string {
	b, _ := json.Marshal(rl)
	return string(b)
}

func (rl RecipientsList) Set(value string) error {
	for _, str := range strings.Split(value, ";") {
		fmt.Println(str)
		rl = append(rl, strings.TrimSpace(str))
	}
	return nil
}
