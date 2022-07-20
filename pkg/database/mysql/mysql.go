package mysql

import (
	"database/sql"
	"fmt"
	"library/internal/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Open(conf *config.Config) (*sql.DB, error) {

	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.DBname))
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	log.Print("Mysql connect sucsses!")
	return conn, nil
}
