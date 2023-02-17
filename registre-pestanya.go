package main

import (
	"ecohortapp/repository"
	"fmt"
	"strconv"

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

	//Invoquem el métode inferior registresActuals()
	registres, err := app.registresActuals()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	//Realitzem un append per incloure els registres obtinguts en forma de files i definint alhora les etiquetes de cada columna per la fila inicial.
	slice = append(slice, []interface{}{"ID", "Data", "Precipitació", "Temp. Màxima", "Temp. Minima", "Humitat", "Opcions"})

	//Executem un for per elaborar tantes files com resultats ha obtingut de la BD
	for _, x := range registres {
		//Creem una interficie buida per la fila actual
		var filaActual []interface{}

		//anem afegint a la fila actual cada un dels valors que corresponen a cada columna definida al inici
		filaActual = append(filaActual, strconv.FormatInt(x.ID, 10))           //Transformem el valor numeric a String en base 10
		filaActual = append(filaActual, x.Data.Format("2006-01-02"))           //Formategem la data al standard americà
		filaActual = append(filaActual, fmt.Sprintf("%d%%", x.Precipitacio))   //Formatagem la sortida a un valor decimal enter
		filaActual = append(filaActual, fmt.Sprintf("%d", x.TempMaxima))       //Formatagem la sortida a un valor decimal enter
		filaActual = append(filaActual, fmt.Sprintf("%d", x.TempMinima))       //Formatagem la sortida a un valor decimal enter
		filaActual = append(filaActual, fmt.Sprintf("%d%%", x.Humitat))        //Formatagem la sortida a un valor decimal enter
		filaActual = append(filaActual, widget.NewButton("Borrar", func() {})) //Definim el boto per eliminar i que invocarà una funció que ja definirem

		//Afegim aquesta fila a el slice de files
		slice = append(slice, filaActual)
	}

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
