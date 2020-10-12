package tools

import (
	"fmt"
	"git.agiletech.de/AgileRCM/support-tools/context"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"net/url"
)

func MigrateDB(ctx context.Context) error {

	connectionString, err := buildConnectionString(ctx)

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

func DownDB(ctx context.Context) error {

	connectionString, err := buildConnectionString(ctx)

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

func buildConnectionString(ctx context.Context) (string, error) {

	var connectionstring string

	dbpassword := fmt.Sprint(url.PathEscape(ctx.DB_Password))
	dbserver := fmt.Sprint(url.PathEscape(ctx.DB_Server))
	dbuser := fmt.Sprint(url.PathEscape(ctx.DB_User))
	dbtype := fmt.Sprint(url.PathEscape(ctx.DB_Type))
	dbport := fmt.Sprint(url.PathEscape(ctx.DB_Port))
	dbname := fmt.Sprint(url.PathEscape(ctx.DB_Name))

	switch db_type := ctx.DB_Type; db_type {
	case "postgres":
		fmt.Println("Using Postgres DB")
		connectionstring = dbtype + "://" + dbuser + ":" + dbpassword + "@" + dbserver + ":" + dbport + "/" + dbname + "?sslmode=disable"
	case "mysql":
		fmt.Println("Using MySQL DB")
		connectionstring = dbtype + "://" + dbuser + ":" + dbpassword + "@" + dbserver + ":" + dbport + "/" + dbname + "?query"
	default:
		fmt.Println("Cannot build connection string!")
		return "", fmt.Errorf("Cannot build connection string!")
	}

	return connectionstring, nil
}
