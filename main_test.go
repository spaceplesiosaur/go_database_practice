package main

import (
  "testing"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
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

func testEndpoints(t *testing.T) {
  testDB := setUp()

  tearDown(testDB)
}
