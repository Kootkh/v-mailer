package types

import "strings"

type Recipient string

func (r *Recipient) String() string {
	return string(*r)
}

func (r *Recipient) Set(value string) {
	*r = Recipient(strings.TrimSpace(value))
}

type RecipientsList []string

func NewRecipientsList() *RecipientsList {
	return &RecipientsList{}
}

func (rl *RecipientsList) Add(value string) {
	*rl = append(*rl, strings.TrimSpace(value))
}

func (rl *RecipientsList) String() string {
	return ""
}

func (rl *RecipientsList) Set(value string) error {
	arr := strings.Split(value, ",")

	for _, v := range arr {
		*rl = append(*rl, strings.TrimSpace(v))
	}

	return nil
}
