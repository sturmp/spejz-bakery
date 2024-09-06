package pastry

import "database/sql"

type PastrySqlRepository struct {
	DB *sql.DB
}

func NewPastrySqlRepository(db *sql.DB) *PastrySqlRepository {
	return &PastrySqlRepository{
		DB: db,
	}
}

func (repository *PastrySqlRepository) FetchAllPastries(languageCode string) (pastries []Pastry, err error) {
	rows, err := repository.DB.Query(`SELECT pastry.id,
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

func (repository *PastrySqlRepository) UpdatePastry(pastry Pastry, languageCode string) error {
	tx, err := repository.DB.Begin()
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

func (repository *PastrySqlRepository) CreatePastry(createRequest CreatePastryRequest, languageCode string) (pastry Pastry, err error) {
	tx, err := repository.DB.Begin()
	if err != nil {
		return pastry, err
	}
	defer tx.Rollback()

	var pastryId int
	err = tx.QueryRow(`INSERT INTO
		pastry(price, quantityperpiece, unitofmeasure)
		VALUES(?, ?, ?)
		RETURNING id`,
		createRequest.Price,
		createRequest.QuantityPerPiece,
		createRequest.UnitOfMeasure).Scan(&pastryId)
	if err != nil {
		return pastry, err
	}

	err = repository.insertPastryLanguages(tx, createRequest, pastryId)
	if err != nil {
		return pastry, err
	}

	if err := tx.Commit(); err != nil {
		return pastry, err
	}

	pastry, err = repository.fetchPastry(pastryId, languageCode)
	if err != nil {
		return Pastry{}, err
	}
	return pastry, nil
}

func (repository *PastrySqlRepository) fetchPastry(pastryId int, languageCode string) (Pastry, error) {
	row := repository.DB.QueryRow(`SELECT pastry.id,
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

func (repository *PastrySqlRepository) insertPastryLanguages(tx *sql.Tx, pastry CreatePastryRequest, pastryId int) error {
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
