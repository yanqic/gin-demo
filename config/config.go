package config

import (
	"github.com/spf13/viper"
	"log"
)

var cfgLog *viper.Viper
var cfgDatabase *viper.Viper
var cfgApplication *viper.Viper
var cfgJwt *viper.Viper

func init() {
	viper.SetConfigName("setting")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	cfgLog = viper.Sub("setting.log")
	if cfgLog == nil {
		panic("config not found setting.log")
	}
	LogConfig = InitLog(cfgLog)

	cfgDatabase = viper.Sub("setting.database")
	if cfgDatabase == nil {
		panic("config not found setting.database")
	}
	DatabaseConfig = InitDatabase(cfgDatabase)

	cfgApplication = viper.Sub("setting.application")
	if cfgApplication == nil {
		panic("config not found setting.application")
	}
	ApplicationConfig = InitApplication(cfgApplication)

	cfgJwt = viper.Sub("setting.jwt")
	if cfgJwt == nil {
		panic("config not found setting.jwt")
	}
	JwtConfig = InitJwt(cfgJwt)
}
