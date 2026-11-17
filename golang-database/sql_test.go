package golangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	
	sqlQuery := "INSERT INTO customers(id, name) VALUES('udin', 'Udin')"
	_, err := db.ExecContext(ctx, sqlQuery)
	if err!= nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT * FROM customers"
	rows, err :=  db.QueryContext(ctx, script)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("ID", id)
		fmt.Println("Name", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customers"
	rows, err :=  db.QueryContext(ctx, script)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt )
		if err != nil {
			panic(err)
		}

		fmt.Println("=====================")
		fmt.Println("id", id)
		fmt.Println("name", name)
		if email.Valid {
			fmt.Println("email", email.String)
		}
		fmt.Println("balance", balance)
		fmt.Println("rating", rating)
		if birthDate.Valid {
			fmt.Println("birthDate", birthDate.Time)
		}
		fmt.Println("createdAt", createdAt)
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; --"
	password := "salah"

	// script := "SELECT username FROM users WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	sqlQuery := "SELECT username FROM users WHERE username = $1 AND password = $2 LIMIT 1"
	fmt.Println(sqlQuery)
	rows, err :=  db.QueryContext(ctx, sqlQuery, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Println("Sukses login", username)
		} else {
			fmt.Println("Gagal login")
		}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "udin' DROP TABLE users; --"
	password := "rahasia"
	
	sqlQuery := "INSERT INTO users(username, password) VALUES($1, $2)"
	_, err := db.ExecContext(ctx, sqlQuery, username, password)
	if err!= nil {
		panic(err)
	}

	fmt.Println("Success insert new user")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "udin@gmail.com"
	comment := "your comment"
	
	sqlQuery := `INSERT INTO comments(email, comment) VALUES($1, $2) RETURNING id`
	result, err := db.ExecContext(ctx, sqlQuery, email, comment)
	if err!= nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id ", id)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	sqlQuery := `INSERT INTO comments(email, comment) VALUES($1, $2) RETURNING id`
	stmt, err := db.PrepareContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "udin" + strconv.Itoa(i) + "@gmail.com";
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("success insert comment with id ", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	sqlQuery := `INSERT INTO comments(email, comment) VALUES($1, $2) RETURNING id`
	// do transaction
	for i := 0; i < 10; i++ {
		email := "udin" + strconv.Itoa(i) + "@gmail.com";
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, sqlQuery, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()

		fmt.Println("success insert comment with id ", id)
	}


	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}