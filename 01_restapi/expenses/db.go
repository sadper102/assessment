package expenses

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/proullon/ramsql/driver"
)

var db *sql.DB

// const (
// 	host     = "host.docker.internal"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "P@ssw0rd"
// 	dbname   = "golang"
// )

func init() {
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)
	// conn, err := sql.Open("postgres", psqlInfo)
	conn, err := sql.Open(os.Getenv("DATABASE_DRIVER"), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("can't connect to database", err)
	}
	db = conn
	createExpenseTable()
}

func createExpenseTable() {
	createCustomerTable := `
	CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);
	`

	if _, err := db.Exec(createCustomerTable); err != nil {
		log.Fatal("can't create table ", err)
	}
}
