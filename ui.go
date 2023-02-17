package main

import "fyne.io/fyne/v2/container"

func (app *Config) makeUI() {
	// obtenir les dades de l'API (Probabilitat de precipitacions, Temperatura Max. i Min. i Humitat)
	precipitacio, tempMax, tempMin, humitat := app.getClimaText()
	//insertar la informació dins del contenidor
	climaDadesContent := container.NewGridWithColumns(4,
		precipitacio,
		tempMax,
		tempMin,
		humitat,
	)
	//afegir el contenidor a la finestra
}
