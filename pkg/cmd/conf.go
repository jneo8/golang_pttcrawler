package cmd

import (
	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

// Init Viper & logrus.
func InitAll(cmd *cobra.Command, args []string) error {
	if err := InitViper(cmd, args); err != nil {
		return err
	}
	if err := InitLogrus(cmd, args); err != nil {
		return err
	}
	return nil
}

// Init logrus.
func InitLogrus(cmd *cobra.Command, args []string) error {
	// Set log formatter.
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000",
		FullTimestamp:   true,
	})

	// Set log level
	logLevel := viper.GetString("LOG_LEVEL")

	switch logLevel {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warning":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default: // default log level.
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.Infof("logLevel: %v", logrus.GetLevel().String())

	// Add filename hook.
	// Print filename & line number as source
	logrus.AddHook(filename.NewHook())

	logrus.Info("Init Logrus success!!")
	return nil
}

func InitViper(cmd *cobra.Command, args []string) error {
	// Get config file if set.
	var cfgFile string
	if cfgFlag := cmd.Flags().Lookup("config"); cfgFlag != nil {
		cfgFile = cfgFlag.Value.String()
	}

	// ENV
	viper.SetDefault("CONFIG_FILE", "default")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("CTA")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	logrus.Infof("CONFIG_FILE: %v", viper.GetString("CONFIG_FILE"))

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Yaml
		viper.SetConfigName(viper.GetString("CONFIG_FILE"))
		viper.SetConfigType("yaml")
		viper.AddConfigPath("../config")
		viper.AddConfigPath("./config")
	}

	if err := bindFlags(cmd); err != nil {
		return err
	}

	// Load yaml settings.
	if err := viper.ReadInConfig(); err == nil {
		logrus.Infof("Using config file: %v", viper.ConfigFileUsed())
	} else {
		logrus.Warning("Unfound config file.")
	}
	return nil
}

// Passing cmd.Commands to viper.
func bindFlags(cmd *cobra.Command) error {
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	for _, subcmd := range cmd.Commands() {
		if err := bindFlags(subcmd); err != nil {
			return err
		}
	}

	return nil
}
