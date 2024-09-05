package dayoff

import (
	"api/internal/utility"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var Repository DayOffRepository

type DayOff struct {
	Id  int
	Day time.Time
}

type DayOffRepository interface {
	FetchDayOffs() ([]DayOff, error)
	DeleteDayOff(id int) error
	CreateDayOff(day time.Time) (int64, error)
}

func RegisterHandler(router *mux.Router, db *sql.DB) {
	Repository = newDayOffRepository(db)

	router.HandleFunc("/dayoff", GetDayOffs).Methods("GET")
	router.HandleFunc("/dayoff", CreateDayOff).Methods("POST")
	router.HandleFunc("/dayoff/{id}", DeleteDayOff).Methods("DELETE")
}

func GetDayOffs(response http.ResponseWriter, request *http.Request) {
	dayOffs, err := Repository.FetchDayOffs()
	if err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}

	encoder := json.NewEncoder(response)
	encoder.SetIndent("", " ")
	encoder.Encode(dayOffs)
}

func DeleteDayOff(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	dayOffId, ok := vars["id"]
	if !ok {
		http.Error(response, "Invalid dayoff id!", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(dayOffId)
	if err != nil {
		http.Error(response, "Invalid dayoff id!", http.StatusBadRequest)
	}

	err = Repository.DeleteDayOff(id)
	if err != nil {
		utility.LogAndErrorResponse(err, response)
	}
}

func CreateDayOff(response http.ResponseWriter, request *http.Request) {
	var dayString string
	if err := json.NewDecoder(request.Body).Decode(&dayString); err != nil {
		utility.LogAndErrorResponseWithCode(err, response, http.StatusBadRequest)
		return
	}

	day, err := time.Parse(time.RFC3339, dayString)
	if err != nil {
		utility.LogAndErrorResponseWithCode(err, response, http.StatusBadRequest)
		return
	}

	id, err := Repository.CreateDayOff(day)
	if err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}

	dayOff := DayOff{
		Id:  int(id),
		Day: day,
	}

	encoder := json.NewEncoder(response)
	encoder.SetIndent("", "  ")
	encoder.Encode(dayOff)
}
