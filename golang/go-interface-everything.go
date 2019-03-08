/**
CREATE TABLE  task (
	id INT(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	name VARCHAR(45) NOT NULL COMMENT '任务名',
	PRIMARY KEY (id)
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
	db, err := sql.Open("mysql", "dbuser:dbpass@tcp(127.0.0.1:3306)/dbname")
	checkErr(err)
	defer db.Close()
	rows, err := db.Query("SELECT * FROM task")
	checkErr(err)
	defer rows.Close()
	var (
		id   int
		name string
	)

	for rows.Next() {
		err = rows.Scan(&id, &name)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(name)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func query_v2() {
	db, err := sql.Open("mysql", "dbuser:dbpass@tcp(127.0.0.1:3306)/dbname")
	checkErr(err)
	defer db.Close()
	rows, err := db.Query("SELECT * FROM task")
	checkErr(err)
	defer rows.Close()
	columns, _ := rows.Columns()
	args := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		args[i] = &values[i]
	}

	record := make(map[string]string)
	for rows.Next() {
		err = rows.Scan(args...)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]uint8))
			}
		}
		fmt.Println(record)
	}
}

func main() {
	query_v2()
}
