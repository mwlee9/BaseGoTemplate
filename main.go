package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func database() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (ID integer Primary Key, Name STRING)")

	if err != nil {
		fmt.Println(err)
	}

}

func home(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("views/home.html", "partials/head.html", "partials/foot.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	t.Execute(w, "Home")

}

func testform(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	r.ParseForm()
	fmt.Println(r.Form["user"])
}

func main() {
	database()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", home)
	http.HandleFunc("/testform", testform)

	http.ListenAndServe(":8000", nil)
}
