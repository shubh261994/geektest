package logs

import (
	"io"
	"os"
	"path/filepath"
	"github.com/op/go-logging"
	"geektest/internal/configuration"
	"geektest/config"
)

var logger *logging.Logger
var appConfig *configuration.Config
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func init()  {
	appConfig = config.GetConfig()
	logger = logging.MustGetLogger(appConfig.Namespace)
	handle := getHandle(appConfig.LogFilePath)
	configureLogger(handle, appConfig.LogLevel)
}

func configureLogger(handle io.Writer, level logging.Level) {
	backend := logging.NewLogBackend(handle, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	backendLeveled := logging.AddModuleLevel(backendFormatter)
	backendLeveled.SetLevel(level, "")
	logging.SetBackend(backendLeveled)
}

func getHandle(logFile string) io.Writer {
	if appConfig.Deployment == configuration.DEBUG {
		return os.Stdout
	}

	var handle io.Writer
	var err error
	if logFile != "" {
		logFilePath := filepath.Join(configuration.BASEPATH, logFile)
		handle, err = openOrCreateFile(logFilePath)
		if err!=nil {
			panic(err)
		}
	} else {
		handle = os.Stdout
	}

	return handle
}

func openOrCreateFile(fileName string) (*os.File, error) {
	var file *os.File
	var _, err = os.Stat(fileName)
	if os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(fileName), os.ModePerm)
	}

	file, err = os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	return file, err
}