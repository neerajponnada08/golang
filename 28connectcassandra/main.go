package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	username = "postgres"
	password = "123456"
	dbname   = "first_db"
)

func main() {
	con := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	db, err := sql.Open("postgres", con)
	CheckErr(err)

	defer db.Close()


	// createstmt := `CREATE TABLE COMPANY(
	// 	NAME TEXT PRIMARY KEY     NOT NULL,
	// 	ID   INT    NOT NULL
	//  );`

	// _, e := db.Exec(createstmt)
	// CheckErr(e)

	// insertStmt := `insert into company("name", "id") values('Rohit', 21)`

	// _, e := db.Exec(insertStmt)
	// CheckErr(e)

	// insertDynStmt := `insert into company("name", "id") values($1, $2)`
	// _, e = db.Exec(insertDynStmt, "kridsh", 03)
	// CheckErr(e)

	
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
