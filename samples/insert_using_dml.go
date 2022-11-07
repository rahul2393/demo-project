package samples

import (
	"database/sql"
	"log"
)

func InsertSingerUsingDML(db *sql.DB) {
	log.Println("running insertSingerUsingDML..")
	_, err := db.Exec("INSERT INTO singers (id, name) VALUES(?, ?)",
		2, "Newly inserted Singer")
	if err != nil {
		log.Panic(err)
	}
	log.Printf("successfully executed insertSingerUsingDML...\n\n")
}
