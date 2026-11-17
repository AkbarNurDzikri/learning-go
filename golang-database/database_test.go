package golangdatabase

import (
	"database/sql"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestEmpty(t *testing.T) {}

func TestOpenConnection(t *testing.T) {
	// db connection based on pgx docs
	// conn, err := pgx.Connect(context.Background(), "postgres://postgres:Krw.030894@localhost:5432/go_db")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer conn.Close(context.Background())

	db, err := sql.Open("pgx", "postgres://postgres:Krw.030894@localhost:5432/go_db")
	if err != nil {
		panic(err)
	}

	defer db.Close()
}