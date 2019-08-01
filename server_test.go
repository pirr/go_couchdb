package couchdb_test

import (
	"mypackages/couchdb"
	"net/http"
	"testing"
)

var server = &couchdb.Server{`http://127.0.0.1:5984`, `admin`, `admin`, &http.Client{}}
var db couchdb.Database

func TestCreate(t *testing.T) {
	db = *server.Create(`go-test`)
	if *db.URL != `http://127.0.0.1:5984/go-test` {
		t.Error("Expected http://127.0.0.1:5984/go-test, got ", db.URL)
	}
}

func TestExsits(t *testing.T) {
	TestCreate(t)
	err := server.ExistsDB(`go-test`)
	if err != nil {
		t.Error(err)
	}
}

func TestSave(t *testing.T) {
	TestCreate(t)
	m := make(map[string]interface{})
	m["bar"] = "Bar"
	m["foo"] = map[string]map[string]interface{}{
		"product1": map[string]interface{}{
			"id":       "p01",
			"name":     "name 1",
			"price":    4.5,
			"quantity": 10,
			"arr":      []string{"1", "2"},
		},
		"product2": map[string]interface{}{
			"id":       "p02",
			"name":     "name 3",
			"price":    6,
			"quantity": 4,
		},
		"product3": map[string]interface{}{
			"id":       "p03",
			"name":     "name 3",
			"price":    13,
			"quantity": 9,
		},
	}
	db.Save(m)
}

func TestDelete(t *testing.T) {
	TestCreate(t)
	err := server.Delete(`go-test`)
	if err != nil {
		t.Error(err)
	}
}

func TestAllDocs(t *testing.T) {
	TestCreate(t)
	TestSave(t)
	data, err := db.AllDocs()
	if err != nil {
		t.Error(err)
		TestDelete(t)
	}
	total_rows := data["total_rows"].(float64)
	total_rows_int := int(total_rows)
	if total_rows_int != 1 {
		t.Error("Expected total rows 1 got", total_rows_int)
		TestDelete(t)
	}
	TestDelete(t)
}
