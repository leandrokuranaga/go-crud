package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDb() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	connString := os.Getenv("CONNECTION_STRINGS")
	if connString == "" {
		log.Fatal("The CONNECTION_STRINGS environment variable is not defined or empty")
	}

	var err error
	DB, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatalf("Error opening connection to the database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	DB.SetConnMaxIdleTime(5)

	createTables()
	return DB, nil
}

func createTables() {

	createUsersTable := `
    IF NOT EXISTS (SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'users')
    BEGIN
		CREATE TABLE users(
			id INT PRIMARY KEY IDENTITY,
			email NVARCHAR(255) NOT NULL UNIQUE,
			password NVARCHAR(MAX) NOT NULL
		);
    END;
    `

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table")
	}

	sqlStmt := `
    IF NOT EXISTS (SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'events')
    BEGIN
        CREATE TABLE events (
            id INT PRIMARY KEY IDENTITY,
            name NVARCHAR(MAX) NOT NULL,
            description NVARCHAR(MAX) NOT NULL,
            location NVARCHAR(MAX) NOT NULL,
            dateTime DATETIME NOT NULL,
            user_id INT
            CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id)
        );
    END;
    `
	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Error creating events table: %v", err)
	}

	createRegistrationsTable := `
    IF NOT EXISTS (SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'registrations')
    BEGIN
    CREATE TABLE registrations(
        id INT PRIMARY KEY IDENTITY,
        event_id INTEGER,
        user_id INTEGER,
        CONSTRAINT fk_registrations_users_id FOREIGN KEY (user_id) REFERENCES users(id),
        CONSTRAINT fk_registrations_event_id FOREIGN KEY (event_id) REFERENCES events(id)
    );
    END;
    `

	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		log.Fatalf("Error creating events table: %v", err)
	}
}
