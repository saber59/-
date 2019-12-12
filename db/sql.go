package db

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	USERNAME  string
	PASSWORD  string
	IPADDRESS string
	DATABASE  string
}

var db *sql.DB

type msl interface {
	LinkMysql() (bool, *sql.DB)
}

func (p *Mysql) LinkMysql() (error, *sql.DB) {
	var conf Mysql
	data, err := ioutil.ReadFile("db/config.json")
	if err != nil {
		log.Println("Read config file error:", err)
	}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		log.Println("Unmarshal config data error:", err)
	}
	LinkString := "root:" + conf.PASSWORD + "@tcp(" + conf.IPADDRESS + ")/" + conf.DATABASE + "?charset=utf8"
	conn, err := sql.Open("mysql", LinkString)
	if err != nil {
		log.Print("MySQL连接失败")
		return err, nil
		//return conn
	}
	err = conn.Ping()
	if err == nil {
		return nil, conn
	}
	log.Print("数据库连接出错")
	return err, nil

}
