package log

import "log"

var logger *log.Logger
var enabled bool

func Enable(enable bool) {
	enabled = enable
}

func Print(message string) {
	if !enabled {
		return
	}

	if logger == nil {
		logger = log.Default()
	}

	logger.Print(message)
}
