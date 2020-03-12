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
  return testDB
}

func tearDown(testDB *gorm.DB) {
  testDB.DB.Close()
}
