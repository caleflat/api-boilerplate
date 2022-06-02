package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// Init initializes the logger
func Init() {
	log.SetFormatter(&log.JSONFormatter{
		PrettyPrint: false,
	})
	// log.SetReportCaller(true) // decreases logger performance by 20-40%
	log.SetOutput(os.Stdout)

	// Sets the level of logging, debug logs everything
	log.SetLevel(log.DebugLevel)

	log.Debug("Logger initialized")
}
