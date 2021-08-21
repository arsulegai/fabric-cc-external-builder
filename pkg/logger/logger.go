package logger

import (
	"external-builder/pkg/env"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

const (
	DebugLevel                       = "DEBUG"
	InfoLevel                        = "INFO"
	WarnLevel                        = "WARN"
	ErrorLevel                       = "ERROR"
	CorePeerChaincodeBuilderLogLevel = "CORE_PEER_CC_BUILDER_LOG_LEVEL"
)

func InitLogger(instance string) {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{
		DataKey: instance,
	})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stderr)
	logLevel := readLogLevel()
	log.SetLevel(logLevel)
}

func readLogLevel() log.Level {
	level := env.GetEnvOrDefault(CorePeerChaincodeBuilderLogLevel, WarnLevel)
	var logLevel log.Level
	switch strings.ToUpper(level) {
	case DebugLevel:
		logLevel = log.DebugLevel
	case InfoLevel:
		logLevel = log.InfoLevel
	case WarnLevel:
		logLevel = log.WarnLevel
	case ErrorLevel:
		logLevel = log.ErrorLevel
	default:
		logLevel = log.WarnLevel
	}
	return logLevel
}
