package config

import (
	"context"
	"fmt"

	"v-mailer/internal/flags"
	"v-mailer/internal/lib/logging"
	"v-mailer/internal/models"
)

// Vars for service config params
var (
	ShowHelp     bool
	ShowSMTPInfo bool
	ShowExamples bool
	ShowVersion  bool
	Debug        bool
	Verbose      bool
	configPath   string
)

var (
	envFile string = ".env"
)

// func GetConfig(ctx context.Context) (*types.Config, error) {
func GetConfig(ctx context.Context, args []string) (*models.VMailSend, error) {
	logging.L(ctx).Debug("GetConfig: starting...")

	// Flags Config Instance
	// var fc = types.NewConfig()
	/* var fc = models.NewVMailSend() */

	// Environment Config Instance
	/* var ec = types.NewConfig() */

	// Conf File Config Instance
	// var cc = types.NewConfig()
	/* var cc = models.NewVMailSend() */

	// App Config Instance (returned by GetConfig)
	// var ac = types.NewConfig()
	/* var ac = models.NewVMailSend() */

	// Item for flags specified SMTP Auth "username" and "password"
	//var fai = types.NewAuthItem()
	/* var fai = models.NewAuth() */

	// Item for flags specified attachments
	/* var fal = models.NewAttachmentsList() */

	// Item for flags specified recipients
	/* var frl = models.NewRecipientsList() */

	// Flags Headers List
	/* fhl := models.NewHeadersList() */

	// --------------------------------------------
	// Парсим значения принятых флагов в конфиг "fc"
	// --------------------------------------------
	//InitFlags(ctx, fc, fai, fal, frl, fhl)

	// Запускаем обработку флагов Cobra
	flags.Execute(ctx)

	fmt.Printf("flag config: %#v\n", flags.FS.ConfigPath)
	fmt.Printf("Sender Address: %#v\n", flags.FC.AppOptions.From)
	fmt.Printf("Sender Name: %#v\n", flags.FC.AppOptions.FromName)
	//fmt.Printf("flag config: %#v\n", flags.FC.ServiceOptions.ConfigPath)

	// --------------------------------------------
	// пробуем загрузить и распарсить файл переменных окружения в конфиг "ec"
	// --------------------------------------------
	/*
		if err := ec.ProcessEnvFile(ctx, envFile); err != nil {
			logging.L(ctx).Warn("GetConfig: unable to process environment variables file", logging.ErrAttr(err))
			// return nil, err
		} else {
			logging.L(ctx).Debug("GetConfig: environment file processed successfully")
		}
	*/

	// --------------------------------------------
	// пробуем обработать файл конфига в конфиг "cc"
	// --------------------------------------------
	/*
		if err := cc.ProcessConfigFile(ctx, configPath, os.Getenv(EnvConfigPathName)); err != nil {
			logging.L(ctx).Warn("GetConfig: unable to process config file", logging.ErrAttr(err))
			return nil, err
		} else {
			logging.L(ctx).Debug("GetConfig: config file processed successfully")
		}
	*/

	// --------------------------------------------
	// пробуем обработать данные для smtp авторизации из флагов
	// --------------------------------------------
	/*
		if err := checkFlagAuth(ctx, fc, fai); err != nil {
			return nil, err
		}
	*/

	// --------------------------------------------
	// пробуем обработать данные для smtp авторизации из файла конфигурации
	// --------------------------------------------
	/*
		if err := checkFileConfigAuth(ctx, cc); err != nil {
			return nil, err
		}
	*/

	// --------------------------------------------
	// Настройки основного файла лога принимаем только из файла конфигурации.
	// Если в файле конфигурации указан файл лога (cc.AppConfig.AppLogFile) - закидываем его в ac.AppConfig.AppLogFile.
	// Иначе закидываем в ac.AppConfig.AppLogFile дефолтное значение.
	// --------------------------------------------
	/*
		if cc.AppConfig.AppLogFile != "" {
					ac.AppConfig.AppLogFile = cc.AppConfig.AppLogFile
				} else {
					ac.AppConfig.AppLogFile = "./log/v-mailer.log"
				}

				if cc.AppConfig.AppLogFormat != "" {
					ac.AppConfig.AppLogFormat = cc.AppConfig.AppLogFormat
				} else {
					ac.AppConfig.AppLogFormat = "json"
				}

				if cc.AppConfig.AppLogLevel != "" {
					ac.AppConfig.AppLogLevel = cc.AppConfig.AppLogLevel
				} else {
					ac.AppConfig.AppLogLevel = "info"
		}
	*/

	// --------------------------------------------
	// Получателей принимаем из флагов и файла конфигурации.
	// Если во флагах указаны получатели (frl) - пробуем заполнить структуру fc.Recipients.
	// Иначе, если в файле конфигурации указаны получатели (cc.RecipientsList) - пробуем заполнить структуру cc.Recipients
	// --------------------------------------------
	/*
		if fc.CountRecipientsList() > 0 {
			fc.MessageConfig.Recipients = make([]string, 0, fc.CountRecipientsList())
			for _, item := range *frl {
				item = strings.TrimSpace(item)
				fc.MessageConfig.Recipients = append(fc.MessageConfig.Recipients, item)
				logging.L(ctx).Debug("GetConfig: flag recipient value added to flag config", logging.StringAttr("email", item))
			}
		} else if cc.CountRecipientsList() > 0 {
			cc.MessageConfig.Recipients = make([]string, 0, cc.CountRecipientsList())
			arr := strings.Split(cc.GetRecipientsList(), ",")
			for _, item := range arr {
				item = strings.TrimSpace(item)
				cc.MessageConfig.Recipients = append(cc.MessageConfig.Recipients, item)
				logging.L(ctx).Debug("GetConfig: config file recipient added to config file config", logging.StringAttr("email", item))
			}
		}
	*/

	// --------------------------------------------
	// Аттачменты принимаем только из флагов.
	// Если во флагах указаны аттачменты (fal) - пробуем заполнить структуру ac.MessageConfig.Attachments
	// --------------------------------------------
	/*
		if len(*fal) > 0 {
			if err := ac.ProcessFlagsAttachments(ctx, fal); err != nil {
				logging.L(ctx).Error("GetConfig: failed to process flags attachments", logging.ErrAttr(err))
				return nil, err
			}
		}
	*/

	// --------------------------------------------
	// Пробуем объединить получившиеся конфиги. (ac = fc > ec > cc)
	// --------------------------------------------
	/*
		logging.L(ctx).Debug("GetConfig: merging config params")

		if err := mergeConfigs(ctx, ac, fc, cc); err != nil {
			logging.L(ctx).Error("GetConfig: failed to merge config files", logging.ErrAttr(err))
			return nil, err
		}

		logging.L(ctx).Debug("GetConfig: merge configs processed successfully.")
	*/

	// --------------------------------------------
	// Возвращаем итоговый конфиг и нулевую ошибку
	// --------------------------------------------
	/* return ac, nil */
	return nil, nil
}

