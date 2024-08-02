package pastry

import (
	"api/internal/utility"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

var DB *sql.DB

type Pastry struct {
	Id               int
	Name             string
	Description      string
	Price            string
	UnitOfMeasure    string
	QuantityPerPiece string
}

func GetPastries(response http.ResponseWriter, request *http.Request) {
	rows, err := DB.Query("SELECT id, name, description, price, unitofmeasure, quantityperpiece FROM pastry")
	if err != nil {
		utility.LogAndErrorResponse(err, response)
	}
	defer rows.Close()

	pastriesFromDB := []Pastry{}
	for rows.Next() {
		var pastry Pastry
		err = rows.Scan(&pastry.Id, &pastry.Name, &pastry.Description, &pastry.Price, &pastry.UnitOfMeasure, &pastry.QuantityPerPiece)
		if err != nil {
			utility.LogAndErrorResponse(err, response)
		}
		pastriesFromDB = append(pastriesFromDB, pastry)
	}
	err = rows.Err()
	if err != nil {
		utility.LogAndErrorResponse(err, response)
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
		utility.LogAndErrorResponse(err, response)
		return
	}

	stmt, err := tx.Prepare("UPDATE pastry SET name=?, description=?, price=?, unitofmeasure=?, quantityperpiece=? WHERE id=?")
	if err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(pastry.Name, pastry.Description, pastry.Price, pastry.UnitOfMeasure, pastry.QuantityPerPiece, pastry.Id)
	if err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}

	if err := tx.Commit(); err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}

	encoder := json.NewEncoder(response)
	encoder.SetIndent("", "  ")
	encoder.Encode(pastry)
}

func FetchPastryName(pastryId int) string {
	var pastryName string
	row := DB.QueryRow("SELECT name FROM pastry WHERE id=?", pastryId)
	if err := row.Scan(&pastryName); err != nil {
		log.Fatal(err)
	}
	return pastryName
}
