package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Order struct {
	Pastry       string
	Customer     string
	Quantity     int
	PreferedDate time.Time
}

func GetOrders(response http.ResponseWriter, request *http.Request) {
	rows, err := db.Query("select pastry, customer, quantity, preferedDate from pastryorder")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	orders := []Order{}
	for rows.Next() {
		var order Order
		var preferedDateText string
		err = rows.Scan(&order.Pastry, &order.Customer, &order.Quantity, &preferedDateText)
		if err != nil {
			log.Fatal(err)
		}
		if preferedDate, err := time.Parse(time.RFC3339, preferedDateText); err == nil {
			order.PreferedDate = preferedDate
		} else {
			log.Println(err)
		}
		orders = append(orders, order)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	encoder := json.NewEncoder(response)
	encoder.SetIndent("", "  ")
	encoder.Encode(orders)
}

func CreateOrder(response http.ResponseWriter, request *http.Request) {
	var order Order

	err := json.NewDecoder(request.Body).Decode(&order)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	InsertOrderToDb(order)

	encoder := json.NewEncoder(response)
	encoder.SetIndent("", "  ")
	encoder.Encode(order)
}

func InsertOrderToDb(order Order) {
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

	_, err = stmt.Exec(order.Pastry, order.Customer, order.Quantity, order.PreferedDate.Format(time.RFC3339))
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
