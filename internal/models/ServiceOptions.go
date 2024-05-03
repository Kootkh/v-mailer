package models

// ServiceOptions - Описание структуры для объекта списка сервисных параметров. Параметры доступны только для флаговых команд.
type ServiceOptions struct {
	ShowHelp     bool
	ShowSMTPInfo bool
	ShowExamples bool
	ShowVersion  bool
	Debug        bool
	Verbose      bool
	Verbosity    int
	Quiet        bool
	ConfigPath   string
}

func NewServiceOptions() *ServiceOptions {
	/* 	service := new(ServiceOptions)
	   	return service */
	return &ServiceOptions{}
}
