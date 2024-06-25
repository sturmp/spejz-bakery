package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
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

var initialSchedules []BakingSchedule = []BakingSchedule{
	{"kakaós csiga", 12, 5, time.Now()},
	{"kenyér", 2, 1, time.Now()},
	{"kenyér", 2, 2, time.Now().AddDate(0, 0, 1)},
	{"Heti különlegesség", 15, 5, time.Now().AddDate(0, 0, 1)},
	{"Biscuit", 0.5, 0.5, time.Now().AddDate(0, 0, 1)},
}

var db *sql.DB

func main() {
	if _, err := os.Stat("bakery.db"); err != nil {
		db = InitDb()
	} else {
		db = ConnectToDb()
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/pastry", GetPastries).Methods("GET")
	router.HandleFunc("/order", GetOrders).Methods("GET")
	router.HandleFunc("/order", CreateOrder).Methods("POST")
	router.HandleFunc("/schedule", GetBakingSchedules).Methods("GET")

	handler := cors.Default().Handler(router)
	http.ListenAndServe(":5555", handler)
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
        create table pastry(name text not null primary key,
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
		quantity real,
		preferedDate text);
		`
	_, err = db.Exec(createOrderTableStatement)
	if err != nil {
		log.Printf("%q: %s\n", err, createOrderTableStatement)
	}

	InsertOrders(db)

	createBakingScheduleTableStatement := `
	create table bakingschedule(pastry text not null,
		quantity real,
		reserved real,
		readyDate text,
		PRIMARY KEY(pastry, readyDate),
		FOREIGN KEY(pastry) REFERENCES pastry(name));
	`
	_, err = db.Exec(createBakingScheduleTableStatement)
	if err != nil {
		log.Printf("%q: %s\n", err, createBakingScheduleTableStatement)
	}

	InsertBakingSchedules(db)

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

func InsertBakingSchedules(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare(`insert into
		bakingschedule(pastry, quantity, reserved, readyDate)
        values(?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for i := 0; i < len(initialSchedules); i++ {
		_, err := stmt.Exec(initialSchedules[i].Pastry, initialSchedules[i].Quantity, initialSchedules[i].Reserved, initialSchedules[i].ReadyDate.Format(time.RFC3339))
		if err != nil {
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
