package couchdb

import (
	"net/http"
	"net/url"
	"path"
)

// Server structure
// URL: database URL
// User: admin name
// Password: admin password
// HTTPClient: http Client
type Server struct {
	URL        string
	User       string
	Password   string
	HTTPClient *http.Client
}

// getDBUrl get url string for database
func getDBUrl(s *Server, dbName *string) string {
	dbURL, err := url.Parse(s.URL)
	check(err)
	dbURL.Path = path.Join(dbURL.Path, *dbName)
	dbURLString := dbURL.String()
	return dbURLString
}

// doRequest do request to database and return response
func (s *Server) doRequest(request *http.Request) *http.Response {
	request.SetBasicAuth(s.User, s.Password)
	response, err := s.HTTPClient.Do(request)
	check(err)
	defer response.Body.Close()
	return response
}

// Create database with dbName
func (s *Server) Create(dbName string) *Database {
	db := &Database{}
	db.Name = &dbName
	db.Server = s
	dbURL := getDBUrl(s, &dbName)
	db.URL = &dbURL
	request, err := http.NewRequest(`PUT`, dbURL, nil)
	check(err)
	s.doRequest(request)
	return db
}

// Delete database with dbName
func (s *Server) Delete(dbName string) error {
	dbURL := getDBUrl(s, &dbName)
	request, err := http.NewRequest(`DELETE`, dbURL, nil)
	check(err)
	response := s.doRequest(request)
	return checkExists(response)
}

// ExistsDB check database existence by dbName
func (s *Server) ExistsDB(dbName string) error {
	dbURL := getDBUrl(s, &dbName)
	request, err := http.NewRequest(`HEAD`, dbURL, nil)
	check(err)
	response := s.doRequest(request)
	return checkExists(response)
}
