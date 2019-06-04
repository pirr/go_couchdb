package couchdb

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Database struct {
	Name   *string
	Url    *string
	Server *Server
}

func (db *Database) Save(data map[string]interface{}) error {
	jsonData, err := json.Marshal(data)
	check(err)
	jsonDataBytes := bytes.NewReader(jsonData)
	response, err := http.Post(*db.Url, "application/json", jsonDataBytes)
	check(err)
	checkExists(response)
	defer response.Body.Close()
	return nil
}

func (db *Database) AllDocs() (map[string]interface{}, error) {
	var result map[string]interface{}
	response, err := http.Get(*db.Url + `/_all_docs`)
	check(err)
	checkExists(response)
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	check(err)
	err = json.Unmarshal(data, &result)
	return result, err
}
