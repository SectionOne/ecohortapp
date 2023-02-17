package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)
//Fem una funció que retornara quatre elements de texte amb fyne i que es realitzarà una inferencia a l'estructura Config
func (app *Config) getClimaText() (*canvas.Text, *canvas.Text, *canvas.Text, *canvas.Text){
	//Definim variables 
	var g Diaria //Variable que conté total la informació de climatologia de tipus Diaria
	var precipitacio, tempMax, tempMin, humitat *canvas.Text //Quatre variables de tipus canvas.Text

	prediccio, err := g.GetPrediccions()
	if err != nil {
		//Definirem el color del text segons els valors mostrats. Actualment el definirem en gris
		//En el cas que s'hagi generat un error al invocar la funció d'obtenir prediccions.
		gris := color.NRGBA{R: 155, G: 155, B:155, A:255}
		precipitacio = canvas.NewText("Precipitació: No Definit", gris)
		tempMax = canvas.NewText("Temp. Max: No Definit", gris)
		tempMin = canvas.NewText("Temp. Min: No Definit", gris)
		humitat = canvas.NewText("Humitat: No Definit", gris)
	} else {
		displayColor := color.NRGBA{R: 0, G: 180, B: 0, A: 255} //Definim un color per defecte

		//Si la precipitacio és menor a 50% mostrarem un color vermell
		if prediccio.ProbPrecipitacio < 50 {
			displayColor = color.NRGBA{R: 180, G: 0, B: 0, A: 255} //Color per quan el valor és inferior al 50%
		}

		//Preparem els strings de cada tipus de dada cimatologica amb el texte corresponent, emprant el Springf per encapsular tot en una variable de tipus string
		precipitacioTxt := fmt.Sprintf("Precipitació: %d%%", prediccio.ProbPrecipitacio)
		tempMaxTxt := fmt.Sprintf("Temp. Max: %d", prediccio.TemperaturaMax)
		tempMinTxt := fmt.Sprintf("Temp. Min: %d", prediccio.TemperaturaMin)
		humitatTxt := fmt.Sprintf("Humitat: %d%%", prediccio.HumitatRelativa)

		//Ja ara creem els nous elements de texte, indicant com a parametres els textes i els colors corresponents.
		precipitacio = canvas.NewText(precipitacioTxt, displayColor)
		tempMax = canvas.NewText(tempMaxTxt, nil)
		tempMin = canvas.NewText(tempMinTxt, nil)
		humitat = canvas.NewText(humitatTxt, nil)

	}

	//I per completar definirem els alineaments d’aquest elements de tipus texte, definint un a l’esquerra, un altre el centre i l’ultim a la dreta.
	precipitacio.Alignment = fyne.TextAlignLeading
	tempMax.Alignment = fyne.TextAlignCenter
	tempMin.Alignment = fyne.TextAlignCenter
	humitat.Alignment = fyne.TextAlignTrailing

	//Retornem les tres variables
	return precipitacio, tempMax, tempMin, humitat
}