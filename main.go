/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"text/tabwriter"
	"v-mailer/internal/config"
	"v-mailer/internal/flags"
	"v-mailer/internal/lib/logging"
	"v-mailer/internal/util"

	"github.com/yassinebenaid/godump"
)

// !! https://clavinjune.dev/en/blogs/ways-to-define-custom-command-line-flags-in-golang/

// "версия" в контекст
const (
	version = "0.0.1"
	// debug info string
	flagIsPassedString = "\tFlagIsPassed:\t"
)

var (
	// init config builder log settings
	logInitLevel  string    = slog.LevelError.String()
	logInitFormat string    = "text"
	logInitDest   io.Writer = os.Stdout
	// app log settings
	mainLogFile *os.File  = nil
	mainLogDest io.Writer = nil
	// flag log settings
	flagLogFile *os.File  = nil
	flagLogDest io.Writer = nil

	//err  error

)

func main() {

	ctx := context.Background()

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	//
	defer cancel()

	// ##########################################################################
	// Нормализация флагов (аргументов командной строки)
	// - Преобразовать флаги указанные "полным названием" в вид "нижний регистр" и заменить "kebab-like-case" на "snake_like_case".
	// - Т.к. "кобра" принимает все указанные в аргументах командной строки инстанции флага, но использует последнее указанное значение , нужно провалидироваться аргументом флага по списку уникальности флагов и крашнуться в случае множественного указания флага из списка уникальности.
	// ##########################################################################

	//fmt.Printf("Original os.Args: &%v\n", os.Args)
	err := error(nil)
	os.Args, err = util.NormalizeFlags(ctx, os.Args)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}

	//fmt.Printf("Normalized os.Args: &%v\n", os.Args)

	// Закидываем в переменные среды "версию"
	os.Setenv(version, version)

	// --------------------------------------------------------------------------
	//  Инициализируем логгер для сборщика конфигурации приложения
	// --------------------------------------------------------------------------

	// Если в переменных среды указан флаг "IS_DEV" - устанавливаем уровень логирования для логгера сборщика конфигурации в "debug"
	if os.Getenv("IS_DEV") == "true" {
		logInitLevel = slog.LevelDebug.String()
	}

	log := logging.NewLLogger(
		logging.WithLLevel(logInitLevel),
		logging.WithFormat(logInitFormat),
		logging.WithLAddSource(true),
		logging.WithLogDest(logInitDest),
		logging.WithLSetDefault(false),
	)
	// Закидываем логгер сборщика конфигурации приложения в контекст
	ctx = logging.ContextWithLogger(ctx, log)

	// --------------------------------------------------------------------------
	// Инициализируем валидатор
	// --------------------------------------------------------------------------
	val := util.Init()

	// Закидываем валидатор в контекст
	ctx = util.ContextWithValidator(ctx, val)

	// --------------------------------------------------------------------------
	// Инициализируем сборщик конфигурации приложения (app *types.Config)
	// --------------------------------------------------------------------------
	logging.L(ctx).Debug("main: config initializing...")

	app, err := config.GetConfig(ctx, os.Args)

	if err != nil {
		logging.L(ctx).Error("main: failed to init configuration", logging.ErrAttr(err))
		os.Exit(1)
	}

	// --------------------------------------------------------------------------
	// Проверяем валидность конфига приложения
	// --------------------------------------------------------------------------

	/*
		logging.L(ctx).Debug("main: app config validation...")
		v_validator.ValidateStruct(ctx, app)
		logging.L(ctx).Debug("main: app config validate successful...")
	*/

	// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	// !!!!! TODO: VALIDATE APP CONFIG !!!!!!
	// !!!!! TODO: deal with VALIDATOR !!!!!!
	// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

	/* 	if err := v_validator.L(ctx).Struct(app); err != nil {
		v_validator.ProcessValidationErrors(ctx, err)
	} */

	// --------------------------------------------------------------------------
	// Обрабатываем сервисные флаги (пример: --help)
	// --------------------------------------------------------------------------

	/*
		// "Show HELP information"
		if config.IsFlagPassed(config.FlagShowHelpNameLong) || config.IsFlagPassed(config.FlagShowHelpNameShort) { // IsFlagPassed
			flag.Usage()
			os.Exit(0)
		}
	*/

	/*
		// "Show SMTP information"
		if config.IsFlagPassed(config.FlagShowSMTPInfoNameLong) || config.IsFlagPassed(config.FlagShowSMTPInfoNameShort) {
			fmt.Println("-------------------------------------------------------------")
			fmt.Printf("SMTP INFO (domain:port)       | %s:%d\n", app.AppConfig.SMTPDomain, app.AppConfig.SMTPPort)
			fmt.Printf("SMTP AUTH (username@password) | %s@%s\n", app.AppConfig.Auth.Username, app.AppConfig.Auth.Password)
			fmt.Printf("SMTP SENDER (from)            | %s\n", app.MessageConfig.From)
			fmt.Printf("SMTP RECIPIENTS (to)          | %#v\n", app.MessageConfig.RecipientsList)
			fmt.Printf("SMTP SSL                      | %#v\n", app.AppConfig.SSL)
			fmt.Printf("SMTP Verify Certs             | %#v\n", app.AppConfig.VerifyCert)
			fmt.Println("-------------------------------------------------------------")
			os.Exit(0)
		}
	*/

	/*
		// "Show examples"
		// fsc.ShowExample
	*/

	/*
		// "Show version"
		if config.IsFlagPassed(config.FlagShowVersionNameLong) || config.IsFlagPassed(config.FlagShowVersionNameShort) {
			fmt.Printf("App Version: %s\n", version)
			os.Exit(0)
		}
	*/

	// --------------------------------------------------------------------------
	// Logger Initialization for application
	// --------------------------------------------------------------------------

	// !! TODO: переделать multiwriter.
	// !! TODO: По умолчанию весь вывод в logfile
	// !! TODO: в stdout пишем только при отладке(IS_DEV, debug)
	// !! TODO: Надо делать отдельный логгер для файла и stdout. !!!

	// !! TODO: добавить в флаги лог-файл для параллельного вывода.
	// !! TODO: flagLogFile добавить в multiwriter

	// ----------------------------------------------------------------
	// Logger Initialization for app
	// ----------------------------------------------------------------

	/*
		// Определяем основной файл лога - app.AppConfig.AppLogFile
		// Значение параметра не может быть пустым. Если в конфиг файле отсутствует - устанавливается дефолтное значение /var/log/v-mailer/v-mailer.log
		mainLogFile, err = os.OpenFile(app.AppConfig.AppLogFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			logging.L(ctx).Error("main: app log file error", logging.StringAttr("file", app.AppConfig.AppLogFile), logging.ErrAttr(err))
			os.Exit(1)
		}

		// Не забываем закрыть файл
		defer mainLogFile.Close()

		if config.Verbose {
			mainLogDest = io.MultiWriter(mainLogFile, os.Stdout)
		} else {
			mainLogDest = io.MultiWriter(mainLogFile)
		}

		// https://www.youtube.com/watch?v=ptoKy-COIlE

		appLogger := logging.NewLLogger(
			logging.WithLLevel(app.AppConfig.AppLogLevel),
			logging.WithFormat(app.AppConfig.AppLogFormat),
			logging.WithLAddSource(false),
			logging.WithLogDest(mainLogDest),
			logging.WithLSetDefault(false),
		)

		// Запихиваем логгер в контекст
		ctx = logging.ContextWithLogger(ctx, appLogger)

		// Если во флаге конфига указан маркер для логгера - добавляем его к основному логгеру
		if app.AppConfig.AppLogMarker != "" {
			appLogger = logging.WithAttrs(ctx, logging.StringAttr("marker", app.AppConfig.AppLogMarker))
		}

		// Обновляем логгер в контексте
		ctx = logging.ContextWithLogger(ctx, appLogger)
	*/

	// ----------------------------------------------------------------
	// Logger Initialization for flag
	// ----------------------------------------------------------------

	/*
		if app.AppConfig.FlagLogFile != "" {
			flagLogFile, err = os.OpenFile(app.AppConfig.FlagLogFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				logging.L(ctx).Error("main: flag log file error", logging.StringAttr("file", app.AppConfig.FlagLogFile), logging.ErrAttr(err))
				os.Exit(1)
			}
			// Не забываем закрыть файл
			defer flagLogFile.Close()
			flagLogDest = io.MultiWriter(flagLogFile, os.Stdout)
		} else {
			flagLogDest = io.MultiWriter(os.Stdout)
		}

		fLog := logging.NewLLogger(
			logging.WithLLevel(app.AppConfig.FlagLogLevel),
			logging.WithFormat(app.AppConfig.FlagLogFormat),
			logging.WithLAddSource(true),
			logging.WithLogDest(flagLogDest),
			logging.WithLSetDefault(true),
		)
	*/

	/*
		logging.L(ctx).Info("appLog: config initialized successfully",
			logging.Group("OS Info",
				logging.StringAttr("OS", runtime.GOOS),
				logging.StringAttr("ARCH", runtime.GOARCH),
				logging.IntAttr("NumCPU", runtime.NumCPU()),
				logging.StringAttr("Go Version", runtime.Version()),
			))

		logging.L(ctx).Debug("appLog: config initialized successfully",
			logging.Group("OS Info",
				logging.StringAttr("OS", runtime.GOOS),
				logging.StringAttr("ARCH", runtime.GOARCH),
				logging.IntAttr("NumCPU", runtime.NumCPU()),
				logging.StringAttr("Go Version", runtime.Version()),
			))

		fLog.Info("fLog: config initialized successfully",
			logging.Group("OS Info",
				logging.StringAttr("OS", runtime.GOOS),
				logging.StringAttr("ARCH", runtime.GOARCH),
				logging.IntAttr("NumCPU", runtime.NumCPU()),
				logging.StringAttr("Go Version", runtime.Version()),
			))

		fLog.Debug("fLog: config initialized successfully",
			logging.Group("OS Info",
				logging.StringAttr("OS", runtime.GOOS),
				logging.StringAttr("ARCH", runtime.GOARCH),
				logging.IntAttr("NumCPU", runtime.NumCPU()),
				logging.StringAttr("Go Version", runtime.Version()),
			))
	*/

	// !!!!! !!!!! !!!!! !!!!! !!!!! !!!!! !!!!! !!!!!
	// !!!!! TODO: Add "show config" flag to app !!!!!
	// !!!!! !!!!! !!!!! !!!!! !!!!! !!!!! !!!!! !!!!!

	// -----------------------------------------------------------
	// DEBUG
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, "Debug:\t", flags.FS.Debug, flagIsPassedString, flags.IsFlagPassed(flags.RootCmd, "debug"))
	fmt.Fprintln(w, "Verbose:\t", flags.FS.Verbose, flagIsPassedString, flags.IsFlagPassed(flags.RootCmd, "verbose"))
	fmt.Fprintln(w, "Version:\t", flags.FS.ShowVersion, flagIsPassedString, flags.IsFlagPassed(flags.RootCmd, "version"))
	fmt.Fprintln(w, "Examples:\t", flags.FS.ShowExamples, flagIsPassedString, flags.IsFlagPassed(flags.RootCmd, "examples"))
	fmt.Fprintln(w, "Help:\t", flags.FS.ShowHelp, flagIsPassedString, flags.IsFlagPassed(flags.RootCmd, "help"))

	w.Flush()

	fmt.Printf("Len of FHMap: %v\n", len(flags.FHMap))
	fmt.Printf("fhMap: %v\n", flags.FHMap)
	fmt.Printf("Len of Recipients: %v\n", len(flags.FC.Recipients))
	fmt.Printf("SMTP Server address: %v\n", flags.FC.AppOptions.SMTPServer)
	fmt.Printf("SMTP Server zone: %v\n", flags.FC.AppOptions.SMTPServerZone)
	fmt.Printf("SMTP Server port: %v\n", flags.FC.AppOptions.SMTPServerPort)
	fmt.Printf("From domain: %v\n", flags.FC.AppOptions.Domain)
	fmt.Printf("Subject: %v\n", flags.FC.AppOptions.Subject)

	for i := range flags.FC.Recipients {
		rt, err := flags.FC.GetRecipient(i)
		if err != nil {
			fmt.Printf("GetRecipient error: %v\n", err)
		}
		fmt.Printf("Recipient: %v\n", flags.FC.GetAllRecipientValues(*rt))
	}

	if app != nil {
		godump.Dump(*app)
	}

	logging.L(ctx).Debug("main finished.")
	os.Exit(0)
}
