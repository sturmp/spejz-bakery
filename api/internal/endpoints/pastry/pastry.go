package pastry

import (
	"api/internal/utility"
	"database/sql"
	"encoding/json"
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
	Enabled          bool
}

type CreatePastryRequest struct {
	Name             string
	Description      string
	Price            string
	UnitOfMeasure    int
	QuantityPerPiece string
}

func GetPastries(response http.ResponseWriter, request *http.Request) {
	languageCode := utility.GetLanguageOrDefault(request)
	rows, err := DB.Query(`SELECT pastry.id,
		pastrytranslation.name,
		pastrytranslation.description,
		pastry.price,
		unitofmeasuretranslation.name,
		pastry.quantityperpiece,
		pastry.enabled
		FROM pastry
			JOIN pastrytranslation ON pastry.id = pastrytranslation.pastryid
				AND pastrytranslation.language = ?
			JOIN unitofmeasuretranslation ON pastry.unitofmeasure = unitofmeasuretranslation.unitofmeasureid
				AND unitofmeasuretranslation.language = ?`, languageCode, languageCode)
	if err != nil {
		utility.LogAndErrorResponse(err, response)
	}
	defer rows.Close()

	pastriesFromDB := []Pastry{}
	for rows.Next() {
		var pastry Pastry
		err = rows.Scan(&pastry.Id,
			&pastry.Name,
			&pastry.Description,
			&pastry.Price,
			&pastry.UnitOfMeasure,
			&pastry.QuantityPerPiece,
			&pastry.Enabled)
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
	languageCode := utility.GetLanguageOrDefault(request)
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
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE pastry SET price=?, quantityperpiece=?, enabled=? WHERE id=?",
		pastry.Price,
		pastry.QuantityPerPiece,
		pastry.Enabled,
		pastry.Id)
	if err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}

	_, err = tx.Exec("UPDATE pastrytranslation SET name=?, description=? WHERE pastryid=? AND language=?",
		pastry.Name,
		pastry.Description,
		pastry.Id,
		languageCode)
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

func CreatePastry(response http.ResponseWriter, request *http.Request) {
	var pastry CreatePastryRequest
	if err := json.NewDecoder(request.Body).Decode(&pastry); err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	tx, err := DB.Begin()
	if err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}
	defer tx.Rollback()

	pastryId, err := insertPastry(tx, pastry)
	if err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}

	err = insertPastryLanguages(tx, pastry, pastryId)
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

func FetchPastryName(pastryId int) (string, error) {
	var pastryName string
	row := DB.QueryRow(`SELECT name
		FROM pastry
			JOIN pastrytranslation ON pastry.id = pastrytranslation.pastryid
				AND pastrytranslation.language = "en"
		WHERE id=?`, pastryId)
	if err := row.Scan(&pastryName); err != nil {
		return "", err
	}
	return pastryName, nil
}

func insertPastry(tx *sql.Tx, pastry CreatePastryRequest) (pastryId int, err error) {
	err = tx.QueryRow(`INSERT INTO
		pastry(price, quantityperpiece, unitofmeasure)
		VALUES(?, ?, ?)
		RETURNING id`,
		pastry.Price,
		pastry.QuantityPerPiece,
		pastry.UnitOfMeasure).Scan(&pastryId)
	if err != nil {
		return 0, err
	}
	return pastryId, nil
}

func insertPastryLanguages(tx *sql.Tx, pastry CreatePastryRequest, pastryId int) error {
	rows, err := tx.Query("SELECT DISTINCT language FROM pastrytranslation")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var language string
		if err := rows.Scan(&language); err != nil {
			return err
		}

		_, err = tx.Exec(`INSERT INTO
			pastrytranslation(language, pastryid, name, description)
			VALUES(?, ?, ?, ?)`,
			language,
			pastryId,
			pastry.Name,
			pastry.Description)
		if err != nil {
			return err
		}
	}
	return nil
}
