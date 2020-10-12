package context

import (
	"io"
	"os"
)

// TODO: Beim schreiben der Doku fällt auf: User/Pass/Endpoint etc. brauchen wir für ALLE tools, nicht nur confluence, hier muss also noch spezifiziert werden werden

// Context struct to hold the config values that will be filled and then used later in functions of the toolkit for example
// to access instances like confluence, jira etc
type Context struct {
	DB_User     string
	DB_Password string
	DB_Server   string
	DB_Name     string
	DB_Type     string
	DB_Port     string
	UserId      string
	Password    string
	Vault       string
	Writer      io.Writer
	Endpoint    string
	Spacekey    string
	Insecure    bool
	Debug       bool
}

// Create a new Context using the Context struct that can then be used in the toolkit
func NewContext() *Context {
	return &Context{
		Writer:   os.Stdout,
		Insecure: false,
		Debug:    false,
	}
}
