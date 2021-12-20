package initDB

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type conn struct {
	Host     string
	Dbname   string
	Username string
	Password string
}
type Config struct {
	Mysql map[string]conn
}

var Db *sql.DB

func init() {
	//传入配置路径
	yamlFile, err := ioutil.ReadFile("./conf/db.yaml")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	conf := new(Config)
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Panicln("err:", err.Error())
	}

	//根据字段需要选择配置节
	db, ok := conf.Mysql["default"]
	if !ok {
		log.Panicln("err:", "db default no set")
	}

	Db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", db.Username, db.Password, db.Host, db.Dbname))
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
}
