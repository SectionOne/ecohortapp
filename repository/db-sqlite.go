package repository

import (
	"database/sql"
	"errors"
	"time"
)

type SQLiteRepository struct {
	Conn *sql.DB
}

// Aquesta funció retornara el struct poblat amb la conexió a la bd
func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		Conn: db,
	}
}

// Desenvolupem les funcions que hem mencionat en la interficie
// Hem de indicar en el receptor que farem servir un punter sobre el metode NewSQLiteRepository per emprar la conexió establerta per aquestes accions
func (repo *SQLiteRepository) Migrate() error {
	query := `
	create table if not exists registres(
		id integer primary key autoincrement,
		data_registre integer not null,
		precipitacio integer not null,
		temp_maxima integer not null,
		temp_minima integer not null,
		humitat integer not null)
		`

	_, err := repo.Conn.Exec(query)
	return err
}

func (repo *SQLiteRepository) InsertRegistre(registres Registres) (*Registres, error) {
	//Preparem la instrucció per afegir un registre en la taula registres
	stmt := "insert into registres (data_registre, precipitacio, temp_maxima, temp_minima, humitat) values (?,?,?,?,?)"

	//Executem la instrucció
	res, err := repo.Conn.Exec(stmt, registres.Data.Unix(), registres.Precipitacio, registres.TempMaxima, registres.TempMinima, registres.Humitat)
	if err != nil {
		return nil, err
	}

	//Afegim una crida a la funció LastInsertId() de la rersposta per obtenir la id que s'ha generat amb aquesta inserció
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	registres.ID = id
	return &registres, nil //Preparem els retorn amb un objecte o nil en cas d'error
}

func (repo *SQLiteRepository) ObtenirTotsRegistres() ([]Registres, error) {
	//Formulem la consulta per obtenir totes les columnes de tots els registres ordenats per el camp purchase_date
	query := "select id, data_registre, precipitacio, temp_maxima, temp_minima, humitat from registres order by data_registre"
	//Executem la consulta
	rows, err := repo.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close() //Tanquem la conexió a la bd i optimitzar

	//Creem una variable anomenada tots de tipus slice Registres
	var tots []Registres
	//Executem una estructura for per consultar un dels resultats i inclourels en el slice
	for rows.Next() {
		var h Registres
		var unixTime int64
		err := rows.Scan(
			&h.ID,
			&unixTime,
			&h.Precipitacio,
			&h.TempMaxima,
			&h.TempMinima,
			&h.Humitat,
		)
		//Si es produeix algun error el gestionem
		if err != nil {
			return nil, err
		}
		h.Data = time.Unix(unixTime, 0)
		tots = append(tots, h) //Apliquem la inclusió de l'objecte dins del slice
	}

	return tots, nil //Retornem el slice o nil segons si es genera algun error
}

// Funció per obtenir dades per ID
func (repo *SQLiteRepository) ObtenirRegistrePerID(id int) (*Registres, error) {
	row := repo.Conn.QueryRow("select id, data_registre, precipitacio, temp_maxima, temp_minima, humitat from registres where id = ?", id)

	//Preparem les variables
	var h Registres
	var unixTime int64
	//Preparem un struct de tipus Holdings amb les dades obtingudes
	err := row.Scan(
		&h.ID,
		&unixTime,
		&h.Precipitacio,
		&h.TempMaxima,
		&h.TempMinima,
		&h.Humitat,
	)

	if err != nil {
		return nil, err
	}

	h.Data = time.Unix(unixTime, 0)

	return &h, nil
}

// Realitzem l'actualització d'un Registre.
func (repo *SQLiteRepository) ActualitzarRegistre(id int64, actualitzar Registres) error {
	//Validem si la id és 0 i aixi controlem possibles errors
	if id == 0 {
		return errors.New("La id a actualitzar és incorrecte")
	}

	//Preparem la petició per updatejar les dades
	stmt := "Update registres set data_registre = ?, precipitacio = ?, temp_maxima = ?, temp_minima = ?, humitat = ?"
	res, err := repo.Conn.Exec(stmt, actualitzar.Data.Unix(), actualitzar.Precipitacio, actualitzar.TempMaxima, actualitzar.TempMinima, actualitzar.Humitat, id)
	//Controlem possibles errors
	if err != nil {
		return err
	}

	//Comprobem el nombre de registres que han sigut afectats per l'actualització, per poder controlar si és el cas que son 0.
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errUpdateFailed
	}

	return nil
}

//Funció per borrar un Registre
func (repo *SQLiteRepository) BorrarRegistre(id int64) error {
	res, err := repo.Conn.Exec("delete from registres where id = ?", id)
	if err != nil {
		return err
	}

	//Comprobem quin numnero de registres es veuen afectats per el procés d'eliminació
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errUpdateFailed
	}

	return nil
}