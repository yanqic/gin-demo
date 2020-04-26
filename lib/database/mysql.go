package database

import (
  "bytes"
  "fmt"
  config2 "gin-demo/config"
  _ "github.com/go-sql-driver/mysql" //加载mysql
  "github.com/jinzhu/gorm"
  "log"
  "strconv"
)

var Eloquent *gorm.DB

func init() {

  dbType := config2.DatabaseConfig.Dbtype
  host := config2.DatabaseConfig.Host
  port := config2.DatabaseConfig.Port
  database := config2.DatabaseConfig.Database
  username := config2.DatabaseConfig.Username
  password := config2.DatabaseConfig.Password

  if dbType != "mysql" && dbType != "sqlite3" {
    fmt.Println("db type unknow")
  }
  var err error

  var conn bytes.Buffer
  conn.WriteString(username)
  conn.WriteString(":")
  conn.WriteString(password)
  conn.WriteString("@tcp(")
  conn.WriteString(host)
  conn.WriteString(":")
  conn.WriteString(strconv.Itoa(port))
  conn.WriteString(")")
  conn.WriteString("/")
  conn.WriteString(database)
  conn.WriteString("?charset=utf8&parseTime=True&loc=Local&timeout=1000ms")

  log.Println(conn.String())

  var db Database
  if dbType != "mysql" {
    panic("db type unknow")
  }
  db = new(Mysql)
  Eloquent, err = db.Open(dbType, conn.String())

  Eloquent.LogMode(true)
  if err != nil {
    log.Fatalf("%s connect error %v\n", dbType, err)
  } else {
    log.Printf("%s connect success!\n", dbType)
  }

  if Eloquent.Error != nil {
    log.Fatalf("database error %v\n", Eloquent.Error)
  }

}

type Database interface {
  Open(dbType string, conn string) (db *gorm.DB, err error)
}

type Mysql struct {
}

func (*Mysql) Open(dbType string, conn string) (db *gorm.DB, err error) {
  eloquent, err := gorm.Open(dbType, conn)
  return eloquent, err
}
