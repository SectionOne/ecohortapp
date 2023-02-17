package main

import (
	"log"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

//Crearem un struct amb totes les configuracions que necessiti la nostre App
type Config struct {
	App fyne.App //Definim que emprara Fyne per construir la GUI de l'App
	InfoLog *log.Logger //Definim un Log d'accions
	ErrorLog *log.Logger //Definim un Log d'errors
}

var myApp Config //Creem una variable que sigui de tipus Config i aixi enmagatzemar la configuració de l'App

func main() {
	// crearem una aplicació fyne

	//crearem els nostres logs

	//conexió amb la base de dades

	//crearem un repositori de base de dades

	//crearem i definim el tamany de una pantalla de fyne

	//mostrar i executar l'aplicació
}
