package logx

import (
	log "github.com/mgutz/logxi/v1"
)

var logger log.Logger

func Logger() log.Logger {
	return logger
}

func InitLogger() {
	envConf := log.Configuration{
		Format: log.FormatJSON,
	}
	log.ProcessEnv(&envConf)

	logger = log.New("rr-backend")
	logger.SetLevel(log.LevelAll)
}
