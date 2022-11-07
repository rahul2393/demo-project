package main

import (
	"database/sql"
	"log"

	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/googleapis/go-sql-spanner"

	"github.com/rahul2393/demo-project/samples"
)

// MYSQL schema
// CREATE TABLE Singers (
//	Id int NOT NULL,
//	Name varchar(255),
//	PRIMARY KEY (Id)
// );
//
// Spanner schema
// CREATE TABLE Singers (
//	Id   INT64 NOT NULL,
//	Name  STRING(1024)
// ) PRIMARY KEY (Id);

func main() {
	db, err := sql.Open("spanner",
		"projects/span-cloud-testing/instances/spanner-testing/databases/irahul-test")
	//db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	samples.ReadSingerByID(db, 1)
	samples.InsertSingerUsingDML(db)
	samples.ReadSingerByID(db, 2)
	samples.UpdateSingersUsingDML(db)
	samples.ReadSingers(db)
	samples.DeleteSingersByID(db, 2)
}
