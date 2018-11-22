package main

import (
	"github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"
	crawler "golang_pttcrawler/crawler"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
	log.AddHook(filename.NewHook())

}

func main() {
	crawler.Init()
}
