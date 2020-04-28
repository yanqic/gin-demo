package log

import (
	"fmt"
	"time"

	"gin-demo/config"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", config.LogConfig.Path)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		config.LogConfig.FileName,
		time.Now().Format(config.LogConfig.TimeFormat),
		config.LogConfig.FileExt,
	)
}
