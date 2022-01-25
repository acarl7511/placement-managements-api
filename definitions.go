package main

import "github.com/jinzhu/gorm"

type company struct {
	gorm.Model
	Name       string   `json:"Name"`
	Address    string   `json:"Address"`
	Desc       string   `json:"Desc"`
	Applicants int      `json:"Applicants"`
	Field      []fields `json:"Field" gorm:"many2many:tracking_Depts"`
}

type fields struct {
	gorm.Model
	Dept string `json:"Dept"`
}

const (
	host     = "localhost"
	port     = 8080
	user     = "postgres"
	password = "Decibels52*"
	dbname   = "mineintern"
)
