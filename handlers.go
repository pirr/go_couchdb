package couchdb

import (
	"log"
	"net/http"
)

// NotExists struct used in Error handler
type NotExists struct {
	Op     string
	DBName string
}

func (e *NotExists) Error() string { return `Not exists` + ` ` + e.DBName }

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkExists(response *http.Response) error {
	if response.StatusCode != 200 {
		url := response.Request.URL
		return &NotExists{`Not exists`, url.Path}
	}
	return nil
}