/* func checkFlagAuth(ctx context.Context, fc *mode.Config, fai *types.AuthItem) error {
	// Check if flag auth credentials are set.

	//v_validator.ValidateStruct(ctx, fai)

	if fai.Username != "" && fai.Password == "" {
		logging.L(ctx).Error("GetConfig: flags smtp auth password is not specified.", logging.StringAttr("username", fai.Username))
		return fmt.Errorf("GetConfig: flags smtp auth password is not specified")
	} else if fai.Password != "" && fai.Username == "" {
		logging.L(ctx).Error("GetConfig: flags smtp auth username is not specified")
		return fmt.Errorf("flags smtp auth username is not specified")
	} else if fai.Username == "" && fai.Password == "" {
		logging.L(ctx).Debug("GetConfig: flags smtp auth credentials not specified")
		return nil
	} else if fai.Username != "" && fai.Password != "" {
		fc.AppConfig.Auth.Username = strings.TrimSpace(fai.Username)
		fc.AppConfig.Auth.Password = strings.TrimSpace(fai.Password)
		logging.L(ctx).Debug("GetConfig: smtp auth credentials from flags values stored in flags config")
	}
	return nil
} */

/* func checkFileConfigAuth(ctx context.Context, cc *types.Config) error {
	if cc.AppConfig.Auth.Username != "" && cc.AppConfig.Auth.Password == "" {
		logging.L(ctx).Error("GetConfig: config smtp auth password is not specified", logging.StringAttr("username", cc.AppConfig.Auth.Username))
		return fmt.Errorf("GetConfig: config smtp auth password is not specified. username: %q", cc.AppConfig.Auth.Username)
	} else if cc.AppConfig.Auth.Password != "" && cc.AppConfig.Auth.Username == "" {
		logging.L(ctx).Error("GetConfig: config smtp auth username is not specified", logging.StringAttr("username", cc.AppConfig.Auth.Username))
		return fmt.Errorf("GetConfig: config smtp auth username is not specified. username: %q", cc.AppConfig.Auth.Username)
	} else if cc.AppConfig.Auth.Username == "" && cc.AppConfig.Auth.Password == "" {
		logging.L(ctx).Debug("GetConfig: config smtp auth credentials not specified")
		return nil
	}
	cc.AppConfig.Auth.Username = strings.TrimSpace(cc.AppConfig.Auth.Username)
	cc.AppConfig.Auth.Password = strings.TrimSpace(cc.AppConfig.Auth.Password)
	logging.L(ctx).Debug("GetConfig: smtp auth credentials from config file values stored in config file config", logging.StringAttr("username", cc.AppConfig.Auth.Username), logging.StringAttr("password", cc.AppConfig.Auth.Password))
	return nil
} */
