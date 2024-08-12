package dayoff

import (
	"api/internal/utility"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var DB *sql.DB

type DayOff struct {
	Id  int
	Day time.Time
}

func GetDayOffs(response http.ResponseWriter, request *http.Request) {
	rows, err := DB.Query("select id, day from dayoff")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	dayOffs := []DayOff{}
	for rows.Next() {
		var dayOff DayOff
		var dayString string
		if err = rows.Scan(&dayOff.Id, &dayString); err != nil {
			utility.LogAndErrorResponse(err, response)
			return
		}

		if day, err := time.Parse(time.RFC3339, dayString); err != nil {
			utility.LogAndErrorResponse(err, response)
			return
		} else {
			dayOff.Day = day
			dayOffs = append(dayOffs, dayOff)
		}
	}
	if err = rows.Err(); err != nil {
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

	tx, err := DB.Begin()
	if err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`delete from dayoff where id = ?`)
	if err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(dayOffId)
	if err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}

	err = tx.Commit()
	if err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}
}

func CreateDayOff(response http.ResponseWriter, request *http.Request) {
	var dayString string
	if err := json.NewDecoder(request.Body).Decode(&dayString); err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}

	day, err := time.Parse(time.RFC3339, dayString)
	if err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}

	tx, err := DB.Begin()
	if err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`insert into dayoff(day) values(?);`)
	if err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(day.Format(time.RFC3339))
	if err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}

	err = tx.Commit()
	if err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}

	id, err := result.LastInsertId()
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
