package tools

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"net/url"
)

// Migrate the database completely up. Uses the files located in "folder".
func UpDB(dbPassword string, dbServer string, dbUser string, dbType string, dbPort string, dbName string, folder string, sslMode string) error {

	connectionString, err := BuildConnectionString(dbPassword, dbServer, dbUser, dbType, dbPort, dbName, sslMode)

	if err != nil {
		return err
	}

	fmt.Println("Up migrating database.")
	migration, err := migrate.New(
		"file://"+folder,
		connectionString,
	)

	if err != nil {
		return err
	}

	if err := migration.Up(); err != nil {
		return err
	}

	return nil
}

// Migrate the database completely down. Uses the files located in "folder".
func DownDB(dbPassword string, dbServer string, dbUser string, dbType string, dbPort string, dbName string, folder string, sslMode string) error {

	connectionString, err := BuildConnectionString(dbPassword, dbServer, dbUser, dbType, dbPort, dbName, sslMode)
	if err != nil {
		return err
	}

	fmt.Println("Down migrating database.")
	migration, err := migrate.New(
		"file://"+folder,
		connectionString,
	)

	if err != nil {
		return err
	}

	if err := migration.Down(); err != nil {
		return err
	}

	return nil
}

// Build a connection string for the provided database type.
func BuildConnectionString(dbPassword string, dbServer string, dbUser string, dbType string, dbPort string, dbName string, sslMode string) (string, error) {

	var connectionstring string

	dbPassword = fmt.Sprint(url.PathEscape(dbPassword))
	dbServer = fmt.Sprint(url.PathEscape(dbServer))
	dbUser = fmt.Sprint(url.PathEscape(dbUser))
	dbType = fmt.Sprint(url.PathEscape(dbType))
	dbPort = fmt.Sprint(url.PathEscape(dbPort))
	dbName = fmt.Sprint(url.PathEscape(dbName))

	switch db_type := dbType; db_type {
	case "postgres":
		fmt.Println("Using Postgres DB")
		connectionstring = dbType + "://" + dbUser + ":" + dbPassword + "@" + dbServer + ":" + dbPort + "/" + dbName + "?sslmode=" + sslMode
	case "mysql":
		fmt.Println("Using MySQL DB")
		connectionstring = dbType + "://" + dbUser + ":" + dbPassword + "@" + dbServer + ":" + dbPort + "/" + dbName + "?query"
	default:
		fmt.Println("Cannot build connection string!")
		return "", fmt.Errorf("Cannot build connection string!")
	}

	return connectionstring, nil
}
