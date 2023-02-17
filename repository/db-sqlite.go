package repository

import (
	"database/sql"
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