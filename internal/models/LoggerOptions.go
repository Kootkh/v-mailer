package models

type AppLoggerOpts struct {
	AppLogLevel     string `yaml:"log-level" env-default:"info" env-description:"sets log level" validate:"omitempty"`
	AppLogFormat    string `yaml:"log-format" env-default:"json" env-description:"sets log format" validate:"omitempty,oneof=text json"`
	AppLogFile      string `yaml:"log-file" env-default:"./log/v-mailer.log" env-description:"sets log file" validate:"omitempty,filepath"`
	AppLogAddSource bool   `yaml:"log-add-source" env-default:"true" env-description:"add source to log" validate:"omitempty,boolean"`
	AppLogMarker    string `env-description:"sets app log marker" validate:"omitempty"`
}

func NewAppLoggerOpts() *AppLoggerOpts {
	return &AppLoggerOpts{}
}

type DebugLoggerOpts struct {
	DebugLogLevel     string `env-description:"sets debug log level" validate:"omitempty"`
	DebugLogFormat    string `env-description:"sets debug log format" validate:"omitempty,oneof=text json"`
	DebugLogFile      string `env-description:"sets debug log file" validate:"omitempty,filepath"`
	DebugLogAddSource bool   `env-description:"add source to debug log" validate:"omitempty,boolean"`
}

func NewDebugLoggerOpts() *DebugLoggerOpts {
	return &DebugLoggerOpts{}
}
