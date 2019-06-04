package couchdb

import (
	"log"
	"net/http"
)

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
