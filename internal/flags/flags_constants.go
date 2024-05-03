package flags

const (
	// https://gobyexample.com/command-line-subcommands
	// https://abhinavg.net/2022/08/13/flag-subcommand/

	// -----------------
	// Service Options Flags. Flag only values.
	// -----------------
	FlagDebugName                = "debug"
	FlagDebugShorthand           = ""
	FlagDebugDefaultValue        = false
	FlagDebugDescription         = "set debug mode"
	FlagVerboseName              = "verbose"
	FlagVerboseShorthand         = "v"
	FlagVerboseDefaultValue      = false
	FlagVerboseDescription       = "set verbose mode"
	FlagShowVersionName          = "version"
	FlagShowVersionShorthand     = "V"
	FlagShowVersionDefaultValue  = false
	FlagShowVersionDescription   = "show app version and exit"
	FlagShowExamplesName         = "examples"
	FlagShowExamplesShorthand    = "e"
	FlagShowExamplesDefaultValue = false
	FlagShowExamplesDescription  = "show examples"
	FlagShowHelpName             = "help"
	FlagShowHelpShorthand        = "h"
	FlagShowHelpDefaultValue     = false
	FlagShowHelpDescription      = "show help"
	FlagShowSMTPInfoName         = "info"
	FlagShowSMTPInfoShorthand    = "i"
	FlagShowSMTPInfoDefaultValue = false
	FlagShowSMTPInfoDescription  = "Print info about SMTP server"
	FlagConfigPathName           = "config_file"
	FlagConfigPathShorthand      = ""
	FlagConfigPathDefaultValue   = "./app/config/config.yaml"
	FlagConfigPathDescription    = "set app config file"

	// -----------------
	// Application Options Flags. Flag & Config values.
	// -----------------
	FlagCopyrightName          = "copyright"
	FlagCopyrightShorthand     = ""
	FlagCopyrightDefaultValue  = false
	FlagCopyrightDescription   = "set copyright"
	FlagIpv4Name               = "ipv4"
	FlagIpv4Shorthand          = ""
	FlagIpv4DefaultValue       = true
	FlagIpv4Description        = "use ipv4 for SMTP server"
	FlagIpv6Name               = "ipv6"
	FlagIpv6Shorthand          = ""
	FlagIpv6DefaultValue       = false
	FlagIpv6Description        = "use ipv6 for SMTP server"
	FlagSMTPServerName         = "smtp_server"
	FlagSMTPServerShorthand    = ""
	FlagSMTPServerDefaultValue = ""
	FlagSMTPServerDescription  = "set smtp server"
	FlagDomainName             = "from_domain"
	FlagDomainShorthand        = ""
	FlagDomainDefaultValue     = "localhost"
	FlagDomainDescription      = "domain name for SMTP HELO."
	FlagSenderName             = "from"
	FlagSenderShorthand        = ""
	FlagSenderDefaultValue     = ""
	FlagSenderDescription      = "set sender email address & name. Syntax is: email_address, name"
	FlagSubjectName            = "subject"
	FlagSubjectShorthand       = "s"
	FlagSubjectDefaultValue    = ""
	FlagSubjectDescription     = "set message subject"

	FlagDebugLogLevelName         = "log_level"
	FlagDebugLogLevelShorthand    = ""
	FlagDebugLogLevelDefaultValue = "debug"
	FlagDebugLogLevelDescription  = "set debug log level"

	FlagDebugLogFormatName         = "log_format"
	FlagDebugLogFormatShorthand    = ""
	FlagDebugLogFormatDefaultValue = "text"
	FlagDebugLogFormatDescription  = "set debug log format"

	FlagDebugLogFileName         = "log_file"
	FlagDebugLogFileShorthand    = ""
	FlagDebugLogFileDefaultValue = ""
	FlagDebugLogFileDescription  = "write log messages to this file"

	FlagAppLogMarkerName         = "log_marker"
	FlagAppLogMarkerShorthand    = "m"
	FlagAppLogMarkerDefaultValue = ""
	FlagAppLogMarkerDescription  = "set app log marker"

	FlagVerifyCertName         = "verify_cert"
	FlagVerifyCertShorthand    = "cert"
	FlagVerifyCertDefaultValue = false
	FlagVerifyCertDescription  = "Verify Certificate in connection. Default is No"

	FlagSSLName         = "ssl"
	FlagSSLShorthand    = ""
	FlagSSLDefaultValue = true
	FlagSSLDescription  = "SMTP over SSL. Default is StartTLS"

	FlagAuthUsernameName         = "username"
	FlagAuthUsernameShorthand    = "u"
	FlagAuthUsernameDefaultValue = ""
	FlagAuthUsernameDescription  = "username for SMTP authentication. Required"

	FlagAuthPasswordName         = "password"
	FlagAuthPasswordShorthand    = "p"
	FlagAuthPasswordDefaultValue = ""
	FlagAuthPasswordDescription  = "password for SMTP authentication. Required"

	FlagCharsetName         = "charset"
	FlagCharsetShorthand    = ""
	FlagCharsetDefaultValue = "utf-8"
	FlagCharsetDescription  = "Character set for text/HTML."

	FlagBodyMessageName         = "message"
	FlagBodyMessageShorthand    = "m"
	FlagBodyMessageDefaultValue = ""
	FlagBodyMessageDescription  = "message to show as body"

	FlagBodyFileName             = "file"
	FlagBodyFileShorthand        = "f"
	FlagBodyFileDefaultValue     = ""
	FlagBodyFileDescription      = "path to file of a text/HTML file to attach as body"
	FlagBodyMimeTypeName         = "mime-type"
	FlagBodyMimeTypeShorthand    = ""
	FlagBodyMimeTypeDefaultValue = ""
	FlagBodyMimeTypeDescription  = "MIME type of the file to attach as body. Default is auto-detected"

	FlagRecipientName         = "to"
	FlagRecipientShorthand    = "t"
	FlagRecipientDefaultValue = ""
	FlagRecipientDescription  = "email address of the recipient"

	FlagRecipientsListFileName         = "recipients_file"
	FlagRecipientsListFileShorthand    = ""
	FlagRecipientsListFileDefaultValue = ""
	FlagRecipientsListFileDescription  = "csv file with list of 'name,email address,mode'.	Syntax is: Name, email_address, mode"

	FlagRecipientsListName         = "recipients_list"
	FlagRecipientsListShorthand    = ""
	FlagRecipientsListDefaultValue = ""
	FlagRecipientsListDescription  = "list of 'email addresses/names/modes'.	Syntax is: email_address, name, mode"

	FlagAttachmentFileName         = "attach"
	FlagAttachmentFileShorthand    = "a"
	FlagAttachmentFileDefaultValue = ""
	FlagAttachmentFileDescription  = "path to attachment file"

	FlagAttachmentNameName         = "name"
	FlagAttachmentNameShorthand    = "n"
	FlagAttachmentNameDefaultValue = ""
	FlagAttachmentNameDescription  = "name of attachment. Default is filename"

	FlagAttachmentModeInlineName         = "inline"
	FlagAttachmentModeInlineShorthand    = "i"
	FlagAttachmentModeInlineDefaultValue = false
	FlagAttachmentModeInlineDescription  = "set attachment content-disposition to `inline` or `attachment`. Default is `false` = `attachment`"

	FlagAttachmentMIMETypeName         = "mime-type"
	FlagAttachmentMIMETypeShorthand    = ""
	FlagAttachmentMIMETypeDefaultValue = ""
	FlagAttachmentMIMETypeDescription  = "MIME type of the file to attach as body. Default is auto-detected"

	FlagHeaderListName         = "headers_list"
	FlagHeaderListShorthand    = ""
	FlagHeaderListDefaultValue = ""
	FlagHeaderListDescription  = "list of 'name,value'.	Syntax is: Name, value;Name, value"

	FlagHeadersName         = "headers"
	FlagHeadersShorthand    = ""
	FlagHeadersDefaultValue = ""
	FlagHeadersDescription  = "Headers list.	Syntax is: --headers name1=value1,name2=value2"
)
