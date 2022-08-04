package log

import (
	"github.com/sirupsen/logrus"
	"log"
	logrusadapter "logur.dev/adapter/logrus"
	"logur.dev/logur"
	"os"
)

//hold details necessary for logging
type Config struct {
	//Format specify output log format
	//Accepted value: json, logfmt
	Format string

	//Level: minimum log level that should appear in our log
	Level string

	//NoColor make sure tht no log output get color
	NoColor bool
}

// NewLogger creates a custom logger, not global library
func NewLogger(config Config) logur.LoggerFacade {
	logger := logrus.New()

	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:               false,
		DisableColors:             config.NoColor,
		ForceQuote:                false,
		EnvironmentOverrideColors: true,
		DisableTimestamp:          false,
		FullTimestamp:             false,
		TimestampFormat:           "",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		PadLevelText:              false,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier:          nil,
	})
	switch config.Format {
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{})

	case "logfmt":
		//default already
	}
	if level, err := logrus.ParseLevel(config.Level); err != nil {
		logger.SetLevel(level)
	}
	return logrusadapter.New(logger)
}

// SetStandardLogger sets the global logger's output to a custom logger instance.
func SetStandaloneLogger(logger logur.Logger) {
	log.SetOutput(logur.NewLevelWriter(logger, logur.Info))
}
