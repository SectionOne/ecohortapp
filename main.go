package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	_ "fyne.io/fyne/v2/widget"
)

// Crearem un struct amb totes les configuracions que necessiti la nostre App
type Config struct {
	App      fyne.App    //Definim que emprara Fyne per construir la GUI de l'App
	InfoLog  *log.Logger //Definim un Log d'accions
	ErrorLog *log.Logger //Definim un Log d'errors
}

var myApp Config //Creem una variable que sigui de tipus Config i aixi enmagatzemar la configuració de l'App

func main() {
	// crearem una aplicació fyne
	fyneApp := app.NewWithID("cat.cibernarium.ecohortapp") //El definit el mètode New amb una id ens permet distribuir la nostre app en un MarketPlace
	myApp.App = fyneApp

	//crearem els nostres logs
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)        //Creem un Log per els registres informatius
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Lshortfile) //Creraem un log per els registres d'error

	//conexió amb la base de dades

	//crearem un repositori de base de dades

	//crearem i definim el tamany de una pantalla de fyne
	win := fyneApp.NewWindow("Eco Hort App")

	//mostrar i executar l'aplicació
	win.ShowAndRun()
}
