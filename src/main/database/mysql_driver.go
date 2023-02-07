package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func SqlConnect(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:1234567890@/mdm_data_20220921")
	checkErr(err)
	rows, err := db.Query("select user_code, user_name from mdm_user limit 1")
	checkErr(err)
	for rows.Next() {
		var user_code string
		var user_name string
		err := rows.Scan(&user_code, &user_name)
		checkErr(err)
		fmt.Println(user_code + "--" + user_name)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
