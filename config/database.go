package config

import "github.com/spf13/viper"

type Database struct {
	Dbtype      string
	Host        string
	Port        int
	DbName      string
	UserName    string
	Password    string
	TablePrefix string
}

func InitDatabase(cfg *viper.Viper) *Database {
	return &Database{
		Port:        cfg.GetInt("port"),
		Dbtype:      cfg.GetString("dbType"),
		Host:        cfg.GetString("host"),
		DbName:      cfg.GetString("dbname"),
		UserName:    cfg.GetString("username"),
		Password:    cfg.GetString("password"),
		TablePrefix: cfg.GetString("tableprefix"),
	}
}

var DatabaseConfig = new(Database)
