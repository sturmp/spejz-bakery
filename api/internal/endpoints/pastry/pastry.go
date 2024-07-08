package pastry

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

var DB *sql.DB

type Pastry struct {
	Name             string
	Description      string
	Price            string
	UnitOfMeasure    string
	QuantityPerPiece string
}

func GetPastries(response http.ResponseWriter, request *http.Request) {
	rows, err := DB.Query("select name, description, price, unitofmeasure, quantityperpiece from pastry")
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

func UpdatePastry(response http.ResponseWriter, request *http.Request) {
	var pastry Pastry
	if err := json.NewDecoder(request.Body).Decode(&pastry); err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	tx, err := DB.Begin()
	if err != nil {
		LogandErrorResponse(err, response)
		return
	}

	stmt, err := tx.Prepare("update pastry set description=?, price=?, unitofmeasure=?, quantityperpiece=? where name=?")
	if err != nil {
		LogandErrorResponse(err, response)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(pastry.Description, pastry.Price, pastry.UnitOfMeasure, pastry.QuantityPerPiece, pastry.Name)
	if err != nil {
		LogandErrorResponse(err, response)
		return
	}

	if err := tx.Commit(); err != nil {
		LogandErrorResponse(err, response)
		return
	}

	encoder := json.NewEncoder(response)
	encoder.SetIndent("", "  ")
	encoder.Encode(pastry)
}

func LogandErrorResponse(err error, response http.ResponseWriter) {
	log.Println(err.Error())
	http.Error(response, "", http.StatusInternalServerError)
}
