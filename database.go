package couchdb

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Database meta
type Database struct {
	Name   *string
	URL    *string
	Server *Server
}

// Save document to database
func (db *Database) Save(data map[string]interface{}) error {
	jsonData, err := json.Marshal(data)
	check(err)
	jsonDataBytes := bytes.NewReader(jsonData)
	response, err := http.Post(*db.URL, "application/json", jsonDataBytes)
	check(err)
	checkExists(response)
	defer response.Body.Close()
	return nil
}

// GetDocByID return doc by _id
func (db *Database) GetDocByID(id string) (map[string]interface{}, error) {
	var result map[string]interface{}
	response, err := http.Get(*db.URL + "/" + id)
	check(err)
	checkExists(response)
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	check(err)
	err = json.Unmarshal(data, &result)
	return result, err
}

// AllDocs return all docs from database
func (db *Database) AllDocs() (map[string]interface{}, error) {
	result, err := db.GetDocByID("_all_docs")
	check(err)
	return result, err
}
