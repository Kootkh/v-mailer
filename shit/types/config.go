package types

import (
	"context"
	"fmt"
	"strings"
	"v-mailer/internal/lib/logging"

	"github.com/caarlos0/env"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	AppConfig struct {
		SMTPDomain    string `yaml:"smtp-domain" env-default:"localhost" env-description:"sets smtp domain" validate:"omitempty,hostname"`
		SMTPPort      int    `yaml:"smtp-port" env-default:"587" env-description:"sets smtp port" validate:"omitempty,number,min=1,max=65535"`
		AppLogLevel   string `yaml:"log-level" env-default:"info" env-description:"sets log level" validate:"omitempty"`
		AppLogFormat  string `yaml:"log-format" env-default:"json" env-description:"sets log format" validate:"omitempty"`
		AppLogFile    string `yaml:"log-file" env-default:"./log/v-mailer.log" env-description:"sets log file" validate:"omitempty,filepath"`
		AppLogMarker  string `env-description:"sets app log marker" validate:"omitempty"`
		FlagLogLevel  string `env-description:"sets flag log level" validate:"omitempty"`
		FlagLogFormat string `env-description:"sets flag log format" validate:"omitempty,oneof=text json"`
		FlagLogFile   string `env-description:"sets log file" validate:"omitempty,filepath"`
		VerifyCert    bool   `yaml:"verify-cert" env-default:"enabled" env-description:"sets certificate verification" validate:"omitempty,boolean"`
		SSL           bool   `yaml:"ssl" env-default:"enabled" env-description:"sets ssl mode" validate:"omitempty,boolean"`
		Auth          struct {
			Username string `yaml:"username" env-description:"sets smtp auth username" validate:"omitempty,required_with=password"`
			Password string `yaml:"password" env-description:"sets smtp auth password" validate:"omitempty,required_with=Username"`
		} `yaml:"smtp-auth" validate:"required_with=Username,required_with=Password"`
		// } `yaml:"app-config" validate:"omitempty,unique,dive,optional"`
	} `yaml:"app-config" validate:"omitempty"`
	MessageConfig struct {
		Charset        string       `yaml:"charset" env-default:"utf-8" env-description:"sets message charset" validate:"omitempty"`
		From           string       `yaml:"from" env-description:"sets sender email" validate:"required,email"`
		Subject        string       `yaml:"subject" env-description:"sets message subject" validate:"required"`
		Attachments    []Attachment `env-description:"sets path to attachments" validate:"omitempty,dive,required"`
		RecipientsList string       `yaml:"default-recipients" validate:"omitempty"`
		Recipients     []string     `validate:"required,dive,required,email"`
		MessageBody    string       `env-description:"sets message body" validate:"required"`
		ContentType    string       `env-description:"sets message content type" validate:"omitempty,oneof=text html"`
		//} `yaml:"message-config" validate:"omitempty,unique,dive,optional"`
	} `yaml:"message-config" validate:"omitempty"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) GetAttachment(index int) Attachment {
	return c.MessageConfig.Attachments[index]
}

func (c *Config) AddAttachment(attachment Attachment) {
	c.MessageConfig.Attachments = append(c.MessageConfig.Attachments, attachment)
}

func (c *Config) GetRecipientsList() string {
	return c.MessageConfig.RecipientsList
}

func (c *Config) SetRecipientsListFromString(value string) {
	if len(c.MessageConfig.RecipientsList) == 0 {
		c.MessageConfig.RecipientsList = value
		return
	}
	c.MessageConfig.RecipientsList = c.MessageConfig.RecipientsList + "," + value
}

func (c *Config) SetRecipientsListFromType(recipientsList *RecipientsList) {
	for _, v := range *recipientsList {
		c.SetRecipientsListFromString(v)
	}
}

func (c *Config) LenRecipientsList() int {
	return len(c.MessageConfig.RecipientsList)
}

func (c *Config) CountRecipientsList() int {
	if len(c.MessageConfig.RecipientsList) == 0 {
		return 0
	}
	return len(strings.Split(c.MessageConfig.RecipientsList, ","))
}

func (c *Config) LenRecipients() int {
	return len(c.MessageConfig.Recipients)
}

func (c *Config) GetRecipient(index int) string {
	if index >= c.LenRecipients() {
		return ""
	}
	return c.MessageConfig.Recipients[index]
}

func (c *Config) AddRecipient(recipient string) {
	c.MessageConfig.Recipients = append(c.MessageConfig.Recipients, recipient)
}

func (ac *Config) ReadConfigFile(ctx context.Context, cf string) error {
	if err := cleanenv.ReadConfig(cf, ac); err != nil {
		logging.L(ctx).Error("GetConfig: can't read config file", logging.ErrAttr(err))
		return fmt.Errorf("GetConfig: can't read config file: %q, error: %w", cf, err)
	}
	return nil
}

func (ec *Config) ProcessEnvFile(ctx context.Context, file string) error {
	// если есть файл переменных окружения - пробуем загрузить из него параметры и распарсить их в конфиг "ec" (Config)
	/* 	if !util.FileExists(file) {
	   		logging.L(ctx).Warn("GetConfig: environment variables file is not exists", logging.StringAttr("file", file))
	   		return fmt.Errorf("GetConfig: environment variables file %q is not exists", file)
	   	}
	   	logging.L(ctx).Debug("GetConfig: environment variables file found - processing", logging.StringAttr("file", file)) */

	if err := godotenv.Load(file); err != nil {
		logging.L(ctx).Warn("GetConfig: unable to load environment file", logging.ErrAttr(err))
		return fmt.Errorf("GetConfig: unable to load environment file: %q, error: %w", file, err)
	}
	logging.L(ctx).Debug("GetConfig: environment variables successfully loaded form file to memory", logging.StringAttr("file", file))

	if err := env.Parse(ec); err != nil {
		logging.L(ctx).Error("GetConfig: unable to parse environment variables", logging.ErrAttr(err))
		return fmt.Errorf("GetConfig: unable to parse environment variables, error: %w", err)
	}
	logging.L(ctx).Debug("GetConfig: environment variables parsed from memory to environment config successfully")
	return nil
}

/* func (cc *Config) ProcessConfigFile(ctx context.Context, fcf, ecf string) error {

	// Если во флагах указан файл конфига, и он существует - пробуем прочитать и распарсить по структуре конфига "cc" (AltConfig). Если не получается - возвращаем ошибку.
	if fcf != "" && util.FileExists(fcf) {
		logging.L(ctx).Debug("GetConfig: reading config file from flag value", logging.StringAttr("config_file", fcf))
		if err := cc.ReadConfigFile(ctx, fcf); err != nil {
			logging.L(ctx).Error("GetConfig: failed to read config file from flag value", logging.ErrAttr(err))
			return fmt.Errorf("GetConfig: failed to read config file from flag value. file: %q, error: %w", fcf, err)
		}
		logging.L(ctx).Debug("GetConfig: config file from flag value readed successfully")
		return nil

		// Если во флагах указан файл конфига, и он не существует - возвращаем ошибку.
	} else if fcf != "" && !util.FileExists(fcf) {
		logging.L(ctx).Error("GetConfig: config file provided in flag is not exists", logging.StringAttr("file", fcf))
		return fmt.Errorf("GetConfig: config file provided in flag is not exists: %q", fcf)

		// Если в переменных окружения указан файл конфига, и он существует - пробуем прочитать и распарсить по структуре конфига "cc" (AltConfig). Если не получается - возвращаем ошибку.
	} else if ecf != "" && util.FileExists(ecf) {
		logging.L(ctx).Debug("GetConfig: reading config file from environment value", logging.StringAttr("config_file", ecf))
		if err := cc.ReadConfigFile(ctx, ecf); err != nil {
			logging.L(ctx).Error("GetConfig: failed to read config file from environment value", logging.ErrAttr(err))
			return fmt.Errorf("GetConfig: failed to read config file from environment value. file: %q, error: %w", ecf, err)
		}
		logging.L(ctx).Debug("GetConfig: config file from environment value readed successfully")
		return nil

		// Если в переменных окружения указан файл конфига, и он не существует - возвращаем ошибку.
	} else if ecf != "" && !util.FileExists(ecf) {
		logging.L(ctx).Error("GetConfig: specified in environment config file not exists", logging.StringAttr("file", fcf))
		return fmt.Errorf("GetConfig: specified in environment config file %q is not exists", ecf)

		// Если файл конфига не указан во флагах и в переменных окружения - выводим предупреждение об этом
	} else {
		logging.L(ctx).Warn("GetConfig: config file is not specified")
		return nil
	}
} */

/* func (ac *Config) ProcessFlagsAttachments(ctx context.Context, fal *AttachmentsList) error {

	ac.MessageConfig.Attachments = make([]Attachment, 0, len(*fal))

	// --------------------------------------------
	// если во флагах указаны приложения (fal) - пробуем заполнить структуру fc.MessageConfig.Attachments
	// --------------------------------------------
	for _, item := range *fal {
		item = strings.TrimSpace(item)

		if !util.FileExists(item) {
			logging.L(ctx).Error("GetConfig: attachment does not exist", logging.StringAttr("attachment", item))
			return fmt.Errorf("GetConfig: attachment %q is not exist", item)
		}

		for _, i := range ac.MessageConfig.Attachments {
			if i.AttachmentName == filepath.Base(item) {
				logging.L(ctx).Error("GetConfig: attachment already exist", logging.StringAttr("attachment", item))
				return fmt.Errorf("GetConfig: attachment %q already exist", item)
			}
		}

		ac.MessageConfig.Attachments = append(ac.MessageConfig.Attachments, Attachment{
			FilePath:       item,
			EncodingType:   "base64",
			AttachmentName: filepath.Base(item),
			MimeType:       mime.TypeByExtension(filepath.Ext(item)),
		})
		logging.L(ctx).Debug("GetConfig: attachment added to app config", logging.StringAttr("attachment", item))
	}

	return nil
} */
