package routing

import (
	"io"
	"os"
	"log"
	"net/http"
	"path/filepath"
	"github.com/go-chi/chi/middleware"
	"geektest/internal/configuration"
	"geektest/config"
)

var DefaultLogger func(next http.Handler) http.Handler
var appConfig *configuration.Config

func init() {
	appConfig = config.GetConfig()
	logHandler := getHandle(appConfig.RequestLogFilePath)
	DefaultLogger = middleware.RequestLogger(
		&middleware.DefaultLogFormatter{
			Logger: log.New(logHandler, "", log.LstdFlags), NoColor: false})
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