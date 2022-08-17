package main

import (
	"fmt"
	"strconv"
	"test/initialize"
)

var (
	sql = "INSERT INTO `audit`.`region_management`( `created_at`, `updated_at`, `deleted_at`, `region_name`, `region_description`, `parent_id`) VALUES " +
		"( '2022-07-27 10:21:29.000', '2022-07-27 10:21:29.000', NULL, '安全大区%s', '安全大区%s', 0);"
)

func main() {
	db := initialize.GormMysql()
	sqlexec := "INSERT INTO `audit`.`region_management`( `created_at`, `updated_at`, `deleted_at`, `region_name`, `region_description`, `parent_id`) VALUES"
	for i := 1197; i < 10000; i++ {
		sqlAlone := "( '2022-07-27 10:21:29.000', '2022-07-27 10:21:29.000', NULL, '安全大区%s', '安全大区%s', 0),"
		sprintf := fmt.Sprintf(sqlAlone, strconv.Itoa(i), strconv.Itoa(i))
		sqlexec += sprintf
	}
	sqlexec = sqlexec[:len(sqlexec)-1] + ";"
	db.Exec(sqlexec)
}
