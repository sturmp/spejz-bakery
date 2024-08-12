package unitofmeasure

import (
	"api/internal/utility"
	"database/sql"
	"encoding/json"
	"net/http"
)

var DB *sql.DB

type UnitOfMeasure struct {
	Id   int
	Name string
}

func GetUnitOfMeasures(response http.ResponseWriter, request *http.Request) {
	languageCode := utility.GetLanguageOrDefault(request)
	rows, err := DB.Query("SELECT unitofmeasureid, name FROM unitofmeasuretranslation WHERE language = ?", languageCode)
	if err != nil {
		utility.LogAndErrorResponse(err, response)
	}
	defer rows.Close()

	unitOfMeasuresFromDB := []UnitOfMeasure{}
	for rows.Next() {
		var unitOfMeasure UnitOfMeasure
		err = rows.Scan(&unitOfMeasure.Id, &unitOfMeasure.Name)
		if err != nil {
			utility.LogAndErrorResponse(err, response)
		}
		unitOfMeasuresFromDB = append(unitOfMeasuresFromDB, unitOfMeasure)
	}
	err = rows.Err()
	if err != nil {
		utility.LogAndErrorResponse(err, response)
	}

	encoder := json.NewEncoder(response)
	encoder.SetIndent("", "  ")
	encoder.Encode(unitOfMeasuresFromDB)
}
