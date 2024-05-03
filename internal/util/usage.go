package util

/* func ShowUsageAndExit() {
	v := " Version: @($) v-mailer v" + os.Getenv("version")
	usage := ` v-mailsend [options]
  Where the options are:
  -debug                 		- Print debug messages
		-d
	-verbose               		- show version and exit
		-v
  -subject <subject> 		 		- Subject
		-s="<subject>"
	to												- Recipient Command. Repeat for multiple recipients
		-email="<email>"				- email address of the recipient. Required
			-a="<email>"
		-name="<name>"					- name of the recipient. Optional
			-n="<name>"
		-mode="<mode>"					- recipients address send mode <to/cc/bcc>.
			-m="<mode>"							Optional. Default: <to>
  -recipients_file="<file>"	- csv file with list of 'name,email address,mode'.
		-rfile="<file>"						Syntax is: <Name, email_address, mode>"
	-recipients_list="<list>" - list of 'email addresses/names/modes'.
		-rlist="<list>"						Syntax is: -rlist <email,email,email...>"
	-from="<email>"*					- email address of the sender. Required
		-f="<email>"
  -from_name="<name>"				- name of sender
		-fname="<name>"
	-cc_recipients_list="<>"	- list of the carbon copy recipient addresses.
		-ccrlist="<list>"					Syntax is: -ccrlist="email,email,email..."
	cc												- Recipient carbon copy Command. Repeat for
															multiple carbon copy recipients
		-email="<email>"				- email address of the carbon copy recipient.
			-a="<email>"
		-name="<name>"					- name of the carbon copy recipient.
			-n="<name>"
	-bcc_recipients_list="<>"	- list of the blind carbon copy recipient addresses
		-bccrlist="<list>"				Syntax is: -bccrlist="email,email,email..."
	bcc												- Recipient blind carbon copy Command. Repeat for
															multiple blind carbon copy recipients
		-email="<email>"				- email address of the blind carbon copy recipient.
			-a="<email>"
		-name="<name>"					- name of the blind carbon copy recipient.
				-n="<name>"
	-reply_to="<email>"				- reply to email address
		-rt="<email>"
	-smtp_server="<host/IP>"	- hostname/IP address of the SMTP server. Required
		-ss="<host/IP>"
	-smtp_port="<port>"				- port of SMTP server. Default is 587
		-sp=<port>
  -domain="<domain>"				- domain name for SMTP HELO. Default is localhost
		-dmn="<domain>"
	-info											- Print info about SMTP server
		-i
	-ssl											- SMTP over SSL. Default is StartTLS
	-verify_cert							- Verify Certificate in connection.
		-cert											Default is "false".
  -examples									- show examples
		-ex
	-help											- show this help
		-h
	-log_file="<file>"				- write debug messages to this file
		-lfile="<file>"
	-log_format="<format>"		- set debug messages format <text|json>.
		-lformat="<format>				Default is <text>
	-log_level="<level>"			- set debug messages level <debug|info|warn|error>.
		-llevel="<level>"					Default is <debug>
	-log_marker="<marker>"    - set app log messages marker
		-mark="<marker>"
	-charset="<charset>"			- Character set for text/HTML. Default is utf-8
		-cs="<charset>"
  auth											- Auth Command for SMTP authentication
		-username="<username>"*	- username for SMTP authentication. Required
			-name="<username>"*
		-password="<password>"*	- password for SMTP authentication. Required
			-pass="<password>"*
  body											- body command for attachment for mail body
		-message="<msg>"				- message to show as body
			-msg="<msg>"
		-file="<file>"					- path to file of a text/HTML file to attach
			-f="<file>"							as body
		-mime-type="<type>"			- MIME type of the body content.
			-mime="<type>"					Default is auto-detected
	attach										- attach command. Repeat for multiple attachments
		-file="<path>"*					- path of the attachment. Required
			-f="<path>"*
		-name="<name>"					- Name of the attachment.
			-n="<name>"							Default is filename
		-mime-type="<type>"			- MIME-Type of the attachment.
			-mime="<type>"					Default is detected
		-inline									- set attachment content-disposition to "inline" or
			-i											"attachment". Default is "false"="attachment".
	header										- Header Command. Repeat for multiple headers.
		-name="<header>"				- Header name
			-n="<header>"
		-value="<value>"				- Header value
			-v="<value>"

The options with * are required.

Environment variables:
	SMTP_USER_PASS for auth password (-pass)
`

	usage = strings.Replace(usage, "\t", "    ", -1)
	fmt.Printf("%s\n\n%s\n", v, usage)
	os.Exit(0)
}
*/
