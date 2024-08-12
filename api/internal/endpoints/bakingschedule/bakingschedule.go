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
	Pastry struct {
		Id   int
		Name string
	}
	Quantity  float32
	Reserved  float32
	ReadyDate time.Time
}

type UpsertBakingScheduleRequest struct {
	PastryId  int
	Quantity  float32
	Reserved  float32
	ReadyDate time.Time
}

func GetBakingSchedules(response http.ResponseWriter, request *http.Request) {
	languageCode := utility.GetLanguageOrDefault(request)
	bakingschedules, err := FetchSchedulesFromDB(languageCode)
	if err != nil {
		utility.LogAndErrorResponse(err, response)
	}

	encoder := json.NewEncoder(response)
	encoder.SetIndent("", "  ")
	encoder.Encode(bakingschedules)
}

func CreateBakingSchedules(response http.ResponseWriter, request *http.Request) {
	var bakingSchedule UpsertBakingScheduleRequest

	if err := json.NewDecoder(request.Body).Decode(&bakingSchedule); err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
	}

	if err := insertScheduleToDb(bakingSchedule); err != nil {
		utility.LogAndErrorResponse(err, response)
	}

	encoder := json.NewEncoder(response)
	encoder.SetIndent("", "  ")
	encoder.Encode(bakingSchedule)
}

func UpdateBakingSchedule(response http.ResponseWriter, request *http.Request) {
	var bakingSchedule UpsertBakingScheduleRequest

	if err := json.NewDecoder(request.Body).Decode(&bakingSchedule); err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
	}

	tx, err := DB.Begin()
	if err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("UPDATE bakingschedule SET quantity=?, reserved=? WHERE pastryid=? AND readyDate=?")
	if err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(bakingSchedule.Quantity, bakingSchedule.Reserved, bakingSchedule.PastryId, bakingSchedule.ReadyDate.Format(time.RFC3339))
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

func insertScheduleToDb(bakingSchedule UpsertBakingScheduleRequest) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`INSERT INTO
		bakingschedule(pastryid, quantity, reserved, readyDate)
        VALUES(?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(bakingSchedule.PastryId, bakingSchedule.Quantity, bakingSchedule.Reserved, bakingSchedule.ReadyDate.Format(time.RFC3339))
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func UpdateScheduleReservedInDB(schedule BakingSchedule) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`UPDATE bakingschedule SET reserved = ? WHERE pastryid = ? AND readydate = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(schedule.Reserved, schedule.Pastry.Id, schedule.ReadyDate.Format(time.RFC3339))
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func FetchSchedulesFromDB(languageCode string) ([]BakingSchedule, error) {
	rows, err := DB.Query(`SELECT
			bakingschedule.pastryid,
			pastrytranslation.name,
			bakingschedule.quantity,
			bakingschedule.reserved,
			bakingschedule.readydate FROM bakingschedule
		JOIN pastry ON bakingschedule.pastryid = pastry.id
		JOIN pastrytranslation ON bakingschedule.pastryid = pastrytranslation.pastryid
			AND pastrytranslation.language = ?`, languageCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bakingschedules := []BakingSchedule{}
	for rows.Next() {
		var bakingschedule BakingSchedule
		var readyDateText string
		err = rows.Scan(&bakingschedule.Pastry.Id, &bakingschedule.Pastry.Name, &bakingschedule.Quantity, &bakingschedule.Reserved, &readyDateText)
		if err != nil {
			return nil, err
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
		return nil, err
	}
	return bakingschedules, nil
}
