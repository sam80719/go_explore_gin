package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	// ref: https://github.com/go-sql-driver/mysql/
	//Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/blog") // use 127.0.0.1:3306 when golang is used external docker
	//Db, err = sql.Open("mysql", "root:root@tcp(mysql:3306)/blog") // use docker should use docker name
	Db, err = sql.Open("mysql", "root:root@tcp(docker.for.mac.localhost:3306)/blog") // or use mac on docker should use "docker.for.mac.localhost"
	if err != nil {
		log.Panicln("err", err.Error())
	}
	Db.SetMaxIdleConns(20)
	Db.SetMaxIdleConns(20)
}
