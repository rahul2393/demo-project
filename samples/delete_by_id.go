package samples

import (
	"database/sql"
	"fmt"
	"log"
)

func DeleteSingersByID(db *sql.DB, Id int) {
	log.Println("running deleteSingersByID..")
	result, err := db.Exec("DELETE FROM singers WHERE id=?", 2)
	if err != nil {
		log.Fatal(err)
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v row(s)\n", affectedRows)
	log.Printf("successfully executed deleteSingersByID...\n\n")
}
