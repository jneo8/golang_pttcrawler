package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	appName = "ptt-crawler"
)

var (
	cfgFile  string
	logLevel string
)

// rootCmd is the root command
var RootCmd = &cobra.Command{
	Use:               appName,
	Short:             appName,
	Long:              appName,
	PersistentPreRunE: InitAll,
	RunE: func(cmd *cobra.Command, arg []string) error {
		logrus.Info("RunE")
		return nil
	},
}

func init() {
	// config file.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is config/%s.yaml)", "default"))

	// Log level.
	RootCmd.PersistentFlags().StringVar(&logLevel, "LOG_LEVEL", "", fmt.Sprintf("Log Level (default is %s)", "DEBUG"))
}
