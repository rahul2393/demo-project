package samples

import (
	"database/sql"
	"fmt"
	"log"
)

func ReadSingerByID(db *sql.DB, id int) {
	log.Println("running readSingerByID..")
	var str string
	err := db.QueryRow(fmt.Sprintf("SELECT name FROM singers WHERE id =%v", id)).Scan(&str)
	if err != nil && err != sql.ErrNoRows {
		log.Panic(err)
	}
	log.Println(str)
	log.Printf("successfully executed readSingerByID...\n\n")
}
