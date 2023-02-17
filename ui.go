package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

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

	pronosticTabContent := app.pronosticTab()

	//obtenim les pestanyes de l'aplicació
	tabs := container.NewAppTabs( //Definim un contenidor per les pestanyes i dins afegim cada una de les pestanyes amb icones
		container.NewTabItemWithIcon("Pronòstic", theme.HomeIcon(), pronosticTabContent),
		container.NewTabItemWithIcon("Diari Meteorològic", theme.InfoIcon(), canvas.NewText("El contingut dels valors enregistrats els situarem aqui", nil)),
	)

	//Afegim aquesta instrucció per determinar la posició de les pestanyes
	tabs.SetTabLocation(container.TabLocationTop)

	//afegir el contenidor a la finestra
	finalContent := container.NewVBox(climaDadesContent, toolBar, tabs) //Definim un nou contenidor i que afegirem al canvas general.

	//Invoquem la pàgina principal i fem servir el mètode SetContent per afegir el contenidor
	app.MainWindow.SetContent(finalContent)

	//Realitzem una funcio anonima que sera invocada per una GoRutine en segon pla
	go func() {
		for range time.Tick(time.Second * 30) {
			app.actualitzarClimaDadesContent() //Invoquem la funcio de refrescar els preus
		}
	}()
}

func (app *Config) actualitzarClimaDadesContent() {
	app.InfoLog.Print("refrescant els preus") //Realitzem un log per tenir constancia que s'esta executant la gorutine
	precipitacio, tempMax, tempMin, humitat := app.getClimaText()
	app.ClimaDadesContainer.Objects = []fyne.CanvasObject{precipitacio, tempMax, tempMin, humitat}
	app.ClimaDadesContainer.Refresh()

	grafic := app.obtenirGrafic()
	app.PronosticGraficContainer.Objects = []fyne.CanvasObject{grafic}
	app.PronosticGraficContainer.Refresh()
}
