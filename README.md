# go_couchdb
simple client to work with CouchDB 2.x with Go

### Connection:

```var server = &couchdb.Server{"http://127.0.0.1:5984", "admin", "admin", &http.Client{}}```

### Create:
```db = *server.Create("go-test")```

### Save doc:
```
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
m["_id"] = "foo_test"
db.Save(m)
```

### Delete database
```
err := server.Delete(`go-test`)
if err != nil {
  t.Error(err)
}
```

### Find (mango query)
```
q := `
  "selector": {
    "_id": "foo_test"
		}
`
res, err := db.Find(q)
```
