package models

import (
	"fmt"
	"strings"
	"v-mailer/internal/util"
)

type AppOptions struct {
	Copyright                bool
	Ipv4                     bool
	Ipv6                     bool
	SMTPServer               string `validate:"string,required,N/A,-smtp"`
	SMTPServerZone           string
	SMTPServerPort           int    `validate:"number,optional,587,1,65535,-port"`
	Domain                   string `validate:"string,optional,localhost,-domain"` // domain name for SMTP HELO/EHLO. HELO адрес — это символическое имя вашего компьютера, которое используется при соединении с SMTP сервером. Default is localhost
	Subject                  string
	FromName                 string
	From                     string `validate:"string,required,N/A,-from"`
	MessageBody              string
	Name                     string
	ReplyToAddress           string // reply to address
	RequestReadReciptAddress string
	ReturnPathAddress        string
	SSL                      bool
	VerifyCert               bool // Verify Certificate in connection. Default is No
	CharacterSet             string
}

func NewAppOptions() *AppOptions {
	return &AppOptions{}
}

func (a *AppOptions) SetSender(sender string) error {

	// new

	const errMsg = "sender: [%v], invalid format"

	var (
		address, name string
		// Раскладываем принятую строку sender на элементы по ";" (удаляя пустые элементы).
		elements = util.RemoveEmptyElementsFromArray(strings.Split(sender, ";"))
	)

	// Если элементов более одного - возвращаем ошибку
	if len(elements) > 1 {
		return fmt.Errorf(errMsg, sender)
	}

	// раскладываем нулевой элемент "отправителя" на элементы по "," (удаляя пустые элементы)
	elements = util.RemoveEmptyElementsFromArray(strings.Split(elements[0], ","))

	// Свитчимся по количеству элементов
	switch len(elements) {

	// Если элементов 1 или 2
	case 1, 2:

		// Если нулевой элемент не является электронной почтой - возвращаем ошибку
		if !util.CheckIfEmail(elements[0]) {
			return fmt.Errorf(errMsg, sender)
		}

		// В противном случае - запоминаем "адрес" из нулевого элемента
		address = elements[0]

		if len(elements) == 2 {

			// Раскладываем первый элемент sender на элементы по " "(удаляя пустые элементы) и проходим по ним циклом
			for _, element := range util.RemoveEmptyElementsFromArray(strings.Split(elements[1], " ")) {
				element = strings.TrimSpace(element)
				// если элемент является электронной почтой - возвращаем ошибку
				if util.CheckIfEmail(element) {
					return fmt.Errorf("sender: [%v], duplicate email address: [%v]", sender, element)
				}
				util.SetName(&name, element)
			}

		}

	default:
		// В остальных случаях - возвращаем ошибку
		return fmt.Errorf(errMsg, sender)
	}

	a.From, a.FromName = address, name
	return nil
}

func (a *AppOptions) SetSMTPServer(server string) error {

	const errMsg = "smtp_server: [%v], invalid format"

	address, zone, port, err := util.DetermineAddress(server)

	//fmt.Printf("address: [%v], zone: [%v], port: [%v]\n", address, zone, port)
	if err != nil {
		return fmt.Errorf(errMsg, server)
	}

	a.SMTPServer, a.SMTPServerZone, a.SMTPServerPort = address, zone, port

	return nil

}

func (a *AppOptions) SetDomain(domain string) error {

	const errMsg = "domain: [%v], invalid format"

	if util.CheckInvalidChars(domain, "domain") {
		return fmt.Errorf(errMsg, domain)
	}

	a.Domain = domain
	return nil
}

func (a *AppOptions) SetSubject(subject string) error {

	const errMsg = "subject: [%v], invalid"

	if util.CheckInvalidChars(subject, "subject") {
		fmt.Printf("SetSubject.subject: [%v], invalid\n", subject)
		return fmt.Errorf(errMsg, subject)
	}

	a.Subject = subject
	return nil
}
