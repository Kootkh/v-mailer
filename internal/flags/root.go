/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package flags

import (
	"context"
	"fmt"
	"os"
	"strings"
	"v-mailer/internal/models"
	"v-mailer/internal/util"

	"github.com/spf13/cobra"
)

var (

	// Flags Config Instance
	FC *models.VMailSend = models.NewVMailSend()

	// Service Options
	// FS *models.ServiceOptions = models.NewServiceOptions()
	FS = models.NewServiceOptions()

	// App Options
	FA = models.NewAppOptions()

	// Flags Auth Item (specified SMTP Auth "username" and "password")
	FAI *models.Auth = models.NewAuth()

	// Flags Attachments List
	FAL *models.AttachmentsList = models.NewAttachmentsList()

	// Flags Recipients List
	FRL models.RecipientsList = models.NewRecipientsList()

	// Flags Recipient item
	FRI string

	// Flags Headers List
	FHL *models.HeadersList = models.NewHeadersList()

	FHMap map[string]string

	// flagCount = make(map[string]int)

	//flagDebugIsSet bool
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "v-mailer",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },

	/*
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			flagCount = make(map[string]int)
		},
	*/

	PreRun: func(cmd *cobra.Command, args []string) {

		// fmt.Printf("cmd: %v\n", cmd)
		// fmt.Printf("flags: %#v\n", cmd.Flags())
		// fmt.Printf("args: %v\n", args)

		// #############
		// Do not show usage on errors
		// #############
		cmd.SilenceUsage = true

		// #############
		// Manage declared SMTP Auth Credentials Flags.
		// Make flags required if FlagAuthUsername or FlagAuthPassword specified.
		// #############
		FAI.Username, _ = cmd.Flags().GetString(FlagAuthUsernameName)
		if FAI.Username != "" {
			cmd.MarkFlagRequired(FlagAuthPasswordName)
		}

		FAI.Password, _ = cmd.Flags().GetString(FlagAuthPasswordName)
		if FAI.Password != "" {
			cmd.MarkFlagRequired(FlagAuthUsernameName)
		}

	},

	RunE: func(cmd *cobra.Command, args []string) error {

		// ==================================================================
		// Manage declared Service Options Flags
		// ==================================================================

		// #############
		// Manage declared "Verbose" Flag
		// #############
		verbosity, _ := cmd.Flags().GetCount(FlagVerboseName)

		if verbosity > 0 {
			if verbosity > 5 {
				verbosity = 5
			}
			FS.Verbose = true
			FS.Verbosity = verbosity
			fmt.Printf("verbosity: %d\n", FS.Verbosity)
		}

		// #############
		// Manage declared "Recipient" Flag.
		// #############
		recipientsArray, _ := cmd.Flags().GetStringArray(FlagRecipientName)
		if len(recipientsArray) > 0 {
			recipientsList := strings.Join(recipientsArray, ";")
			if err := FC.SetRecipients(strings.TrimSpace(recipientsList)); err != nil {
				return err
			}
		}

		// #############
		// Manage declared "Recipients List" Flags
		// #############
		recipientsListFileArray, _ := cmd.Flags().GetStringArray(FlagRecipientsListFileName)
		if len(recipientsListFileArray) > 0 {
			recipientsListFiles := strings.Join(recipientsListFileArray, ";")
			if err := ProcessRecipientsListFilesFlag(strings.TrimSpace(recipientsListFiles)); err != nil {
				return err
			}
		}

		// #############
		// Manage declared "Sender" Flag
		// #############
		if sender, _ := cmd.Flags().GetString(FlagSenderName); sender != "" {
			FC.AppOptions.SetSender(strings.TrimSpace(sender))
		}

		// #############
		// Manage declared "SMTP Server" Flag
		// #############
		if server, _ := cmd.Flags().GetString(FlagSMTPServerName); server != "" {
			FC.AppOptions.SetSMTPServer(strings.ToLower(strings.TrimSpace(server)))
		}

		// #############
		// Manage declared "Domain" Flag
		// #############
		if domain, _ := cmd.Flags().GetString(FlagDomainName); domain != "" {
			FC.AppOptions.SetDomain(strings.ToLower(strings.TrimSpace(domain)))
		}

		// #############
		// Manage declared "Subject" Flag
		// #############
		if subject, _ := cmd.Flags().GetString(FlagSubjectName); subject != "" {
			FC.AppOptions.SetSubject(strings.TrimSpace(subject))
		}

		// Get all flags
		// cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		// 	fmt.Println("Flag:", flag.Name, "Value:", flag.Value)
		// })

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(ctx context.Context) {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	// -----------------------------------------------------------
}

/* func addSubCommands(cmd *cobra.Command) {
	for _, subCmd := range cmd.Commands() {
		addSubCommands(subCmd)
	}
} */

func init() {

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.v-mailer.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// rootCmd.PersistentFlags().BoolVarP(&config.Debug, "debug", "d", false, "Print debug messages")

	// -----------------------------------------------------------
	// Declare "Service Options" Flags
	// -----------------------------------------------------------

	// Declare "Debug" Flag
	RootCmd.Flags().BoolVarP(&FS.Debug, FlagDebugName, FlagDebugShorthand, FlagDebugDefaultValue, FlagDebugDescription)

	// Declare "Verbose" & "verbosity" Flags
	RootCmd.Flags().CountP(FlagVerboseName, FlagVerboseShorthand, FlagVerboseDescription)

	// Declare "Version" Flag
	RootCmd.Flags().BoolVarP(&FS.ShowVersion, FlagShowVersionName, FlagShowVersionShorthand, FlagShowVersionDefaultValue, FlagShowVersionDescription)

	// Declare "Examples" Flag
	RootCmd.Flags().BoolVarP(&FS.ShowExamples, FlagShowExamplesName, FlagShowExamplesShorthand, FlagShowExamplesDefaultValue, FlagShowExamplesDescription)

	// Declare "Help" Flag
	RootCmd.Flags().BoolVarP(&FS.ShowHelp, FlagShowHelpName, FlagShowHelpShorthand, FlagShowHelpDefaultValue, FlagShowHelpDescription)

	// Declare "SMTP Info" Flag
	RootCmd.Flags().BoolVarP(&FS.ShowSMTPInfo, FlagShowSMTPInfoName, FlagShowSMTPInfoShorthand, FlagShowSMTPInfoDefaultValue, FlagShowSMTPInfoDescription)

	// Declare "Config Path" Flag
	RootCmd.Flags().StringVarP(&FS.ConfigPath, FlagConfigPathName, FlagConfigPathShorthand, FlagConfigPathDefaultValue, FlagConfigPathDescription)

	// -----------------------------------------------------------
	// Declare "Application Options" Flags
	// -----------------------------------------------------------

	/*
		v Copyright                bool	// FlagCopyrightName
		v Ipv4                     bool // FlagIpv4Name
		v Ipv6                     bool // FlagIpv6Name
		v SMTPServer               string // FlagSMTPServerName
		  SMTPServerZone           string // FlagSMTPServerName
		  SMTPServerPort           int // FlagSMTPServerName
		v Domain                   string	// FlagDomainName
		Subject                  string
		FromName                 string
		From                     string
		MessageBody              string
		Name                     string
		ReplyToAddress           string	// reply to address
		RequestReadReciptAddress string
		ReturnPathAddress        string
		SSL                      bool
		VerifyCert               bool	// Verify Certificate in connection. Default is No
		CharacterSet             string
	*/

	// Declare "Copyright" Flag
	RootCmd.Flags().BoolVarP(&FC.AppOptions.Copyright, FlagCopyrightName, FlagSSLShorthand, FlagSSLDefaultValue, FlagSSLDescription)

	// Declare "Ipv4" & "Ipv6" Flags
	RootCmd.Flags().BoolVarP(&FC.AppOptions.Ipv4, FlagIpv4Name, FlagIpv4Shorthand, FlagIpv4DefaultValue, FlagIpv4Description)
	RootCmd.Flags().BoolVarP(&FC.AppOptions.Ipv6, FlagIpv6Name, FlagIpv6Shorthand, FlagIpv6DefaultValue, FlagIpv6Description)

	// Declare "SMTP Server" Flag
	RootCmd.Flags().StringP(FlagSMTPServerName, FlagSMTPServerShorthand, FlagSMTPServerDefaultValue, FlagSMTPServerDescription)

	// Declare "Domain" Flag
	RootCmd.Flags().StringP(FlagDomainName, FlagDomainShorthand, FlagDomainDefaultValue, FlagDomainDescription)

	// Declare "Sender" Flag
	RootCmd.Flags().StringP(FlagSenderName, FlagSenderShorthand, FlagSenderDefaultValue, FlagSenderDescription)

	// Declare "Subject" Flag
	RootCmd.Flags().StringP(FlagSubjectName, FlagSubjectShorthand, FlagSubjectDefaultValue, FlagSubjectDescription)

	// ==============================

	RootCmd.Flags().BoolVarP(&FC.AppOptions.SSL, FlagSSLName, FlagSSLShorthand, FlagSSLDefaultValue, FlagSSLDescription)

	// Declare "Headers List" Flags.
	RootCmd.Flags().StringToStringVarP(&FHMap, FlagHeadersName, FlagHeadersShorthand, nil, FlagHeadersDescription)
	//viper.BindPFlag(FlagHeadersName, RootCmd.Flags().Lookup(FlagHeadersName))

	// Declare "SMTP Authentication Credentials" Flags.
	RootCmd.Flags().StringP(FlagAuthUsernameName, FlagAuthUsernameShorthand, FlagAuthUsernameDefaultValue, FlagAuthUsernameDescription)
	//RootCmd.MarkFlagRequired(FlagAuthUsernameName)
	RootCmd.Flags().StringP(FlagAuthPasswordName, FlagAuthPasswordShorthand, FlagAuthPasswordDefaultValue, FlagAuthPasswordDescription)
	//RootCmd.MarkFlagRequired(FlagAuthPasswordName)

	// Declare "Body Options" Flags.
	RootCmd.Flags().StringVarP(&FC.Body.Message, FlagBodyMessageName, FlagBodyMessageShorthand, FlagBodyMessageDefaultValue, FlagBodyMessageDescription)
	RootCmd.Flags().StringVarP(&FC.Body.File, FlagBodyFileName, FlagBodyFileShorthand, FlagBodyFileDefaultValue, FlagBodyFileDescription)
	RootCmd.Flags().StringVarP(&FC.Body.MimeType, FlagBodyMimeTypeName, FlagBodyMimeTypeShorthand, FlagBodyMimeTypeDefaultValue, FlagBodyMimeTypeDescription)

	// Declare "Recipient" Flag. Repeat for multiple recipients.
	RootCmd.Flags().StringArrayP(FlagRecipientName, FlagRecipientShorthand, nil, FlagRecipientDescription)

	// Declare "Recipients List" Flag. Repeat for multiple files.
	RootCmd.Flags().StringArrayP(FlagRecipientsListFileName, FlagRecipientsListFileShorthand, nil, FlagRecipientsListFileDescription)

	/*
		FlagDebugLogLevelName = "log_level"
		FlagDebugLogLevelShorthand = ""
		FlagDebugLogLevelDefaultValue = "debug"
		FlagDebugLogLevelDescription = "set debug log level"

		FlagDebugLogFormatName = "log_format"
		FlagDebugLogFormatShorthand = ""
		FlagDebugLogFormatDefaultValue = "text"
		FlagDebugLogFormatDescription = "set debug log format"

		FlagDebugLogFileName = "log_file"
		FlagDebugLogFileShorthand = ""
		FlagDebugLogFileDefaultValue = ""
		FlagDebugLogFileDescription = "write log messages to this file"
		FlagAppLogMarkerName = "log_marker"
		FlagAppLogMarkerShorthand = "m"
		FlagAppLogMarkerDefaultValue = ""
		FlagAppLogMarkerDescription = "set app log marker"
		FlagVerifyCertName = "verify_cert"
		FlagVerifyCertShorthand = "cert"
		FlagVerifyCertDefaultValue = false
		FlagVerifyCertDescription = "Verify Certificate in connection. Default is No"
	*/

	/*
		// Attachments Subcommands.  Repeat for multiple attachments
		attachCmd := flag.NewFlagSet("attachments", flag.ExitOnError)
		attachCmd.StringVar(&fal.File, FlagAttachmentFileName, FlagAttachmentFileDefaultValue, FlagAttachmentFileDescription)
		attachCmd.StringVar(&fal.File, FlagAttachmentFileShorthand, FlagAttachmentFileDefaultValue, FlagAttachmentFileDescription+" (shorthand)")
		attachCmd.StringVar(&fal.Name, FlagAttachmentNameName, FlagAttachmentNameDefaultValue, FlagAttachmentNameDescription)
		attachCmd.StringVar(&fal.Name, FlagAttachmentNameShorthand, FlagAttachmentNameDefaultValue, FlagAttachmentNameDescription+" (shorthand)")
		attachCmd.BoolVar(&fal.Inline, FlagAttachmentModeInlineName, FlagAttachmentModeInlineDefaultValue, FlagAttachmentModeInlineDescription)
		attachCmd.BoolVar(&fal.Inline, FlagAttachmentModeInlineShorthand, FlagAttachmentModeInlineDefaultValue, FlagAttachmentModeInlineDescription+" (shorthand)")
		attachCmd.StringVar(&fal.MIMEType, FlagAttachmentMIMETypeName, FlagAttachmentMIMETypeDefaultValue, FlagAttachmentMIMETypeDescription)
		attachCmd.StringVar(&fal.MIMEType, FlagAttachmentMIMETypeShorthand, FlagAttachmentMIMETypeDefaultValue, FlagAttachmentMIMETypeDescription+" (shorthand)")





		// App Config Flags
		flag.StringVar(&fc.AppConfig.FlagLogLevel, FlagDebugLogLevelName, FlagDebugLogLevelDefaultValue, FlagDebugLogLevelDescription)
		flag.StringVar(&fc.AppConfig.FlagLogLevel, FlagDebugLogLevelShorthand, FlagDebugLogLevelDefaultValue, FlagDebugLogLevelDescription+" (shorthand)")
		flag.StringVar(&fc.AppConfig.FlagLogFormat, FlagDebugLogFormatName, FlagDebugLogFormatDefaultValue, FlagDebugLogFormatDescription)
		flag.StringVar(&fc.AppConfig.FlagLogFormat, FlagDebugLogFormatName, FlagDebugLogFormatDefaultValue, FlagDebugLogFormatDescription+" (shorthand)")
		flag.StringVar(&fc.AppConfig.FlagLogFile, FlagDebugLogFileName, FlagDebugLogFileDefaultValue, FlagDebugLogFileDescription)
		flag.StringVar(&fc.AppConfig.FlagLogFile, FlagDebugLogFileShorthand, FlagDebugLogFileDefaultValue, FlagDebugLogFileDescription+" (shorthand)")
		flag.StringVar(&fc.AppConfig.AppLogMarker, FlagAppLogMarkerName, FlagAppLogMarkerDefaultValue, FlagAppLogMarkerDescription)
		flag.StringVar(&fc.AppConfig.AppLogMarker, FlagAppLogMarkerShorthand, FlagAppLogMarkerDefaultValue, FlagAppLogMarkerDescription+" (shorthand)")
		flag.BoolVar(&fc.AppConfig.VerifyCert, FlagVerifyCertName, FlagVerifyCertDefaultValue, FlagVerifyCertDescription)
		flag.BoolVar(&fc.AppConfig.VerifyCert, FlagVerifyCertShorthand, FlagVerifyCertDefaultValue, FlagVerifyCertDescription+" (shorthand)")
		flag.BoolVar(&fc.AppConfig.SSL, FlagSSLName, FlagSSLDefaultValue, FlagSSLDescription)

		// Message Config Flags
		flag.StringVar(&fc.MessageConfig.Charset, FlagCharsetName, FlagCharsetDefaultValue, FlagCharsetDescription)
		flag.StringVar(&fc.MessageConfig.Charset, FlagCharsetShorthand, FlagCharsetDefaultValue, FlagCharsetDescription+" (shorthand)")



		flag.StringVar(&fc.MessageConfig.Subject, FlagSubjectName, FlagSubjectDefaultValue, FlagSubjectDescription)
		flag.StringVar(&fc.MessageConfig.Subject, FlagSubjectShorthand, FlagSubjectDefaultValue, FlagSubjectDescription+" (shorthand)")



		FlagCommandAuthName = "auth" // auth command for SMTP authentication
		FlagAuthUsernameName = "username"
		FlagAuthUsernameShorthand = "n"
		FlagAuthUsernameDefaultValue = ""
		FlagAuthUsernameDescription = "username for SMTP authentication. Required"
		FlagAuthPasswordName = "password"
		FlagAuthPasswordShorthand = "p"
		FlagAuthPasswordDefaultValue = ""
		FlagAuthPasswordDescription = "password for SMTP authentication. Required"
		FlagCharsetName = "charset"
		FlagCharsetShorthand = ""
		FlagCharsetDefaultValue = "utf-8"
		FlagCharsetDescription = "Character set for text/HTML."
		FlagSenderNameName = "from_name"
		FlagSenderNameShorthand = ""
		FlagSenderNameDefaultValue = ""
		FlagSenderNameDescription = "name of sender"
		FlagSenderAddressName = "from"
		FlagSenderAddressShorthand = "f"
		FlagSenderAddressDefaultValue = ""
		FlagSenderAddressDescription = "set sender email address"
		FlagSubjectName = "subject"
		FlagSubjectShorthand = "s"
		FlagSubjectDefaultValue = ""
		FlagSubjectDescription = "set message subject"
		FlagCommandBodyName = "body" //body command for attachment for mail body
		FlagBodyMessageName = "message"
		FlagBodyMessageShorthand = "m"
		FlagBodyMessageDefaultValue = ""
		FlagBodyMessageDescription = "message to show as body"
		FlagBodyFileName = "file"
		FlagBodyFileShorthand = "f"
		FlagBodyFileDefaultValue = ""
		FlagBodyFileDescription = "path to file of a text/HTML file to attach as body"
		FlagBodyMIMETypeName = "mime-type"
		FlagBodyMIMETypeShorthand = "mime"
		FlagBodyMIMETypeDefaultValue = ""
		FlagBodyMIMETypeDescription = "MIME type of the file to attach as body. Default is auto-detected"
		FlagCommandRecipientName = "to" // recipient command. Repeat for multiple recipients
		FlagRecipientAddressName = "email"
		FlagRecipientAddressShort = "a"
		FlagRecipientAddressDefaultValue = ""
		FlagRecipientAddressDescription = "email address of the recipient"
		FlagRecipientNameName = "name"
		FlagRecipientNameShorthand = "n"
		FlagRecipientNameDefaultValue = ""
		FlagRecipientNameDescription = "name of the recipient"
		FlagRecipientModeName = "mode"
		FlagRecipientModeShort = "m"
		FlagRecipientModeDefaultValue = ""
		FlagRecipientModeDescription = "recipients address send mode `to/cc/bcc`"
		FlagRecipientsListFileName = "recipients_file"
		FlagRecipientsListFileShorthand = ""
		FlagRecipientsListFileDefaultValue = ""
		FlagRecipientsListFileDescription = "csv file with list of 'name,email address,mode'.	Syntax is: Name, email_address, mode"
		FlagRecipientListName = "recipients_list"
		FlagRecipientListShorthand = ""
		FlagRecipientListDefaultValue = ""
		FlagRecipientListDescription = "list of 'email addresses/names/modes'.	Syntax is: email_address, name, mode"
		FlagCommandAttachmentName = "attach" // attach command. Repeat for multiple attachments
		FlagAttachmentFileName = "file"
		FlagAttachmentFileShorthand = "f"
		FlagAttachmentFileDefaultValue = ""
		FlagAttachmentFileDescription = "path to attachment file"
		FlagAttachmentNameName = "name"
		FlagAttachmentNameShorthand = "n"
		FlagAttachmentNameDefaultValue = ""
		FlagAttachmentNameDescription = "name of attachment. Default is filename"
		FlagAttachmentModeInlineName = "inline"
		FlagAttachmentModeInlineShorthand = "i"
		FlagAttachmentModeInlineDefaultValue = false
		FlagAttachmentModeInlineDescription = "set attachment content-disposition to `inline` or `attachment`. Default is `false` = `attachment`"
		FlagAttachmentMIMETypeName = "mime-type"
		FlagAttachmentMIMETypeShorthand = ""
		FlagAttachmentMIMETypeDefaultValue = ""
		FlagAttachmentMIMETypeDescription = "MIME type of the file to attach as body. Default is auto-detected"
		FlagHeaderListName = "headers_list"
		FlagHeaderListShorthand = ""
		FlagHeaderListDefaultValue = ""
		FlagHeaderListDescription = "list of 'name,value'.	Syntax is: Name, value;Name, value"
		FlagCommandHeaderName = "header" // header command. Repeat for multiple headers.
		FlagHeaderNameName = "name"
		FlagHeaderNameShorthand = "n"
		FlagHeaderNameDefaultValue = ""
		FlagHeaderNameDescription = "Header name"
		FlagHeaderValueName = "value"
		FlagHeaderValueShorthand = "v"
		FlagHeaderValueDefaultValue = ""
		FlagHeaderValueDescription = "Header value"
	*/
}

/* func IsFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
} */

func IsFlagPassed(cmd *cobra.Command, flagName string) bool {
	flagName, shorthand := GetFlagNameAndShorthand(cmd, flagName)
	// fmt.Println("Name:", flagName)
	// fmt.Println("Shorthand:", shorthand)

	flag := cmd.Flags().Lookup(flagName)
	flagShort := cmd.Flags().Lookup(shorthand)

	if flag != nil {
		return flag.Changed
	} else if flagShort != nil {
		return flagShort.Changed
	}
	return false
}

func GetFlagNameAndShorthand(cmd *cobra.Command, flagName string) (string, string) {
	flag := cmd.Flags().Lookup(flagName)
	if flag != nil {
		return flag.Name, flag.Shorthand
	}
	return "", ""
}

func ProcessRecipientsListFilesFlag(value string) error {

	// проверка на дубликаты
	duplicates := util.CheckStringForDuplicates(strings.TrimSpace(value))
	if len(duplicates) > 0 {
		return fmt.Errorf("duplicate entries found in recipients list: %v", duplicates)
	}

	// fmt.Printf("recipients list: %v\n", value)

	// разбираем принятую строку на элементы по ";"
	for _, entry := range strings.Split(strings.TrimSpace(value), ";") {

		// удаляем лишние пробелы у каждого элемента entry
		entry = strings.TrimSpace(entry)

		// если файл не существует - Возвращаем ошибку!
		if _, err := util.FileExists(entry); err != nil {
			return err
		}

		file, err := util.ReadFile(entry)
		if err != nil {
			// Не смогли прочитать файл - Возвращаем ошибку!
			return err
		}

		FC.ParseRecipientsListFile(file)
	}
	return nil
}
