package logger

import (
	"os"
	"path"
	"time"

	"git.garena.com/shopee/seller-server/seller-marketing/marketing-common/mock/constant"
	"git.garena.com/shopee/seller-server/seller-marketing/marketing-common/mock/utils"

	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// Log ...
var Log *logrus.Logger

// MockLogger ...
type MockLogger struct {
	Log *logrus.Logger
}

func init() {
	fileName := path.Join(constant.LogFilePath, constant.LogFileName)
	if !utils.FileExists(fileName) {
		utils.FileCreate(fileName)
	}
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		println(err)
	}

	Log = logrus.New()
	Log.Out = src
	SetLogLevel(constant.LogLevel)
	writeMap := lfshook.WriterMap{
		logrus.DebugLevel: SetRotateLogs(fileName + ".debug"),
		logrus.InfoLevel:  SetRotateLogs(fileName + ".info"),
		logrus.WarnLevel:  SetRotateLogs(fileName + ".warn"),
		logrus.ErrorLevel: SetRotateLogs(fileName + ".error"),
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	Log.AddHook(lfHook)
	Log.SetReportCaller(true)
}

// SetRotateLogs ...
func SetRotateLogs(fileName string) *rotateLogs.RotateLogs {
	logWriter, err := rotateLogs.New(
		fileName+".%Y%m%d.log",
		rotateLogs.WithLinkName(fileName),
		rotateLogs.WithMaxAge(7*24*time.Hour),
		rotateLogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		println(err)
		return nil
	}
	return logWriter
}

// SetLogLevel ...
func SetLogLevel(logLevel string) {
	if Log == nil {
		return
	}
	if logLevel == "debug" {
		Log.SetLevel(logrus.DebugLevel)
	} else if logLevel == "info" {
		Log.SetLevel(logrus.InfoLevel)
	} else if logLevel == "warn" {
		Log.SetLevel(logrus.WarnLevel)
	} else {
		Log.SetLevel(logrus.ErrorLevel)
	}
}
