package gocommon

import (
	"os"
	"runtime/debug"
	"time"

	logStack "github.com/Gurpartap/logrus-stack"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

const (
	directoryName    = "logs"
	errorFileName    = directoryName + "/" + "Error.log"
	warnFileName     = directoryName + "/" + "Warn.log"
	infoFileName     = directoryName + "/" + "Info.log"
	debugFileName    = directoryName + "/" + "Debug.log"
	rotateFormat     = ".%Y%m%d"
	logStackConstant = "Stack trace: "
	rotationTime     = 84600
	maxAge           = 2538000
)

var fileLogger *logrus.Logger

func init() {
	os.Mkdir(directoryName, 0700)
	writerError, _ := rotatelogs.New(
		errorFileName+rotateFormat,
		//rotatelogs.WithLinkName(errorFileName),
		rotatelogs.WithMaxAge(time.Duration(maxAge)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(rotationTime)*time.Second),
	)
	writerWarn, _ := rotatelogs.New(
		warnFileName+rotateFormat,
		//rotatelogs.WithLinkName(warnFileName),
		rotatelogs.WithMaxAge(time.Duration(maxAge)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(rotationTime)*time.Second),
	)
	writerInfo, _ := rotatelogs.New(
		infoFileName+rotateFormat,
		//rotatelogs.WithLinkName(infoFileName),
		rotatelogs.WithMaxAge(time.Duration(maxAge)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(rotationTime)*time.Second),
	)
	writerDebug, _ := rotatelogs.New(
		debugFileName+rotateFormat,
		//rotatelogs.WithLinkName(debugFileName),
		rotatelogs.WithMaxAge(time.Duration(maxAge)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(rotationTime)*time.Second),
	)

	formatter := &logrus.JSONFormatter{}
	formatter.TimestampFormat = "2006-01-02 15:04:05"

	pathMap := lfshook.WriterMap{
		logrus.WarnLevel:  writerWarn,
		logrus.ErrorLevel: writerError,
		logrus.DebugLevel: writerDebug,
		logrus.InfoLevel:  writerInfo,
	}
	fileLogger = logrus.New()
	fileLogger.SetFormatter(formatter)

	fileLogger.Hooks.Add(lfshook.NewHook(
		pathMap,
		formatter,
	))
	fileLogger.AddHook(logStack.StandardHook())
}

// LogErrorObject This method logs error
func LogErrorObject(err error, logStack bool) {
	if logStack {
		fileLogger.WithField(logStackConstant, string(debug.Stack())).Error(err.Error())
	} else {
		fileLogger.Error(err.Error())
	}
}

// LogError This method logs error
func LogError(message string, logStack bool) {
	if logStack {
		fileLogger.WithField(logStackConstant, string(debug.Stack())).Error(message)
	} else {
		fileLogger.Error(message)
	}
}

// LogInfo This method logs error
func LogInfo(message string, logStack bool) {
	if logStack {
		fileLogger.WithField(logStackConstant, string(debug.Stack())).Info(message)
	} else {
		fileLogger.Info(message)
	}
}

// LogWarn This method logs error
func LogWarn(message string, logStack bool) {
	if logStack {
		fileLogger.WithField(logStackConstant, string(debug.Stack())).Warn(message)
	} else {
		fileLogger.Warn(message)
	}
}

// LogDebug This method logs error
func LogDebug(message string, logStack bool) {
	if logStack {
		fileLogger.WithField(logStackConstant, string(debug.Stack())).Debug(message)
	} else {
		fileLogger.Debug(message)
	}
}

// LogErrorWithFields This method logs error
func LogErrorWithFields(fields map[string]interface{}, message string, logStack bool) {
	if logStack {
		fields[logStackConstant] = string(debug.Stack())
	}
	fileLogger.WithFields(fields).Error(message)
}

// LogInfoWithFields This method logs error
func LogInfoWithFields(fields map[string]interface{}, message string, logStack bool) {
	if logStack {
		fields[logStackConstant] = string(debug.Stack())
	}
	fileLogger.WithFields(fields).Info(message)
}

// LogWarnWithFields This method logs error
func LogWarnWithFields(fields map[string]interface{}, message string, logStack bool) {
	if logStack {
		fields[logStackConstant] = string(debug.Stack())
	}
	fileLogger.WithFields(fields).Warn(message)
}

// LogDebugWithFields This method logs error
func LogDebugWithFields(fields map[string]interface{}, message string, logStack bool) {
	if logStack {
		fields[logStackConstant] = string(debug.Stack())
	}
	fileLogger.WithFields(fields).Debug(message)
}
