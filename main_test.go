package main

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
  // "fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setUp() *gorm.DB {
  var err error

  testDB := &gorm.DB{}
  testDB, err = gorm.Open("sqlite3", ":memory:")
  if err != nil {
		panic(err)
	}
  db.AutoMigrate(&songModel{})

  return testDB
}

func tearDown(testDB *gorm.DB) {
  testDB.Close()
}

func TestAddSong(t *testing.T) {
  testDB := setUp()

  testBody :=  "{\"Title\": \"blue\", \"SpotifyId\": \"12345\", \"URL\": \"www.moose.com\", \"Delay\": 2, \"AvBarDuration\": 3, \"Duration\": 123, \"Tempo\": 4, \"TimeSignature\": 8}"

  stringifiedBody := strings.NewReader(testBody)

  //if I hadn't written this as a string already, I'd have had to use the .Encode method
  router := MakeRouter()

  req, err := http.NewRequest("POST", "/api/v1/songs/", stringifiedBody)
	if err != nil {
		t.Fatal(err)
	}
	// Our API expects a form body, so set the content-type header to make sure it's treated as one.
	req.Header.Add("Content-Type", "application/json")

	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

  // assert.Equal(t, http.StatusOK, rr.Code)
  //use with testify
  if res.Code != http.StatusCreated {
    t.Errorf("ouchie")
  }

  tearDown(testDB)
}
