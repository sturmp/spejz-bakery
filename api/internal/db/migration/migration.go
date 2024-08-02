package migration

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func RunMigrations(migrationDir string, db *sql.DB) error {
	log.Println("Running migrations")
	defer log.Println("Finished migrations")

	migrationNames, migrationScripts, err := parseMigrations(migrationDir)
	if err != nil {
		return err
	}
	appliedMigrations, err := fetchAppliedMigrations(db)
	if err != nil {
		return err
	}

	indexOfMigrationsToRun := getIndexesOfMigrationsToRun(migrationNames, appliedMigrations)
	if len(indexOfMigrationsToRun) == 0 {
		log.Println("DB is up-to-date")
		return nil
	}

	for i := 0; i < len(indexOfMigrationsToRun); i++ {
		name := migrationNames[indexOfMigrationsToRun[i]]
		script := migrationScripts[indexOfMigrationsToRun[i]]

		tx, err := db.Begin()
		if err != nil {
			return err
		}

		log.Printf("[%d/%d]: %s\n", i+1, len(indexOfMigrationsToRun), name)
		if err := runMigration(name, script, tx); err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Commit(); err != nil {
			return err
		}
	}

	return nil
}

func parseMigrations(migrationDir string) (migrationNames []string, migrationScripts []string, err error) {
	migrationNames = []string{}
	migrationScripts = []string{}

	files, err := os.ReadDir(migrationDir)
	if err != nil {
		return nil, nil, err
	}

	for i := 0; i < len(files); i++ {
		fileNameParts := strings.Split(files[i].Name(), ".")
		extension := fileNameParts[len(fileNameParts)-1]
		if files[i].Type().IsRegular() && extension == "sql" {
			rawFile, err := os.ReadFile(filepath.Join(migrationDir, files[i].Name()))
			if err != nil {
				return nil, nil, err
			}
			migrationNames = append(migrationNames, files[i].Name())
			migrationScripts = append(migrationScripts, string(rawFile))
		}
	}

	return
}

func fetchAppliedMigrations(db *sql.DB) ([]string, error) {
	appliedMigrations := []string{}
	tableRows, err := db.Query("SELECT name FROM sqlite_schema WHERE type='table' AND name='migration'")
	if err != nil {
		return nil, err
	}
	defer tableRows.Close()

	if tableRows.Next() {
		rows, err := db.Query("SELECT name FROM migration")
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var appliedMigration string
			err = rows.Scan(&appliedMigration)
			if err != nil {
				return nil, err
			}
			appliedMigrations = append(appliedMigrations, appliedMigration)
		}

		rows.Close()
	}

	return appliedMigrations, nil
}

func getIndexesOfMigrationsToRun(migrationNames []string, appliedMigrations []string) []int {
	indexOfMigrationsToRun := []int{}
	for i := 0; i < len(migrationNames); i++ {
		if contains(appliedMigrations, migrationNames[i]) {
			continue
		}
		indexOfMigrationsToRun = append(indexOfMigrationsToRun, i)
	}
	return indexOfMigrationsToRun
}

func runMigration(name string, script string, tx *sql.Tx) error {
	_, err := tx.Exec(script)
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO migration (name, date) VALUES (?, ?)", name, time.Now().Format(time.RFC3339))
	if err != nil {
		return err
	}
	return nil
}

func contains(array []string, element string) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == element {
			return true
		}
	}
	return false
}
