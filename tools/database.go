package tools

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"net/url"
)

func MigrateDB(dbPassword string, dbServer string, dbUser string, dbType string, dbPort string, dbName string) error {

	connectionString, err := buildConnectionString(dbPassword, dbServer, dbUser, dbType, dbPort, dbName)

	if err != nil {
		log.Fatal(err)
	}

	migration, err := migrate.New(
		"file://migrations",
		connectionString,
	)

	if err != nil {
		log.Fatal(err)
	}

	if err := migration.Up(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func DownDB(dbPassword string, dbServer string, dbUser string, dbType string, dbPort string, dbName string) error {

	connectionString, err := buildConnectionString(dbPassword, dbServer, dbUser, dbType, dbPort, dbName)

	if err != nil {
		log.Fatal(err)
	}

	migration, err := migrate.New(
		"file://migrations",
		connectionString,
	)

	if err != nil {
		log.Fatal(err)
	}

	if err := migration.Down(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func buildConnectionString(dbPassword string, dbServer string, dbUser string, dbType string, dbPort string, dbName string) (string, error) {

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
		connectionstring = dbType + "://" + dbUser + ":" + dbPassword + "@" + dbServer + ":" + dbPort + "/" + dbName + "?sslmode=disable"
	case "mysql":
		fmt.Println("Using MySQL DB")
		connectionstring = dbType + "://" + dbUser + ":" + dbPassword + "@" + dbServer + ":" + dbPort + "/" + dbName + "?query"
	default:
		fmt.Println("Cannot build connection string!")
		return "", fmt.Errorf("Cannot build connection string!")
	}

	return connectionstring, nil
}
