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
	rows, err := db.Query("select name, description, price, unitofmeasure, quantityperpiece from pastry")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	pastriesFromDB := []Pastry{}
	for rows.Next() {
		var pastry Pastry
		err = rows.Scan(&pastry.Name, &pastry.Description, &pastry.Price, &pastry.UnitOfMeasure, &pastry.QuantityPerPiece)
		if err != nil {
			log.Fatal(err)
		}
		pastriesFromDB = append(pastriesFromDB, pastry)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	encoder := json.NewEncoder(response)
	encoder.SetIndent("", "  ")
	encoder.Encode(pastriesFromDB)
}
