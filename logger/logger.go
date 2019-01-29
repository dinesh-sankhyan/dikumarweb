package logger

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

//Level log level
type Level uint8
type logfunc func()

const (
	//PanicLevel panic level
	PanicLevel Level = iota
	//FatalLevel fatal level
	FatalLevel
	//ErrorLevel error level
	ErrorLevel
	//WarnLevel warn level
	WarnLevel
	//InfoLevel info level
	InfoLevel
	//DebugLevel debug level
	DebugLevel
)

//LogEntry log entry
var LogEntry *logrus.Entry

//InitLogger initialize looger
func InitLogger(logLevel string, appName string, buildVersion string,
	 logFile string, appender string) error {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.Fatal(err.Error())
	}
	logger := logrus.New()
	formatter := &logrus.JSONFormatter{
		DisableTimestamp: true,
	}
	logger.Formatter = formatter
	logger.Level = level
	// use logrus for standard log output:
	if appender == "file" {
		f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			log.SetOutput(os.Stderr)
		} else {
			log.SetOutput(f)
		}
	}
	if appender == "console" {
		log.SetOutput(os.Stderr)
	}
	LogEntry = logger.WithField("app", appName)
	LogEntry = logger.WithField("version", buildVersion)
	return err
}


//Print print the info
func Print(args ...interface{}) {
	LogEntry.Print(args...)
}

//Info log level
func Info(args ...interface{}) {
	LogEntry.Info(args...)
}

//Warn log level
func Warn(args ...interface{}) {
	LogEntry.Warn(args...)
}

//Warning log level
func Warning(args ...interface{}) {
	LogEntry.Warning(args...)
}

//Error log level
func Error(args ...interface{}) {
	LogEntry.Error(args...)
}

//Fatal Fatal level
func Fatal(args ...interface{}) {
	LogEntry.Fatal(args...)
}

//Fatalf Fatalf level
func Fatalf(format string, args ...interface{}) {
	LogEntry.Fatalf(format, args...)
}

//Debugf log level
func Debugf(format string, args ...interface{}) {
	LogEntry.Debugf(format, args...)
}

//Infof log level
func Infof(format string, args ...interface{}) {
	LogEntry.Infof(format, args...)
}

//Printf log level
func Printf(format string, args ...interface{}) {
	LogEntry.Printf(format, args...)
}

//Warnf log level
func Warnf(format string, args ...interface{}) {
	LogEntry.Warnf(format, args...)
}

//Warningf log level
func Warningf(format string, args ...interface{}) {
	LogEntry.Warningf(format, args...)
}

//Errorf log level
func Errorf(format string, args ...interface{}) {
	LogEntry.Errorf(format, args...)
}

//Infoln log level
func Infoln(args ...interface{}) {
	LogEntry.Infoln(args...)
}

//Println log level
func Println(args ...interface{}) {
	LogEntry.Println(args...)
}

//WithError Add an error as single field (using the key defined in ErrorKey) to the Entry.
func WithError(err error) *logrus.Entry {
	return LogEntry.WithError(err)
}
