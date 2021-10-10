package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "daniel:12345678@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Update
	stmt, _ := db.Prepare("update usuarios set nome = ? where id= ?")
	stmt.Exec("novo usuario 1", 1)
	stmt.Exec("novo usuario 2", 2)

	//delete
	stmt.Exec("delete from usuarios where id = ?")
	stmt.Exec(3)
}
