package dayoff

import (
	"database/sql"
	"time"
)

type DayOffRepositoryImp struct {
	DB *sql.DB
}

func newDayOffRepository(db *sql.DB) *DayOffRepositoryImp {
	return &DayOffRepositoryImp{
		DB: db,
	}
}

func (repository *DayOffRepositoryImp) FetchDayOffs() ([]DayOff, error) {
	rows, err := repository.DB.Query("select id, day from dayoff")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dayOffs := []DayOff{}
	for rows.Next() {
		var dayOff DayOff
		var dayString string
		if err = rows.Scan(&dayOff.Id, &dayString); err != nil {
			return nil, err
		}

		if day, err := time.Parse(time.RFC3339, dayString); err != nil {
			return nil, err
		} else {
			dayOff.Day = day
			dayOffs = append(dayOffs, dayOff)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return dayOffs, nil
}

func (repository *DayOffRepositoryImp) DeleteDayOff(id int) error {
	tx, err := repository.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`delete from dayoff where id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (repository *DayOffRepositoryImp) CreateDayOff(day time.Time) (int64, error) {
	tx, err := repository.DB.Begin()
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`insert into dayoff(day) values(?);`)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(day.Format(time.RFC3339))
	if err != nil {
		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}
