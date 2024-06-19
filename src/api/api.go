package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var pastries []Pastry = []Pastry{
	{"Biscuit", "Nem az az édes, ez sós... Vajas pogácsa jó?! Vajas pogácsa!", "3500", "kg", ""},
	{"Foccacia", "Olasz olajos kenyér lángos... Feltét nélkül.", "~440", "db", "~125g"},
	{"kenyér", "Sima kenyér. Semmi extra.", "600", "db", "750g"},
	{"English muffin", "Nem, ez nem az édesség. <a href=\"https://www.google.com/search?client=firefox-b-d&q=english+muffin\">Nézz utána!</a>", "150", "db", "65g"},
	{"Kakaós csiga", "Kakaós és fel van tekerve.", "500", "db", "100g"},
	{"Tortilla", "Mexikói lapos lángos. Kaja origami.", "100", "db", "~20cm"},
	{"Heti különlegesség", "Kísérleti sütések jól sikerült egyedei.", "TBD", "db", ""},
}

var initialOrders []Order = []Order{
	{"kenyér", "Zizi", 1, time.Now().AddDate(0, 0, 1)},
	{"kakaós csiga", "Andi", 5, time.Now().AddDate(0, 0, 2)},
	{"English muffin", "Roland", 14, time.Now().AddDate(0, 0, 2)},
}

var db *sql.DB

func corsAllowAllOrigin(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		f(w, r)
	}
}

func main() {
	if _, err := os.Stat("bakery.db"); err != nil {
		db = InitDb()
	} else {
		db = ConnectToDb()
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/pastry", corsAllowAllOrigin(GetPastries)).Methods("GET")
	router.HandleFunc("/order", corsAllowAllOrigin(GetOrders)).Methods("GET")
	router.HandleFunc("/order", corsAllowAllOrigin(CreateOrder)).Methods("POST")

	http.ListenAndServe(":5555", router)
}

func ConnectToDb() *sql.DB {
	db, err := sql.Open("sqlite3", "bakery.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func InitDb() *sql.DB {
	if _, err := os.Stat("bakery.db"); err == nil {
		return nil
	}

	db := ConnectToDb()

	createPastryTableStatement := `
        create table pastry(id integer not null primary key,
            name text,
            description text,
            price integer,
            unitofmeasure text,
            quantityperpiece text);
    `
	_, err := db.Exec(createPastryTableStatement)
	if err != nil {
		log.Printf("%q: %s\n", err, createPastryTableStatement)
	}

	InsertPastries(db)

	createOrderTableStatement := `
	create table pastryorder(id integer not null primary key,
		pastry text,
		customer text,
		quantity integer,
		preferedDate text);
		`
	_, err = db.Exec(createOrderTableStatement)
	if err != nil {
		log.Printf("%q: %s\n", err, createOrderTableStatement)
	}

	InsertOrders(db)

	return db
}

func InsertPastries(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare(`insert into
        pastry(name, description, price, unitofmeasure, quantityperpiece)
        values(?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for i := 0; i < len(pastries); i++ {
		_, err := stmt.Exec(pastries[i].Name, pastries[i].Description, pastries[i].Price, pastries[i].UnitOfMeasure, pastries[i].QuantityPerPiece)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func InsertOrders(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare(`insert into
		pastryorder(pastry, customer, quantity, preferedDate)
        values(?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for i := 0; i < len(initialOrders); i++ {
		_, err := stmt.Exec(initialOrders[i].Pastry, initialOrders[i].Customer, initialOrders[i].Quantity, initialOrders[i].PreferedDate.Format(time.RFC3339))
		if err != nil {
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
