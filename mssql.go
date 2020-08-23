package main

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func Mssql() {
	dsn := "sqlserver://sa:Zps05.....@localhost:1433?database=master"
	_, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		println(err.Error())
	} else {
		println("connected")
	}
}
