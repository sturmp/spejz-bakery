package dayoff

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var DB *sql.DB

func GetDayOffs(response http.ResponseWriter, request *http.Request) {
	rows, err := DB.Query("select  day from dayoff")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	dayOffs := []time.Time{}
	for rows.Next() {
		var dayText string
		if err = rows.Scan(&dayText); err != nil {
			log.Fatal(err)
		}

		if day, err := time.Parse(time.RFC3339, dayText); err != nil {
			log.Fatal(err)
		} else {
			dayOffs = append(dayOffs, day)
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	encoder := json.NewEncoder(response)
	encoder.SetIndent("", " ")
	encoder.Encode(dayOffs)
}
