package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/theSmallwhiteMe/make_table_struct/get_database_schema"
)

func main() {
	dbConfigString := "root:root@tcp(10.25.10.126:3306)/audit?charset=utf8"
	db, err := sql.Open("mysql", dbConfigString)
	if err == nil {
		dbOption := &get_database_schema.DBOption{
			"audit",
			db,
		}
		get_database_schema.Run(dbOption, get_database_schema.NewMysql())

	}
}
