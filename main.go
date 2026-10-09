package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"time"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Entry struct {
	Name      string
	Email *string
	Message   string
	Timestamp time.Time 
}

var db *sql.DB

func main() {
	username := "user"
	password := "12356"
	ipAddress := "127.0.0.1"
	port := "3060"
	dbName := "Guest_book"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", 
		username, password, ipAddress, port, dbName,
	)

	var err error

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("DB Error:", err)
	}
	fmt.Println("Done")
	handeReqest()
}

func index(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT name, email, message, timestamp FROM entries")
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, "Error:", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var entries []Entry
	for rows.Next() {
		var e Entry
		err := rows.Scan(&e.Name, &e.Email, &e.Message, &e.Timestamp)
		if err != nil {
			log.Println("Error:", err)
			continue
		}
		entries = append(entries, e)
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println("Error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(w, "guest", entries)
}

func handleSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		email := r.FormValue("email")
		message := r.FormValue("message")
		if name != "" && message != "" && email != "" {
			db.Exec("INSERT INTO entries (name, email, message) VALUES (?, ?, ?)", name, email, message)
		}
	}
	http.Redirect(w, r, "guest", http.StatusSeeOther)
}

func handeReqest() {
	http.HandleFunc("/", index)
	http.HandleFunc("/guest", index)
	http.HandleFunc("/submit", handleSubmit)
	log.Println("Running")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
