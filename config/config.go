package config

import (
  "github.com/spf13/viper"
  "log"
)

var cfgDatabase *viper.Viper
var cfgApplication *viper.Viper

func init() {
  viper.SetConfigName("setting")
  viper.AddConfigPath("./config")
  err := viper.ReadInConfig()
  if err != nil {
    log.Println(err)
  }
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
}

func SetApplicationIsInit() {
  SetConfig("./config","settings.application.isInit", false)
}

func SetConfig(configPath string,key string,value interface{}){
  viper.AddConfigPath(configPath)
  viper.Set(key, value)
  err := viper.WriteConfig()
  if err != nil {
    panic("config write error")
  }
}
