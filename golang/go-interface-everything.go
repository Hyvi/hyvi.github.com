/**
CREATE TABLE IF NOT EXISTS `test`.`task` (

	`id` INT(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`name` VARCHAR(45) NOT NULL COMMENT '任务名'
)
ENGINE = InnoDB
AUTO_INCREMENT = 1
DEFAULT CHARACTER SET = utf-8
*/

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func query() {
	db, err := sql.Open("mysql", "dbuser:dbpass@tcp(127.0.0.1:3306)/test?charset=utf-8")
	checkErr(err)
	defer db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
