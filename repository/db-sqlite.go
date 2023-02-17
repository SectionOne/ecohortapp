package repository

import "database/sql"

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
