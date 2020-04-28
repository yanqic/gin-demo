package config

import "github.com/spf13/viper"

type logConf struct {
	Path       string
	FileName   string
	FileExt    string
	TimeFormat string
}

func InitLog(cfg *viper.Viper) *logConf {
	return &logConf{
		Path:       cfg.GetString("path"),
		FileName:   cfg.GetString("filename"),
		FileExt:    cfg.GetString("fileext"),
		TimeFormat: cfg.GetString("timeformat"),
	}
}

var LogConfig = new(logConf)
