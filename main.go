package main

import (
	"database/sql"
	"ecohortapp/repository"
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	_ "github.com/glebarez/go-sqlite"
)

// Crearem un struct amb totes les configuracions que necessiti la nostre App
type Config struct {
	App                      fyne.App              //Definim que emprara Fyne per construir la GUI de l'App
	InfoLog                  *log.Logger           //Definim un Log d'accions
	ErrorLog                 *log.Logger           //Definim un Log d'errors
	DB                       repository.Repository //Definim la referencia a la conexió a SQLite
	MainWindow               fyne.Window           //Aqui enmagatzemem la referencia a certes arees de la ui per controlar les actualitzacions de les mateixes.
	ClimaDadesContainer      *fyne.Container       //Guardem el contenidor de les dades del clima, referenciant el punter de memòria del contenidor de fyne.
	PronosticGraficContainer *fyne.Container       //Definim un camp a on enmagatzem el contenidor del gràfic de clima, que ara sera de tipus contenidor fyne
	Registres                [][]interface{}       //Per emmagatzemar el slice de slices en forma de interfície a on esta contingut les dades obtingudes de la bd
	RegistresTable           *widget.Table         //Per emmagatzemar la referencia al punter que correspon el widget de la Taula.
	HTTPClient               http.Client           //Afegim la referència al client http sence necessitat de invocar la llibreria
	AfegirRegistresDataRegistreEntrada *widget.Entry //Afegim la referencia a la entrada del valor data registre per a nous registres que guardem en la bd
	AfegirRegistresPrecipitacioEntrada *widget.Entry //Afegim la referencia a la entrada del valor precipitacio per a nous registres que guardem en la bd
	AfegirRegistresTempMaximaEntrada *widget.Entry //Afegim la referencia a la entrada del valor tempMaxima per a nous registres que guardem en la bd
	AfegirRegistresTempMinimaEntrada *widget.Entry //Afegim la referencia a la entrada del valor tempMinima per a nous registres que guardem en la bd
	AfegirRegistresHumitatEntrada *widget.Entry //Afegim la referencia a la entrada del valor humitat per a nous registres que guardem en la bd
}

func main() {
	var myApp Config //Creem una variable que sigui de tipus Config i aixi enmagatzemar la configuració de l'App

	// crearem una aplicació fyne
	fyneApp := app.NewWithID("cat.cibernarium.ecohortapp") //El definit el mètode New amb una id ens permet distribuir la nostre app en un MarketPlace
	myApp.App = fyneApp

	//crearem els nostres logs
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)        //Creem un Log per els registres informatius
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Lshortfile) //Creraem un log per els registres d'error

	//conexió amb la base de dades
	sqlDB, err := myApp.connectSQL() //Invoquem la funció d'establiment de la conexió
	if err != nil {
		//Recordem que Panic es l'equivalent a Print pero acompanyat a una crida a Panic
		log.Panic(err)
	}

	//crearem un repositori de base de dades
	myApp.setupDB(sqlDB)

	//Definim la capacitat de que l'usuari modifiqui el municipi i la apiKey
	municipi = fyneApp.Preferences().StringWithFallback("municipi","08001")
	municipi = fyneApp.Preferences().StringWithFallback("apiKey","eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJvZGlnaW9jaW9AZ21haWwuY29tIiwianRpIjoiYjRlZTViMjctZDhhMS00YmIxLWFiZjgtYmFjYTViOTc5ZDhjIiwiaXNzIjoiQUVNRVQiLCJpYXQiOjE2NzU2MTY3OTIsInVzZXJJZCI6ImI0ZWU1YjI3LWQ4YTEtNGJiMS1hYmY4LWJhY2E1Yjk3OWQ4YyIsInJvbGUiOiIifQ.y-WKC8DkAJ4O__aNkvWS60AwmYl6dVHcBZKcowfmNKs")

	//crearem i definim el tamany de una pantalla de fyne
	myApp.MainWindow = fyneApp.NewWindow("Eco Hort App")
	myApp.MainWindow.Resize(fyne.NewSize(900, 700)) //Definim el tamany de la finestra
	myApp.MainWindow.SetFixedSize(true)              //Definim que tindra un tamany fixe
	myApp.MainWindow.SetMaster()                     //Indiquem que es la pantalla principal. Si tanquem aquesta pantalla la aplicacio finalitza

	myApp.makeUI() //Crearem una invocació a una funció externa que creara la interficié grafica a partir del contingut.

	//mostrar i executar l'aplicació
	myApp.MainWindow.ShowAndRun()
}

// Realitzarem una funció per invocar la conexió a la BD
func (app *Config) connectSQL() (*sql.DB, error) {
	path := ""

	//Treballarem amb variables d'entorn per establir les configuracions i en aquest cas comprobem si DB_PATH (que fa referencia a la ruta de la db) te un valor o no.
	if os.Getenv("DB_PATH") != "" {
		//En cas de tenir valor, el recupera
		path = os.Getenv("DB_PATH")
	} else {
		//En cas contrari crearà aquest arxiu de Bases de dades dins de la ruta de l'aplicació
		path = app.App.Storage().RootURI().Path() + "/sql.db"
		//Incloem un registre de control al log
		app.InfoLog.Println("db in:", path)
	}

	//A continuació establim la conexió i controlem possibles errors
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Creem un repositori per a la Base de Dades
func (app *Config) setupDB(sqlDB *sql.DB) {
	app.DB = repository.NewSQLiteRepository(sqlDB)

	err := app.DB.Migrate()
	if err != nil {
		app.ErrorLog.Println(err)
		log.Panic()
	}
}
