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
	if *db.Url != `http://127.0.0.1:5984/go-test` {
		t.Error("Expected http://127.0.0.1:5984/go-test, got ", db.Url)
	}
}

func TestExsits(t *testing.T) {
	err := server.ExistsDB(`go-test`)
	if err != nil {
		t.Error(err)
	}
}

func TestSave(t *testing.T) {
	m := make(map[string]interface{})
	m["foo"] = "Foo"
	m["bar"] = "Bar"
	db.Save(m)
}

func TestDelete(t *testing.T) {
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
