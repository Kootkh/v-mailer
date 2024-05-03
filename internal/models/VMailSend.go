package models

import (
	"fmt"
	"reflect"
	"strings"
	"v-mailer/internal/util"
)

type VMailSend struct {
	AppOptions      AppOptions
	MessageOptions  MessageOptions
	Auth            Auth
	Body            Body
	Headers         []Header
	Attachments     []Attachment `env-description:"sets path to attachments" validate:"omitempty,dive,required"`
	Recipients      []Recipient  `validate:"required,dive,required,email"`
	CCRecipients    []Recipient  `validate:"required,dive,required,email"`
	BCCRecipients   []Recipient  `validate:"required,dive,required,email"`
	AppLoggerOpts   AppLoggerOpts
	DebugLoggerOpts DebugLoggerOpts
	//ServiceOptions  ServiceOptions
}

func NewVMailSend() *VMailSend {
	return &VMailSend{
		AppOptions:      *NewAppOptions(),
		MessageOptions:  *NewMessageOptions(),
		Auth:            *NewAuth(),
		Body:            *NewBody(),
		Headers:         []Header{},
		Attachments:     []Attachment{},
		Recipients:      []Recipient{},
		CCRecipients:    []Recipient{},
		BCCRecipients:   []Recipient{},
		AppLoggerOpts:   *NewAppLoggerOpts(),
		DebugLoggerOpts: *NewDebugLoggerOpts(),
		//ServiceOptions:  *NewServiceOptions(),
	}
}

// SetRecipients sets the recipients for the VMailSend object.
//
// It takes a string value and parses it to create Recipient objects, then adds them to the Recipients array.
// It returns an error.
func (c *VMailSend) SetRecipients(value string) error {

	// Разбираем принятую строку value на элементы entry по ";"
	for _, entry := range strings.Split(value, ";") {
		recipient := NewRecipient()

		// Разбираем каждый элемент entry по "," на item-ы
		item := strings.Split(strings.TrimSpace(entry), ",")
		//fmt.Printf("splitted items: %v\n", len(item))

		// Если item меньше или равен 1 и первый элемент пуст - пропускаем запись
		if len(item) < 2 && item[0] == "" {
			continue
		}

		// Нормализация параметров получателя
		address, name, mode, err := util.NormalizeRecipient(item)

		// Если нормализация не удалась - пропускаем запись
		if err != nil {
			fmt.Printf("normalize error: %v\n", err)
			continue
		}

		// Если нормализация удалась - создаем объект Recipient с нормализованными данными
		*recipient = Recipient{
			Address: address,
			Name:    name,
			Mode:    mode,
		}

		// Проверка entry на дублирование
		exists := false
		for _, person := range c.Recipients {
			if person.Address == strings.TrimSpace(item[0]) {
				exists = true
				break
			}
		}
		// Если дубликат - пропускаем entry
		if exists {
			fmt.Printf("The address %q already exists in the array. Skipping entry: %#q\n", strings.TrimSpace(item[0]), *recipient)
			//break
			continue
		}
		// Если не дубликат - добавляем получателя в массив
		c.Recipients = append(c.Recipients, *recipient)
	}

	return nil
}

func (c *VMailSend) ParseRecipientsListFile(recipients [][]string) error {

	// Проходим циклом по каждой прочитанной строке
	for _, recipient := range recipients {

		fmt.Printf("\nrecipient:\t\t%v\n", recipient)

		recipient := strings.Join(util.RemoveWhitespacesFromArray(recipient), ",")

		// Т.к. после строки может быть комментарий - разбираем принятую строку recipient по ";"
		for i, recipientItem := range strings.Split(recipient, ";") {

			recipientItem = strings.TrimSpace(recipientItem)
			//fmt.Printf("recipientItem:\t\t[%v]\n", recipientItem)
			// удаляем лишние пробелы у каждого получателя
			//entry := strings.Join(util.RemoveWhitespacesFromArray(recipientItem), ",")

			//fmt.Printf("recipient entry:\t[%v]\n", entry)

			// Пропускаем пустые строки и комментарии
			//if util.IsComment(entry) || len(entry) == 0 {
			if util.CheckIsStringCommented(recipientItem) || len(recipientItem) == 0 {
				//fmt.Println("\tSKIP comment")
				fmt.Printf("skip comment(%v):\t[%v]\n", i+1, recipientItem)
				continue
			}

			fmt.Printf("recipientItem(%v):\t[%v]\n", i+1, recipientItem)

			//
			//if err := c.SetRecipients(entry); err != nil {
			if err := c.SetRecipients(recipientItem); err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *VMailSend) LenRecipients() int {
	return len(c.Recipients)
}

func (c *VMailSend) GetRecipient(i int) (Recipient *Recipient, err error) {
	if i < 0 || i >= len(c.Recipients) {
		return nil, fmt.Errorf("index out of range")
	}
	return &c.Recipients[i], nil
}

func (c *VMailSend) GetAllRecipientValues(p Recipient) map[string]interface{} {
	paramValues := make(map[string]interface{})

	val := reflect.ValueOf(p)
	typ := reflect.TypeOf(p)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		paramName := typ.Field(i).Name
		paramValues[paramName] = field.Interface()
	}

	return paramValues
}
