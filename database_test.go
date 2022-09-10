package golang_database

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/absensi")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()
	fmt.Println("Successfully connect to database")
}
