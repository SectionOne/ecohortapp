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
	) //Definim un contenidor amb una graella amb quatre columnes

	app.ClimaDadesContainer = climaDadesContent
	
	//obtenim la barra d'eines o toolbar
	toolBar := app.getToolBar(app.MainWindow)

	//afegir el contenidor a la finestra
	finalContent := container.NewVBox(climaDadesContent, toolBar) //Definim un nou contenidor i que afegirem al canvas general.
	
	//Invoquem la pàgina principal i fem servir el mètode SetContent per afegir el contenidor
	app.MainWindow.SetContent(finalContent)
}
