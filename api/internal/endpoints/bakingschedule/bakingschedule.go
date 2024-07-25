package bakingschedule

import (
	"api/internal/utility"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var DB *sql.DB

type BakingSchedule struct {
	Pastry    string
	Quantity  float32
	Reserved  float32
	ReadyDate time.Time
}

func GetBakingSchedules(response http.ResponseWriter, request *http.Request) {
	bakingschedules := FetchSchedulesFromDB()

	encoder := json.NewEncoder(response)
	encoder.SetIndent("", "  ")
	encoder.Encode(bakingschedules)
}

func CreateBakingSchedules(response http.ResponseWriter, request *http.Request) {
	var bakingSchedule BakingSchedule

	if err := json.NewDecoder(request.Body).Decode(&bakingSchedule); err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
	}

	insertScheduleToDb(bakingSchedule)

	encoder := json.NewEncoder(response)
	encoder.SetIndent("", "  ")
	encoder.Encode(bakingSchedule)
}

func UpdateBakingSchedule(response http.ResponseWriter, request *http.Request) {
	var bakingSchedule BakingSchedule

	if err := json.NewDecoder(request.Body).Decode(&bakingSchedule); err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
	}

	bakingSchedule.ReadyDate = bakingSchedule.ReadyDate.Local()

	tx, err := DB.Begin()
	if err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}

	stmt, err := tx.Prepare("update bakingschedule set quantity=?, reserved=? where pastry=? and readyDate=?")
	if err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(bakingSchedule.Quantity, bakingSchedule.Reserved, bakingSchedule.Pastry, bakingSchedule.ReadyDate.Format(time.RFC3339))
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
	encoder.Encode(bakingSchedule)
}

func insertScheduleToDb(bakingSchedule BakingSchedule) {
	tx, err := DB.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare(`insert into
		bakingschedule(pastry, quantity, reserved, readyDate)
        values(?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(bakingSchedule.Pastry, bakingSchedule.Quantity, bakingSchedule.Reserved, bakingSchedule.ReadyDate.Format(time.RFC3339))
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateScheduleReservedInDB(schedule BakingSchedule) {
	tx, err := DB.Begin()
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt, err := tx.Prepare(`update bakingschedule set reserved = ? where pastry = ? and readydate = ?`)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(schedule.Reserved, schedule.Pastry, schedule.ReadyDate.Format(time.RFC3339))
	if err != nil {
		log.Fatal(err.Error())
	}

	if err = tx.Commit(); err != nil {
		log.Fatal(err.Error())
	}
}

func FetchSchedulesFromDB() []BakingSchedule {
	rows, err := DB.Query("select pastry, quantity, reserved, readydate from bakingschedule")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	bakingschedules := []BakingSchedule{}
	for rows.Next() {
		var bakingschedule BakingSchedule
		var readyDateText string
		err = rows.Scan(&bakingschedule.Pastry, &bakingschedule.Quantity, &bakingschedule.Reserved, &readyDateText)
		if err != nil {
			log.Fatal(err)
		}
		if readyDate, err := time.Parse(time.RFC3339, readyDateText); err == nil {
			bakingschedule.ReadyDate = readyDate
		} else {
			log.Println(err)
		}
		bakingschedules = append(bakingschedules, bakingschedule)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return bakingschedules
}
