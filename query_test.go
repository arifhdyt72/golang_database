package golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestExecContext(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	// query := `
	// 			INSERT INTO customer(name, email, balance, rating, birth_date, married)VALUES
	// 			('Arif', 'arif@test.com', 1000000, 9.0, '1994-08-12', false),
	// 			('Arif2', 'arif2@test.com', 500000, 7.5, '1984-08-12', true)
	// 		`
	query := `UPDATE customer set email = NULL, birth_date = NULL WHERE id = 5`

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Successfully Insert DB")
}

func TestQueryContext(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "SELECT * FROM customer"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err = rows.Scan(
			&id,
			&name,
			&email,
			&balance,
			&rating,
			&createdAt,
			&birthDate,
			&married,
		)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("id:", id, "name:", name, "email:", email.String, "balance:", balance, "rating:", rating, "birth date:", birthDate.Time, "married:", married, "created at:", createdAt)
	}
}

func TestSQLInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "salah"
	password := "admin"

	query := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Success login for User", username)
	} else {
		fmt.Println("Failed Login")
	}
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	email := "arif@test.com"
	comment := "Comment 1"

	query := "INSERT INTO comments(email, comment) VALUES(?,?)"
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		log.Panic(err.Error())
	}

	insertID, err := result.LastInsertId()
	if err != nil {
		log.Panic(err.Error())
	}

	fmt.Println("Success insert comments with ID", insertID)

}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "INSERT INTO comments(email, comment) VALUES(?,?)"
	statement, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Panic(err.Error())
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "arif" + strconv.Itoa(i) + "@test.com"
		comment := "Ini komentar ke - " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			log.Panic(err.Error())
		}

		lastId, err := result.LastInsertId()
		if err != nil {
			log.Panic(err.Error())
		}

		fmt.Println("Success insert comment with ID", lastId)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		log.Panic(err.Error())
	}

	query := "INSERT INTO comments(email,comment) VALUES(?,?)"

	// do transaction
	for i := 0; i < 10; i++ {
		email := "hidayat" + strconv.Itoa(i) + "@test.com"
		comment := "ini komentar ke - " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, query, email, comment)
		if err != nil {
			log.Panic(err.Error())
		}

		id, err := result.LastInsertId()
		if err != nil {
			log.Panic(err.Error())
		}

		fmt.Println("Successfully insert comment with id", id)
	}

	err = tx.Rollback()
	if err != nil {
		log.Panic(err.Error())
	}
}
