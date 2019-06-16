package main

import (
	"github.com/sirupsen/logrus"
	"golang_pttcrawler/pkg/cmd"
	"os"
)

func main() {

	if err := cmd.RootCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
