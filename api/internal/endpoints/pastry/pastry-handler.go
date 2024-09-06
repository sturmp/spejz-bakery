package pastry

import (
	"api/internal/utility"
	"encoding/json"
	"net/http"
)

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

type PastryRepository interface {
	FetchAllPastries(languageCode string) ([]Pastry, error)
	UpdatePastry(pastry Pastry, languageCode string) error
	CreatePastry(createRequest CreatePastryRequest, languageCode string) (Pastry, error)
}

var Repository PastryRepository

func GetPastries(response http.ResponseWriter, request *http.Request) {
	languageCode := utility.GetLanguageOrDefault(request)

	pastries, err := Repository.FetchAllPastries(languageCode)
	enabledPastries := []Pastry{}
	for _, pastry := range pastries {
		if pastry.Enabled {
			enabledPastries = append(enabledPastries, pastry)
		}
	}

	if err != nil {
		utility.LogAndErrorResponse(err, response)
	}

	encoder := json.NewEncoder(response)
	encoder.SetIndent("", "  ")
	encoder.Encode(enabledPastries)
}

func GetAllPastries(response http.ResponseWriter, request *http.Request) {
	languageCode := utility.GetLanguageOrDefault(request)

	pastriesFromDB, err := Repository.FetchAllPastries(languageCode)
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
		utility.LogAndErrorResponseWithCode(err, response, http.StatusBadRequest)
		return
	}

	if err := Repository.UpdatePastry(pastry, languageCode); err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}

	encoder := json.NewEncoder(response)
	encoder.SetIndent("", "  ")
	encoder.Encode(pastry)
}

func CreatePastry(response http.ResponseWriter, request *http.Request) {
	languageCode := utility.GetLanguageOrDefault(request)
	var createRequest CreatePastryRequest
	if err := json.NewDecoder(request.Body).Decode(&createRequest); err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	pastry, err := Repository.CreatePastry(createRequest, languageCode)
	if err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}

	encoder := json.NewEncoder(response)
	encoder.SetIndent("", "  ")
	encoder.Encode(pastry)
}
