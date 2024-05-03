package config

/*
func mergeConfigs(ctx context.Context, ac, fc, cc *types.Config) error {

	// Merge AppConfig
	ac.AppConfig.SMTPDomain = mergeStrings(ac.AppConfig.SMTPDomain, fc.AppConfig.SMTPDomain, cc.AppConfig.SMTPDomain)
	ac.AppConfig.SMTPPort = mergeInts(ac.AppConfig.SMTPPort, fc.AppConfig.SMTPPort, cc.AppConfig.SMTPPort)

	// Merge App Logs Settings
	ac.AppConfig.AppLogFile = mergeStrings(ac.AppConfig.AppLogFile, cc.AppConfig.AppLogFile)
	ac.AppConfig.AppLogFormat = mergeStrings(ac.AppConfig.AppLogFormat, cc.AppConfig.AppLogFormat)
	ac.AppConfig.AppLogLevel = mergeStrings(ac.AppConfig.AppLogLevel, cc.AppConfig.AppLogLevel)
	ac.AppConfig.AppLogMarker = mergeStrings(ac.AppConfig.AppLogMarker, fc.AppConfig.AppLogMarker)

	// Merge Flag Logs Settings
	ac.AppConfig.FlagLogFile = mergeStrings(ac.AppConfig.FlagLogFile, fc.AppConfig.FlagLogFile)
	ac.AppConfig.FlagLogFormat = mergeStrings(ac.AppConfig.FlagLogFormat, fc.AppConfig.FlagLogFormat)
	ac.AppConfig.FlagLogLevel = mergeStrings(ac.AppConfig.FlagLogLevel, fc.AppConfig.FlagLogLevel)

	ac.AppConfig.VerifyCert = *mergeBool(ctx, FlagVerifyCertName, fc.AppConfig.VerifyCert, cc.AppConfig.VerifyCert)
	ac.AppConfig.SSL = *mergeBool(ctx, FlagSSLName, fc.AppConfig.SSL, cc.AppConfig.SSL)

	// Merge AppConfig SMTP Auth
	if err := mergeAuth(ctx, ac, fc, cc); err != nil {
		return err
	}

	// Merge MessageConfig
	ac.MessageConfig.Charset = mergeStrings(ac.MessageConfig.Charset, fc.MessageConfig.Charset, cc.MessageConfig.Charset)
	ac.MessageConfig.From = mergeStrings(ac.MessageConfig.From, fc.MessageConfig.From, cc.MessageConfig.From)
	ac.MessageConfig.Subject = mergeStrings(ac.MessageConfig.Subject, fc.MessageConfig.Subject, cc.MessageConfig.Subject)
	ac.MessageConfig.MessageBody = mergeStrings(ac.MessageConfig.MessageBody, fc.MessageConfig.MessageBody, cc.MessageConfig.MessageBody)

	// Merge RecipientsList
	if err := mergeRecipientsLists(ctx, ac, fc, cc); err != nil {
		return err
	}

	// Merge Recipients
	if err := mergeRecipients(ctx, ac, fc, cc); err != nil {
		return err
	}

	return nil
}
*/

/* func mergeStrings(ac string, s ...string) string {
	for _, v := range s {
		if v != "" {
			return v
		}
	}
	return ac
} */

/* func mergeInts(ac int, i ...int) int {
	for _, v := range i {
		if v != 0 {
			return v
		}
	}
	return ac
} */

/* func mergeBool(ctx context.Context, name string, fc, cc bool) *bool {
	if IsFlagPassed(name) {
		logging.L(ctx).Debug("mergeConfigs: flag provided, setting flag value", logging.StringAttr("flag", name), logging.BoolAttr("value", fc))
		return &fc
	}
	return &cc
} */

/* func mergeAuth(ctx context.Context, ac, fc, cc *types.Config) error {
	if fc.AppConfig.Auth.Username == "" && cc.AppConfig.Auth.Username == "" {
		logging.L(ctx).Debug("mergeConfigs: smtp auth credentials is not set")
		return errors.New("mergeConfigs: smtp auth credentials is not set")
	} else if fc.AppConfig.Auth.Username != "" {
		ac.AppConfig.Auth.Username = fc.AppConfig.Auth.Username
		ac.AppConfig.Auth.Password = fc.AppConfig.Auth.Password
		logging.L(ctx).Debug("mergeConfigs: smtp auth credentials is setted from flags values", logging.StringAttr("username", ac.AppConfig.Auth.Username), logging.StringAttr("password", ac.AppConfig.Auth.Password))
	} else if cc.AppConfig.Auth.Username != "" {
		ac.AppConfig.Auth.Username = cc.AppConfig.Auth.Username
		ac.AppConfig.Auth.Password = cc.AppConfig.Auth.Password
		logging.L(ctx).Debug("mergeConfigs: default smtp auth credentials is setted from config file values", logging.StringAttr("username", ac.AppConfig.Auth.Username), logging.StringAttr("password", ac.AppConfig.Auth.Password))
	}
	return nil
} */

/* func mergeRecipientsLists(ctx context.Context, ac, fc, cc *types.Config) error {
	if fc.LenRecipientsList() > 0 {
		ac.MessageConfig.RecipientsList = fc.MessageConfig.RecipientsList
		logging.L(ctx).Debug("mergeConfigs: recipients list added from flags config", logging.StringAttr("recipients-list", fc.MessageConfig.RecipientsList))
		return nil
	} else if cc.LenRecipientsList() > 0 {
		ac.MessageConfig.RecipientsList = cc.MessageConfig.RecipientsList
		logging.L(ctx).Debug("mergeConfigs: recipients list added from config file", logging.StringAttr("recipients-list", cc.MessageConfig.RecipientsList))
		return nil
	}
	logging.L(ctx).Error("mergeConfigs: Email recipients not specified")
	return errors.New("mergeConfigs: Email recipients not specified")
} */

/* func mergeRecipients(ctx context.Context, ac, fc, cc *types.Config) error {
	if len(fc.MessageConfig.Recipients) > 0 {
		ac.MessageConfig.Recipients = make([]string, 0, len(fc.MessageConfig.Recipients))
		for i := 0; i < len(fc.MessageConfig.Recipients); i++ {
			ac.AddRecipient(fc.GetRecipient(i))
			logging.L(ctx).Debug("mergeConfigs: recipient added from flags config", logging.StringAttr("recipient", fc.GetRecipient(i)))
		}
		return nil

	} else if len(cc.MessageConfig.Recipients) > 0 {
		ac.MessageConfig.Recipients = make([]string, 0, len(cc.MessageConfig.Recipients))
		for i := 0; i < len(cc.MessageConfig.Recipients); i++ {
			ac.AddRecipient(cc.GetRecipient(i))
			logging.L(ctx).Debug("mergeConfigs: recipient added from config file", logging.StringAttr("recipient", cc.GetRecipient(i)))
		}
		return nil

	} else {
		logging.L(ctx).Error("mergeConfigs: Email recipients does not specified")
		//os.Exit(1)
		return errors.New("mergeConfigs: Email recipients does not specified")
	}
} */
