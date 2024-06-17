package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Pastry struct {
	Name             string
	Description      string
	Price            string
	UnitOfMeasure    string
	QuantityPerPiece string
}

func GetPastries(response http.ResponseWriter, request *http.Request) {
	rows, err := db.Query("select name, price from pastry")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	pastriesFromDB := map[string]string{}
	for rows.Next() {
		var name string
		var price string
		err = rows.Scan(&name, &price)
		if err != nil {
			log.Fatal(err)
		}
		pastriesFromDB[name] = price
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	encoder := json.NewEncoder(response)
	encoder.SetIndent("", "  ")
	encoder.Encode(pastriesFromDB)
}
