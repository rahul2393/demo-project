package samples

import (
	"database/sql"
	"log"
)

func UpdateSingersUsingDML(db *sql.DB) {
	log.Println("running updateSingersUsingDML..")
	stmt, err := db.Prepare("UPDATE singers SET name=? WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	res, err := stmt.Exec("Updated newly inserted Singer", 2)
	if err != nil {
		log.Fatal(err)
	}
	affected, err := res.RowsAffected()
	log.Printf("Updated %v row(s)\n", affected)
	log.Printf("successfully executed UpdateSingersUsingDML...\n\n")
}
