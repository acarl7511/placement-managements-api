package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func initMigration() {
	var (
		c company
		f fields
	)
	psqlconnector := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode = disable", host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlconnector)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&c, &f)
	defer db.Close()
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EndPoint Hit: hello")
}

func getplacement(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getPlacement")

	psqlconnector := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode = disable", host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlconnector)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()

	var c []company
	db.Model(&c).Find(&c)
	for _, u := range c {
		byteArray, e := json.MarshalIndent(u, "", " ")
		if e != nil {
			panic("Error Marshalling Response...")
		}
		fmt.Fprint(w, string(byteArray))
	}
}

func postplacement(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: postPlacement")
	var c company
	fetchedData, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(fetchedData, &c)

	psqlconnector := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode = disable", host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlconnector)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()

	db.Create(&c)
	byteArray, e := json.MarshalIndent(c, "", " ")
	if e != nil {
		panic("Unable to Marshal response")
	}
	fmt.Fprint(w, string(byteArray), "\nSuccessfully Added...")
}

func putplacement(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: putPlacement")

	psqlconnector := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode = disable", host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlconnector)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()

	var list []company
	db.Model(&list).Find(&list)

	var c company
	fetchedData, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(fetchedData, &c)
	search := c.Name

	for _, i := range list {
		if i.Name == search {
			db.Model(&i).Update(&c)
			fmt.Fprint(w, "Successfully Updated...")
		}
	}
}

func deleteplacement(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deletePlacement")

	psqlconnector := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode = disable", host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlconnector)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()

	var list []company
	db.Model(&list).Find(&list)

	var c company
	fetchedData, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(fetchedData, &c)
	search := c.Name

	for _, i := range list {
		if i.Name == search {
			db.Model(&i).Delete(&i)
			fmt.Fprint(w, "Successfully Deleted...")
		}
	}
}
