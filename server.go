package couchdb

import (
	"net/http"
	"net/url"
	"path"
)

type Server struct {
	Url        string
	User       string
	Password   string
	HttpClient *http.Client
}

type NotExists struct {
	Op     string
	DBName string
}

func (e *NotExists) Error() string { return `Not exists` + ` ` + e.DBName }

func getDBUrl(s *Server, dbName *string) string {
	dbUrl, err := url.Parse(s.Url)
	check(err)
	dbUrl.Path = path.Join(dbUrl.Path, *dbName)
	dbUrlString := dbUrl.String()
	return dbUrlString
}

func (s *Server) doRequest(request *http.Request) *http.Response {
	request.SetBasicAuth(s.User, s.Password)
	response, err := s.HttpClient.Do(request)
	check(err)
	defer response.Body.Close()
	return response
}

func (s *Server) Create(dbName string) *Database {
	db := &Database{}
	db.Name = &dbName
	db.Server = s
	dbUrl := getDBUrl(s, &dbName)
	db.Url = &dbUrl
	request, err := http.NewRequest(`PUT`, dbUrl, nil)
	check(err)
	s.doRequest(request)
	return db
}

func (s *Server) Delete(dbName string) error {
	dbUrl := getDBUrl(s, &dbName)
	request, err := http.NewRequest(`DELETE`, dbUrl, nil)
	check(err)
	response := s.doRequest(request)
	return checkExists(response)
}

func (s *Server) ExistsDB(dbName string) error {
	dbUrl := getDBUrl(s, &dbName)
	request, err := http.NewRequest(`HEAD`, dbUrl, nil)
	check(err)
	response := s.doRequest(request)
	return checkExists(response)
}
