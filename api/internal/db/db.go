package dbUtil

import (
	"api/internal/configuration"
	"api/internal/endpoints/bakingschedule"
	"api/internal/endpoints/order"
	"api/internal/endpoints/pastry"
	"database/sql"
	"log"
	"os"
	"time"
)

var pastries []pastry.Pastry = []pastry.Pastry{
	{"Biscuit", "Nem az az édes, ez sós... Vajas pogácsa jó?! Vajas pogácsa!", "3500", "kg", ""},
	{"Foccacia", "Olasz olajos kenyér lángos... Feltét nélkül.", "~440", "db", "~125g"},
	{"kenyér", "Sima kenyér. Semmi extra.", "600", "db", "750g"},
	{"English muffin", "Nem, ez nem az édesség. <a href=\"https://www.google.com/search?client=firefox-b-d&q=english+muffin\">Nézz utána!</a>", "150", "db", "65g"},
	{"Kakaós csiga", "Kakaós és fel van tekerve.", "500", "db", "100g"},
	{"Tortilla", "Mexikói lapos lángos. Kaja origami.", "100", "db", "~20cm"},
	{"Heti különlegesség", "Kísérleti sütések jól sikerült egyedei.", "TBD", "db", ""},
}

var initialOrders []order.Order = []order.Order{
	{1, "kenyér", "Zizi", 1, time.Now().AddDate(0, 0, 1), time.Time{}},
	{2, "kakaós csiga", "Andi", 5, time.Now().AddDate(0, 0, 2), time.Time{}},
	{3, "English muffin", "Roland", 14, time.Now().AddDate(0, 0, 2), time.Time{}},
}

var initialSchedules []bakingschedule.BakingSchedule = []bakingschedule.BakingSchedule{
	{"kakaós csiga", 12, 5, time.Now()},
	{"kenyér", 2, 1, time.Now()},
	{"kenyér", 2, 2, time.Now().AddDate(0, 0, 1)},
	{"Heti különlegesség", 15, 5, time.Now().AddDate(0, 0, 1)},
	{"Biscuit", 0.5, 0.5, time.Now().AddDate(0, 0, 1)},
}

var initialOffDays []time.Time = []time.Time{
	getFirstDayOfWeek(time.Now()),
}

func ConnectToDb() *sql.DB {
	db, err := sql.Open("sqlite3", configuration.AppConfig.Database.Path)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func InitDb() *sql.DB {
	if _, err := os.Stat(configuration.AppConfig.Database.Path); err == nil {
		return nil
	}

	db := ConnectToDb()

	createPastryTableStatement := `
        create table pastry(name text not null primary key,
            description text,
            price integer,
            unitofmeasure text,
            quantityperpiece text);
    `
	_, err := db.Exec(createPastryTableStatement)
	if err != nil {
		log.Printf("%q: %s\n", err, createPastryTableStatement)
	}
	insertPastries(db)

	createOrderTableStatement := `
		create table pastryorder(id integer not null primary key,
			pastry text,
			customer text,
			quantity real,
			preferedDate text,
			scheduledDate text);
	`
	_, err = db.Exec(createOrderTableStatement)
	if err != nil {
		log.Printf("%q: %s\n", err, createOrderTableStatement)
	}
	insertOrders(db)

	createBakingScheduleTableStatement := `
		create table bakingschedule(pastry text not null,
			quantity real,
			reserved real,
			readyDate text,
			PRIMARY KEY(pastry, readyDate),
			FOREIGN KEY(pastry) REFERENCES pastry(name));
	`
	_, err = db.Exec(createBakingScheduleTableStatement)
	if err != nil {
		log.Printf("%q: %s\n", err, createBakingScheduleTableStatement)
	}
	insertBakingSchedules(db)

	createDayOffTableStatement := `
		create table dayoff( day text not null primary key );
	`
	_, err = db.Exec(createDayOffTableStatement)
	if err != nil {
		log.Printf("%q: %s\n", err, createBakingScheduleTableStatement)
	}
	insertDaysOff(db)

	return db
}

func getFirstDayOfWeek(time time.Time) time.Time {
	weekDay := time.Weekday()
	if weekDay == 0 {
		weekDay = 7
	}
	return time.AddDate(0, 0, -int(weekDay)+1)
}

func insertPastries(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare(`insert into
        pastry(name, description, price, unitofmeasure, quantityperpiece)
        values(?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for i := 0; i < len(pastries); i++ {
		_, err := stmt.Exec(pastries[i].Name, pastries[i].Description, pastries[i].Price, pastries[i].UnitOfMeasure, pastries[i].QuantityPerPiece)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func insertOrders(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare(`insert into
		pastryorder(pastry, customer, quantity, preferedDate, scheduledDate)
        values(?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for i := 0; i < len(initialOrders); i++ {
		_, err := stmt.Exec(initialOrders[i].Pastry,
			initialOrders[i].Customer,
			initialOrders[i].Quantity,
			initialOrders[i].PreferedDate.Format(time.RFC3339),
			formatDateOrDefault(initialOrders[i].ScheduledDate))
		if err != nil {
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func formatDateOrDefault(t time.Time) string {
	if t == (time.Time{}) {
		return ""
	}
	return t.Format(time.RFC3339)
}

func insertBakingSchedules(db *sql.DB) {
	tx, err := db.Begin()
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

	for i := 0; i < len(initialSchedules); i++ {
		_, err := stmt.Exec(initialSchedules[i].Pastry, initialSchedules[i].Quantity, initialSchedules[i].Reserved, initialSchedules[i].ReadyDate.Format(time.RFC3339))
		if err != nil {
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func insertDaysOff(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare(`insert into
		dayoff(day)
        values(?)`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for i := 0; i < len(initialOffDays); i++ {
		_, err := stmt.Exec(initialOffDays[i].Format(time.RFC3339))
		if err != nil {
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
