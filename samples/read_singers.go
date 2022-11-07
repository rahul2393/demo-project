package samples

import (
	"database/sql"
	"log"
)

func ReadSingers(db *sql.DB) {
	log.Println("running readSingers..")
	rows, err := db.Query("SELECT * FROM singers")
	if err != nil && err != sql.ErrNoRows {
		log.Panic(err)
	}
	defer rows.Close()
	var (
		id, str string
	)
	for rows.Next() {
		err = rows.Scan(&id, &str)
		if err != nil {
			log.Panic(err)
		}
		log.Printf("found Id: %v, Name: %v\n", id, str)
	}
	if err = rows.Err(); err != nil {
		rows.Close()
		log.Panic(err)
	}
	log.Printf("successfully executed readSingers...\n\n")
}
