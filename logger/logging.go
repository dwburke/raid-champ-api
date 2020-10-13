package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
	"github.com/spf13/viper"
)

var Logfile *os.File

func init() {
	viper.SetDefault("logger.stdout.level", "info")
	viper.SetDefault("logger.file.maxsize", 5)
	viper.SetDefault("logger.file.maxbackups", 7)
	viper.SetDefault("logger.file.maxage", 7)
	viper.SetDefault("logger.file.enabled", false)
}

func interpretLogLevel(level string) (log_level log.Level) {

	switch level {
	case "trace":
		log_level = log.TraceLevel
	case "debug":
		log_level = log.DebugLevel
	case "info":
		log_level = log.InfoLevel
	case "warn":
		log_level = log.WarnLevel
	case "error":
		log_level = log.ErrorLevel
	case "fatal":
		log_level = log.FatalLevel
	case "panic":
		log_level = log.PanicLevel
	default:
		log_level = log.DebugLevel
	}

	return
}

func InitLogging() {
	log.SetOutput(os.Stdout)
	log.SetLevel(interpretLogLevel(viper.GetString("logger.stdout.level")))

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	setupFileLogging()
}

func setupFileLogging() {
	if !viper.GetBool("logger.file.enabled") {
		return
	}

	filename := viper.GetString("logger.file.name")

	log.Infof("logrus: opening logfile: %s", filename)

	if filename == "" {
		panic("logger.file.enabled == true && logger.file.name == \"\"\n")
	}

	var err error
	Logfile, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   filename,
		MaxSize:    viper.GetInt("logger.file.maxsize"),
		MaxBackups: viper.GetInt("logger.file.maxbackups"),
		MaxAge:     viper.GetInt("logger.file.maxage"),
		Level:      interpretLogLevel(viper.GetString("logger.file.level")),
		Formatter: &log.TextFormatter{
			FullTimestamp: true,
		},
	})

	log.AddHook(rotateFileHook)

	log.Infof("logrus: logger to file: %s", filename)
}

func Cleanup() {
	log.Info("logger: Cleanup()")
	if Logfile != nil {
		Logfile.Close()
	}
}
