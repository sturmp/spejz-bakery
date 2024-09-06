package pastry

import "database/sql"

var DB *sql.DB

func fetchAllPastries(languageCode string) (pastries []Pastry, err error) {
	rows, err := DB.Query(`SELECT pastry.id,
	pastrytranslation.name,
	pastrytranslation.description,
	pastry.price,
	unitofmeasuretranslation.name,
	pastry.quantityperpiece,
	pastry.enabled
	FROM pastry
		JOIN pastrytranslation ON pastry.id = pastrytranslation.pastryid
			AND pastrytranslation.language = ?
		JOIN unitofmeasuretranslation ON pastry.unitofmeasure = unitofmeasuretranslation.unitofmeasureid
			AND unitofmeasuretranslation.language = ?`, languageCode, languageCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pastries = []Pastry{}
	for rows.Next() {
		var pastry Pastry
		err = rows.Scan(&pastry.Id,
			&pastry.Name,
			&pastry.Description,
			&pastry.Price,
			&pastry.UnitOfMeasure,
			&pastry.QuantityPerPiece,
			&pastry.Enabled)
		if err != nil {
			return nil, err
		}
		pastries = append(pastries, pastry)
	}

	return pastries, rows.Err()
}

func fetchPastry(pastryId int, languageCode string) (Pastry, error) {
	row := DB.QueryRow(`SELECT pastry.id,
		pastrytranslation.name,
		pastrytranslation.description,
		pastry.price,
		unitofmeasuretranslation.name,
		pastry.quantityperpiece,
		pastry.enabled
		FROM pastry
			JOIN pastrytranslation ON pastry.id = pastrytranslation.pastryid
				AND pastrytranslation.language = ?
			JOIN unitofmeasuretranslation ON pastry.unitofmeasure = unitofmeasuretranslation.unitofmeasureid
				AND unitofmeasuretranslation.language = ?
		WHERE pastry.id = ?`, languageCode, languageCode, pastryId)

	var pastry Pastry

	err := row.Scan(&pastry.Id,
		&pastry.Name,
		&pastry.Description,
		&pastry.Price,
		&pastry.UnitOfMeasure,
		&pastry.QuantityPerPiece,
		&pastry.Enabled)
	if err != nil {
		return pastry, err
	}

	return pastry, nil
}

func updatePastry(pastry Pastry, languageCode string) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE pastry SET price=?, quantityperpiece=?, enabled=? WHERE id=?",
		pastry.Price,
		pastry.QuantityPerPiece,
		pastry.Enabled,
		pastry.Id)
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE pastrytranslation SET name=?, description=? WHERE pastryid=? AND language=?",
		pastry.Name,
		pastry.Description,
		pastry.Id,
		languageCode)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func createPastry(pastry CreatePastryRequest) (id int, err error) {
	tx, err := DB.Begin()
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()

	var pastryId int
	err = tx.QueryRow(`INSERT INTO
		pastry(price, quantityperpiece, unitofmeasure)
		VALUES(?, ?, ?)
		RETURNING id`,
		pastry.Price,
		pastry.QuantityPerPiece,
		pastry.UnitOfMeasure).Scan(&pastryId)
	if err != nil {
		return -1, err
	}

	err = insertPastryLanguages(tx, pastry, pastryId)
	if err != nil {
		return -1, err
	}

	if err := tx.Commit(); err != nil {
		return -1, err
	}

	return pastryId, nil
}

func insertPastryLanguages(tx *sql.Tx, pastry CreatePastryRequest, pastryId int) error {
	rows, err := tx.Query("SELECT DISTINCT language FROM pastrytranslation")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var language string
		if err := rows.Scan(&language); err != nil {
			return err
		}

		_, err = tx.Exec(`INSERT INTO
			pastrytranslation(language, pastryid, name, description)
			VALUES(?, ?, ?, ?)`,
			language,
			pastryId,
			pastry.Name,
			pastry.Description)
		if err != nil {
			return err
		}
	}
	return nil
}
