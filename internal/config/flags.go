/*
https://stackoverflow.com/questions/45487377/golang-flag-ignore-missing-flag-and-parse-multiple-duplicate-flags
*/
package config

import (
	"flag"
)

var (
// !! NOTE: Move flags definitions to var declaration
)

/*
func InitFlags(ctx context.Context, fc *models.VMailSend, fai *types.AuthItem, fal *types.AttachmentsList, frl *types.RecipientsList, fhl *models.HeadersList) error {

	// Body Subcommands for attachment in mail body
	bodyCmd := flag.NewFlagSet("body", flag.ExitOnError)
	bodyCmd.StringVar(&fc.Body.Message, FlagBodyMessageNameLong, FlagBodyMessageDefaultValue, FlagBodyMessageDescription)
	bodyCmd.StringVar(&fc.Body.Message, FlagBodyMessageNameLong, FlagBodyMessageDefaultValue, FlagBodyMessageDescription+" (shorthand)")
	bodyCmd.StringVar(&fc.Body.File, FlagBodyFileNameLong, FlagBodyFileDefaultValue, FlagBodyFileDescription)
	bodyCmd.StringVar(&fc.Body.File, FlagBodyFileNameShort, FlagBodyFileDefaultValue, FlagBodyFileDescription+" (shorthand)")
	bodyCmd.StringVar(&fc.Body.MimeType, FlagBodyMIMETypeNameLong, FlagBodyMIMETypeDefaultValue, FlagBodyMIMETypeDescription)
	bodyCmd.StringVar(&fc.Body.MimeType, FlagBodyMIMETypeNameShort, FlagBodyMIMETypeDefaultValue, FlagBodyMIMETypeDescription+" (shorthand)")

	// Header Subcommands. Repeat for multiple headers
	headerCmd := flag.NewFlagSet("header", flag.ExitOnError)
	headerCmd.StringToStringVarP(&fhl.HeaderName, FlagHeaderNameNameLong, FlagHeaderNameDefaultValue, FlagHeaderNameDescription)
	headerCmd.StringVar(&c.HeaderName, FlagHeaderNameNameShort, FlagHeaderNameDefaultValue, FlagHeaderNameDescription+" (shorthand)")
	headerCmd.StringVar(&c.HeaderValue, FlagHeaderValueNameLong, FlagHeaderValueDefaultValue, FlagHeaderValueDescription)
	headerCmd.StringVar(&c.HeaderValue, FlagHeaderValueNameShort, FlagHeaderValueDefaultValue, FlagHeaderValueDescription+" (shorthand)")

	// Header List Flags.
	flag.Var(fhl, FlagHeaderListNameLong, FlagHeaderListDescription)
	flag.Var(fhl, FlagHeaderListNameShort, FlagHeaderListDescription)

	// SMTP Auth Subcommands
	authCmd := flag.NewFlagSet("auth", flag.ExitOnError)
	authCmd.StringVar(&fai.Username, FlagAuthUsernameNameLong, FlagAuthUsernameDefaultValue, FlagAuthUsernameDescription)
	authCmd.StringVar(&fai.Username, FlagAuthUsernameNameShort, FlagAuthUsernameDefaultValue, FlagAuthUsernameDescription+" (shorthand)")
	authCmd.StringVar(&fai.Password, FlagAuthPasswordNameLong, FlagAuthPasswordDefaultValue, FlagAuthPasswordDescription)
	authCmd.StringVar(&fai.Password, FlagAuthPasswordNameShort, FlagAuthPasswordDefaultValue, FlagAuthPasswordDescription+" (shorthand)")

	// Attachments Subcommands.  Repeat for multiple attachments
	attachCmd := flag.NewFlagSet("attachments", flag.ExitOnError)
	attachCmd.StringVar(&fal.File, FlagAttachmentFileNameLong, FlagAttachmentFileDefaultValue, FlagAttachmentFileDescription)
	attachCmd.StringVar(&fal.File, FlagAttachmentFileNameShort, FlagAttachmentFileDefaultValue, FlagAttachmentFileDescription+" (shorthand)")
	attachCmd.StringVar(&fal.Name, FlagAttachmentNameNameLong, FlagAttachmentNameDefaultValue, FlagAttachmentNameDescription)
	attachCmd.StringVar(&fal.Name, FlagAttachmentNameNameShort, FlagAttachmentNameDefaultValue, FlagAttachmentNameDescription+" (shorthand)")
	attachCmd.BoolVar(&fal.Inline, FlagAttachmentModeInlineNameLong, FlagAttachmentModeInlineDefaultValue, FlagAttachmentModeInlineDescription)
	attachCmd.BoolVar(&fal.Inline, FlagAttachmentModeInlineNameShort, FlagAttachmentModeInlineDefaultValue, FlagAttachmentModeInlineDescription+" (shorthand)")
	attachCmd.StringVar(&fal.MIMEType, FlagAttachmentMIMETypeNameLong, FlagAttachmentMIMETypeDefaultValue, FlagAttachmentMIMETypeDescription)
	attachCmd.StringVar(&fal.MIMEType, FlagAttachmentMIMETypeNameShort, FlagAttachmentMIMETypeDefaultValue, FlagAttachmentMIMETypeDescription+" (shorthand)")

	// Recipients Subcommands.  Repeat for multiple recipients
	toCmd := flag.NewFlagSet("recipients", flag.ExitOnError)
	toCmd.StringVar(&frl.Address, FlagRecipientAddressNameLong, FlagRecipientAddressDefaultValue, FlagRecipientAddressDescription)
	toCmd.StringVar(&frl.Address, FlagRecipientAddressShort, FlagRecipientAddressDefaultValue, FlagRecipientAddressDescription+" (shorthand)")
	toCmd.StringVar(&frl.Name, FlagRecipientNameNameLong, FlagRecipientNameDefaultValue, FlagRecipientNameDescription)
	toCmd.StringVar(&frl.Name, FlagRecipientNameNameShort, FlagRecipientNameDefaultValue, FlagRecipientNameDescription+" (shorthand)")
	toCmd.StringVar(&frl.Mode, FlagRecipientModeNameLong, FlagRecipientModeDefaultValue, FlagRecipientModeDescription)
	toCmd.StringVar(&frl.Mode, FlagRecipientModeShort, FlagRecipientModeDefaultValue, FlagRecipientModeDescription+" (shorthand)")

	// Recipients List Flag
	flag.Var(frl, FlagRecipientListNameLong, FlagRecipientListDescription)
	flag.Var(frl, FlagRecipientListNameShort, FlagRecipientListDescription)

	// Recipients File Flag
	flag.StringVar(frlf, FlagRecipientsListFileNameLong, FlagRecipientsListFileDefaultValue, FlagRecipientsListFileDescription)
	flag.StringVar(frlf, FlagRecipientsListFileNameShort, FlagRecipientsListFileDefaultValue, FlagRecipientsListFileDescription+" (shorthand)")

	// Service Flags
	flag.StringVar(&configPath, FlagConfigPathNameLong, FlagConfigPathDefaultValue, FlagConfigPathDescription)
	flag.StringVar(&configPath, FlagConfigPathNameShort, FlagConfigPathDefaultValue, FlagConfigPathDescription+" (shorthand)")
	flag.BoolVar(&showHelp, FlagShowHelpNameLong, FlagShowHelpDefaultValue, FlagShowHelpDescription)
	flag.BoolVar(&showHelp, FlagShowHelpNameShort, FlagShowHelpDefaultValue, FlagShowHelpDescription+" (shorthand)")
	flag.BoolVar(&showSMTPInfo, FlagShowSMTPInfoNameLong, FlagShowSMTPInfoDefaultValue, FlagShowSMTPInfoDescription)
	flag.BoolVar(&showSMTPInfo, FlagShowSMTPInfoNameShort, FlagShowSMTPInfoDefaultValue, FlagShowSMTPInfoDescription+" (shorthand)")
	flag.BoolVar(&showExample, FlagShowExamplesNameLong, FlagShowExamplesDefaultValue, FlagShowExamplesDescription)
	flag.BoolVar(&showExample, FlagShowExamplesNameShort, FlagShowExamplesDefaultValue, FlagShowExamplesDescription+" (shorthand)")
	flag.BoolVar(&showVersion, FlagShowVersionNameLong, FlagShowVersionDefaultValue, FlagShowVersionDescription)
	flag.BoolVar(&showVersion, FlagShowVersionNameShort, FlagShowVersionDefaultValue, FlagShowVersionDescription+" (shorthand)")
	flag.BoolVar(&Debug, FlagDebugNameLong, FlagDebugDefaultValue, FlagDebugDescription)
	flag.BoolVar(&Debug, FlagDebugNameShort, FlagDebugDefaultValue, FlagDebugDescription+" (shorthand)")
	flag.BoolVar(&Verbose, FlagVerboseNameLong, FlagVerboseDefaultValue, FlagVerboseDescription)
	flag.BoolVar(&Verbose, FlagVerboseNameShort, FlagVerboseDefaultValue, FlagVerboseDescription+" (shorthand)")

	// App Config Flags
	flag.StringVar(&fc.AppConfig.SMTPServer, FlagSMTPServerNameLong, FlagSMTPServerDefaultValue, FlagSMTPServerDescription)
	flag.StringVar(&fc.AppConfig.SMTPServer, FlagSMTPServerNameShort, FlagSMTPServerDefaultValue, FlagSMTPServerDescription+" (shorthand)")
	flag.IntVar(&fc.AppConfig.SMTPPort, FlagSMTPPortNameLong, FlagSMTPPortDefaultValue, FlagSMTPPortDescription)
	flag.IntVar(&fc.AppConfig.SMTPPort, FlagSMTPPortNameShort, FlagSMTPPortDefaultValue, FlagSMTPPortDescription+" (shorthand)")
	flag.StringVar(&fc.AppConfig.Domain, FlagDomainNameLong, FlagDomainDefaultValue, FlagDomainDescription)
	flag.StringVar(&fc.AppConfig.Domain, FlagDomainNameShort, FlagDomainDefaultValue, FlagDomainDescription)
	flag.StringVar(&fc.AppConfig.FlagLogLevel, FlagDebugLogLevelNameLong, FlagDebugLogLevelDefaultValue, FlagDebugLogLevelDescription)
	flag.StringVar(&fc.AppConfig.FlagLogLevel, FlagDebugLogLevelNameShort, FlagDebugLogLevelDefaultValue, FlagDebugLogLevelDescription+" (shorthand)")
	flag.StringVar(&fc.AppConfig.FlagLogFormat, FlagDebugLogFormatNameLong, FlagDebugLogFormatDefaultValue, FlagDebugLogFormatDescription)
	flag.StringVar(&fc.AppConfig.FlagLogFormat, FlagDebugLogFormatNameLong, FlagDebugLogFormatDefaultValue, FlagDebugLogFormatDescription+" (shorthand)")
	flag.StringVar(&fc.AppConfig.FlagLogFile, FlagDebugLogFileNameLong, FlagDebugLogFileDefaultValue, FlagDebugLogFileDescription)
	flag.StringVar(&fc.AppConfig.FlagLogFile, FlagDebugLogFileNameShort, FlagDebugLogFileDefaultValue, FlagDebugLogFileDescription+" (shorthand)")
	flag.StringVar(&fc.AppConfig.AppLogMarker, FlagAppLogMarkerNameLong, FlagAppLogMarkerDefaultValue, FlagAppLogMarkerDescription)
	flag.StringVar(&fc.AppConfig.AppLogMarker, FlagAppLogMarkerNameShort, FlagAppLogMarkerDefaultValue, FlagAppLogMarkerDescription+" (shorthand)")
	flag.BoolVar(&fc.AppConfig.VerifyCert, FlagVerifyCertNameLong, FlagVerifyCertDefaultValue, FlagVerifyCertDescription)
	flag.BoolVar(&fc.AppConfig.VerifyCert, FlagVerifyCertNameShort, FlagVerifyCertDefaultValue, FlagVerifyCertDescription+" (shorthand)")
	flag.BoolVar(&fc.AppConfig.SSL, FlagSSLName, FlagSSLDefaultValue, FlagSSLDescription)

	// Message Config Flags
	flag.StringVar(&fc.MessageConfig.Charset, FlagCharsetNameLong, FlagCharsetDefaultValue, FlagCharsetDescription)
	flag.StringVar(&fc.MessageConfig.Charset, FlagCharsetNameShort, FlagCharsetDefaultValue, FlagCharsetDescription+" (shorthand)")

	flag.StringVar(&fc.MessageConfig.From, FlagSenderAddressNameLong, FlagSenderAddressDefaultValue, FlagSenderAddressDescription)
	flag.StringVar(&fc.MessageConfig.From, FlagSenderAddressNameShort, FlagSenderAddressDefaultValue, FlagSenderAddressDescription+" (shorthand)")
	flag.StringVar(&fc.MessageConfig.FromName, FlagSenderNameNameLong, FlagSenderNameDefaultValue, FlagSenderNameDescription)
	flag.StringVar(&fc.MessageConfig.FromName, FlagSenderNameNameShort, FlagSenderNameDefaultValue, FlagSenderNameDescription+" (shorthand)")

	flag.StringVar(&fc.MessageConfig.Subject, FlagSubjectNameLong, FlagSubjectDefaultValue, FlagSubjectDescription)
	flag.StringVar(&fc.MessageConfig.Subject, FlagSubjectNameShort, FlagSubjectDefaultValue, FlagSubjectDescription+" (shorthand)")

	// Parse Flags
	switch os.Args[1] {
	case "body":
		bodyCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'body'")
		fmt.Println("  Message:", fc.Body.Message)
		fmt.Println("  File:", fc.Body.File)
		fmt.Println("  MimeType:", fc.Body.MimeType)
		fmt.Println("  tail:", bodyCmd.Args())

	case "header":
		headerCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'body'")
		fmt.Println("  Message:", fc.Body.Message)

		headerCmd.StringVar(&c.HeaderName, FlagHeaderNameNameLong, FlagHeaderNameDefaultValue, FlagHeaderNameDescription)
		headerCmd.StringVar(&c.HeaderName, FlagHeaderNameNameShort, FlagHeaderNameDefaultValue, FlagHeaderNameDescription+" (shorthand)")
		headerCmd.StringVar(&c.HeaderValue, FlagHeaderValueNameLong, FlagHeaderValueDefaultValue, FlagHeaderValueDescription)
		headerCmd.StringVar(&c.HeaderValue, FlagHeaderValueNameShort, FlagHeaderValueDefaultValue, FlagHeaderValueDescription+" (shorthand)")

	case "auth":
		barCmd.Parse(os.Args[2:])
		fmt.Println("bar")
	case "recipients":
		barCmd.Parse(os.Args[2:])
		fmt.Println("bar")
	case "attachments":
		barCmd.Parse(os.Args[2:])
		fmt.Println("bar")
	default:
		log.Fatalf("[ERROR] unknown subcommand '%s', see help for more details.", os.Args[1])
	}

	flag.Parse()

	if frl != nil {
		fc.SetRecipientsListFromType(frl)
	}

	logging.L(ctx).Debug("GetConfig: flags parsed successfully")
}
*/

func IsFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
