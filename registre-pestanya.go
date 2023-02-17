package main

import (
	"ecohortapp/repository"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Realitzem una funció que retornara un contenidor de Fyne amb el contingut
func (app *Config) registresTab() *fyne.Container {
	return nil
}

// Realitzem una funcio adicional que ens retornara el punter a la widget en forma de taula i a on situarem les dades
func (app *Config) getRegistresTable() *widget.Table {
	return nil
}

// Realitzem una funció per obtenir tots els Registres en un Slice de Slices através d'una interficie que ens sera retornada
func (app *Config) getRegistresSlice() [][]interface{} {
	var slice [][]interface{}

	return slice
}

// Realitzem una altre funció per obtenir tots els Registres amb un slice pero del nostre repositori en la DB
func (app *Config) registresActuals() ([]repository.Registres, error) {
	registres, err := app.DB.ObtenirTotsRegistres()
	if err != nil {
		//Capturem el possible error en el log d'errors
		app.ErrorLog.Println(err)
		return nil, err
	}

	return registres, nil
}
