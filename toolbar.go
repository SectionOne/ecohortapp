package main

import (
	"ecohortapp/repository"
	"strconv"
	"time"

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

	validacioData := func(s string) error {
		//Apliquem el format de la data definit en el primer parametre al string indicat en el segon parametre i que correspon el aportat per l'usuari
		//S’ha de referenciar el format amb el layout standard 2006-01-02 i no pas amb cualsevol data
		if _, err := time.Parse("2006-01-02", s); err != nil {
			//Si es produeix algun error en el procés el retornem
			return err
		}
		return nil
	}
	dataRegistreEntrada.Validator = validacioData

	esIntValidador := func(s string) error {
		_, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		return nil
	}
	precipitacioEntrada.Validator = esIntValidador
	tempMaximaEntrada.Validator = esIntValidador
	tempMinimaEntrada.Validator = esIntValidador
	humitatEntrada.Validator = esIntValidador

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
				//Desenvolupar un filtratge, convertint les dades a els formats de la bd
				//S’ha de referenciar el format amb el layout standard 2006-01-02 i no pas amb cualsevol data
				dataRegistre, _ := time.Parse("2006-01-02", dataRegistreEntrada.Text)
				precipitacio, _ := strconv.Atoi(precipitacioEntrada.Text)
				tempMaxima, _ := strconv.Atoi(tempMaximaEntrada.Text)
				tempMinima, _ := strconv.Atoi(tempMinimaEntrada.Text)
				humitat, _ := strconv.Atoi(humitatEntrada.Text)
				
				//Invoquem el mètode de la base de dades per insertar registres i que poblarem amb les dades formatejades
				_, err := app.DB.InsertRegistre(repository.Registres{
					Data:        dataRegistre,
					Precipitacio:  precipitacio,
					TempMaxima: tempMaxima,
					TempMinima: tempMinima,
					Humitat: humitat,
				})
				//Controlem si és produeix algun error
				if err != nil {
					app.ErrorLog.Println(err)
				}
				//Invoquem el paremetre del struct Config per permetre que refresqui el widget de la taula amb el nou registre
				app.actualitzarRegistresTable()
			}
		},
		app.MainWindow)

	//Establim el tamany de la finestra i mostrem el dialeg
	addForm.Resize(fyne.Size{Width: 400})
	addForm.Show()

	return addForm
}
