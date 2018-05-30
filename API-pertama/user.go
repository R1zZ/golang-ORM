package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB //deklarasi variabel Database dari package gorm
var err error

//detail dari isi gorm.DB
type User struct {
	gorm.Model
	Name  string
	Email string
}

//function untuk konek DB
func InitialMigration() {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Gagal Konek ke databases")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

//fungsi untuk menampilkan list data user
func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("tidak dapat terhubung ke databases")
	}
	defer db.Close()

	var users []User // pembuatan dumi untuk json data user
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("tidak dapat terhubung ke databases")
	}
	defer db.Close()

	vars := mux.Vars(r) //deklarasi user baru dengan var name dan email
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email}) //write name & email dengan dumi nya

	fmt.Fprintf(w, "User baru berhasil dibuat")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("tidak dapat terhubung ke databases")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user) //seleksi berdasarkan user
	db.Delete(&user)                       //penghapusan menurut seleksi name

	fmt.Fprintf(w, "Penghapusan berhasil")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("tidak dapat terhubung ke databases")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)

	user.Email = email //update email

	db.Save(&user)
	fmt.Fprintf(w, "update user Berhasil")
}
