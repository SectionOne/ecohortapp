package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (app *Config) getToolBar(win fyne.Window) *widget.Toolbar {
	toolBar := widget.NewToolbar(
		widget.NewToolbarSpacer(), //Crearem un espaciador que empenyi els diferents items cap a la dreta
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			app.addRegistresDialog()
		}), //Crearem una nova acció indicant quina icona i quina funcio estaràn involucrades
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			app.actualitzarClimaDadesContent()
		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {}),
	)

	return toolBar
}

// Funció per afegir Registres a on referenciem el struct Config
func (app *Config) addRegistresDialog() dialog.Dialog {
	//Definim les variables a on guardarem el resultat del mètode d'entrada
	dataRegistreEntrada := widget.NewEntry()
	precipitacioEntrada := widget.NewEntry()
	tempMaximaEntrada := widget.NewEntry()
	tempMinimaEntrada := widget.NewEntry()
	humitatEntrada := widget.NewEntry()

	app.AfegirRegistresDataRegistreEntrada = dataRegistreEntrada
	app.AfegirRegistresPrecipitacioEntrada = precipitacioEntrada
	app.AfegirRegistresTempMaximaEntrada = tempMaximaEntrada
	app.AfegirRegistresTempMaximaEntrada = tempMinimaEntrada
	app.AfegirRegistresHumitatEntrada = humitatEntrada

	//Definim un placeholder i aixi facilitar l'usabilitat en el camp data de compra
	dataRegistreEntrada.PlaceHolder = "YYYY-MM-DD"

	//Crearem el dialeg creant un formulari
	addForm := dialog.NewForm(
		"Afegir Registre",
		"Afegir",
		"Cancelar",
		//Afegirem les etiquetes en forma de item per el formulari
		[]*widget.FormItem{
			{Text: "Data Registre", Widget: dataRegistreEntrada},
			{Text: "Probabilitat de precipitació", Widget: precipitacioEntrada},
			{Text: "Temperatura màxima", Widget: tempMaximaEntrada},
			{Text: "Temperatura minima", Widget: tempMinimaEntrada},
			{Text: "Humitat", Widget: humitatEntrada},
		},
		//A continuació realitzem la validació de les dades
		func(valid bool) {
			if valid {
				//TODO: Desenvolupar un filtratge
			}
		},
		app.MainWindow)

	//Establim el tamany de la finestra i mostrem el dialeg
	addForm.Resize(fyne.Size{Width: 400})
	addForm.Show()

	return addForm
}
