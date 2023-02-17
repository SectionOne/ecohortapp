package repository

import (
	"errors"
	"time"
)

var (
	errUpdateFailed = errors.New("l'actualització a fallat")
	errDeleteFailed = errors.New("el borrat a fallat")
)

type Repository interface {
	//Establim la funció migrate per crear totes les taules que necessitem en la nostre bd
	Migrate() error
	InsertRegistre(h Registres) (*Registres, error)
	//Realitzarem una nova inserció en la interficie per poder obtener tots els resultats que hem enmagatzemat atraves d'un slice
	ObtenirTotsRegistres() ([]Registres, error)
    ObtenirRegistrePerID(id int) (*Registres, error)
    ActualitzarRegistre(id int64, actualitzar Registres) error
    BorrarRegistre(id int64) error

}

// A continuació definim un struct amb els camps i el tipus de dades que emprarem
type Registres struct {
	ID           int64     `json:"id"`
	Data         time.Time `json:"data_registre"`
	Precipitacio int       `json:"precipitacio"`
	TempMaxima   int       `json:"temp_maxima"`
	TempMinima   int       `json:"temp_minima"`
	Humitat      int       `json:"humitat"`
}
