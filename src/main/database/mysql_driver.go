package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func SqlConnect(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(ip:3306)/xxx?charset=utf8mb4")
	checkErr(err)
	rows, err := db.Query("select id, username from t_user")
	checkErr(err)
	for rows.Next() {
		var id int
		var username string
		err := rows.Scan(&id, &username)
		checkErr(err)
		fmt.Printf("%d -- %s", id, username)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
