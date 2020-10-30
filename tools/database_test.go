package tools_test

import (
	"git.agiletech.de/AgileRCM/support-tools/tools"
	"testing"
)

func TestBuildConnectionStringMysql(t *testing.T) {
	conntring, _ := tools.BuildConnectionString("password", "localhost", "user", "mysql", "3306", "testdb", "require")

	if conntring != "mysql://user:password@localhost:3306/testdb?query"{
		t.Errorf("String content was wrong. Want: mysql://user:password@localhost:3306/testdb?query But got: %s", conntring)
	}
}

func TestBuildConnectionStringPostgres(t *testing.T) {
	conntring, _ := tools.BuildConnectionString("password", "localhost", "user", "postgres", "5432", "testdb", "disable")

	if conntring != "postgres://user:password@localhost:5432/testdb?sslmode=disable"{
		t.Errorf("String content was wrong. Want: postgres://user:password@localhost:5432/testdb?sslmode=disable But got: %s", conntring)
	}
}

