package logging

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/covalenthq/lumberjack/v3"
	"jy.org/verse/src/config"
)

var cfg = config.Config.Log

type logger struct {
    ERROR *log.Logger
    WARN *log.Logger
    INFO *log.Logger
}

func isValidLogfile(path string) bool {
    // check if parent exists
    parent := filepath.Dir(path)
    if _, err := os.Stat(parent); err != nil {
        return false
    }
    return true
}

var Logger = &logger{
    ERROR: log.New(os.Stderr, "ERROR:", log.LstdFlags|log.Lshortfile),
    WARN: log.New(os.Stdout, "WARN:", log.LstdFlags|log.Lshortfile),
    INFO: log.New(os.Stdout, "INFO:", log.LstdFlags|log.Lshortfile),
}

func Init() {
    // Setting up lumberjack logger for log rotation
	logFile := &lumberjack.Logger{
		Filename:   cfg.Path,
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
		Compress:   true, // disabled by default
	}

    flags := log.LstdFlags | log.Lshortfile
    Logger.ERROR = log.New(io.MultiWriter(os.Stderr, logFile), "ERROR:", flags)
    Logger.WARN = log.New(io.MultiWriter(os.Stdout, logFile), "WARN:", flags)
    Logger.INFO = log.New(io.MultiWriter(os.Stdout, logFile), "INFO:", flags)
}

