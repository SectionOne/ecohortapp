package main

import (
	"ecohortapp/repository"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Realitzem una funció que retornara un contenidor de Fyne amb el contingut
func (app *Config) registresTab() *fyne.Container {
	//Invoquem la funcio anterior per carregar l'estructura de dedes amb la interficie de slice de slices
	app.Registres = app.getRegistresSlice()
	//També invoquem el mètode getRegistresTable() i l'asignem al item RegistresTable del struct
	app.RegistresTable = app.getRegistresTable()
	//Creem un contenidor amb una capça vertical i a on situem el widget que em general de la taula Registres
	registresContainer := container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		//definirem un contenidor que ens permetra realizar graelles adaptatives i que indicarem amb dos parametres: el nombre de files/columnas i l’objecte que situarem.
		container.NewAdaptiveGrid(1, app.RegistresTable),
	)

	return registresContainer
}

// Realitzem una funcio adicional que ens retornara el punter a la widget en forma de taula i a on situarem les dades
func (app *Config) getRegistresTable() *widget.Table {
	//Definim l'estructura del widget per crear una nova taula amb fyne
	t := widget.NewTable(
		func() (int, int) {
			return len(app.Registres), len(app.Registres[0])
		},
		func() fyne.CanvasObject {
			ctr := container.NewVBox(widget.NewLabel(""))
			return ctr
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Col == (len(app.Registres[0])-1) && i.Row != 0 {
				//Ultima cel.la - situa un botò
				w := widget.NewButtonWithIcon("Borrar", theme.DeleteIcon(), func() {
					//Presentem un dialeg de confirmació
					dialog.ShowConfirm("Borrar?", "", func(deleted bool) {
						if deleted {
							id, _ := strconv.Atoi(app.Registres[i.Row][0].(string)) //Transformem el identificador a decimal sencer
							err := app.DB.BorrarRegistre(int64(id))        //Invoquem el metode per borrar a partir d'un id
							//Capturem possibles errors
							if err != nil {
								app.ErrorLog.Println(err)
							}
						}
						//Forcem el refresc de la taula
						app.actualitzarRegistresTable()
					}, app.MainWindow)
				})
				//Creem un widget d'alta importancia per mostrar un missatge destacat
				w.Importance = widget.HighImportance

				//Definim el contenidor a on situarem el objecte corresponent a el boto.
				o.(*fyne.Container).Objects = []fyne.CanvasObject{
					w,
				}
			} else {
				//situarem la informació rebuda en el slice, recordem que primer gestiona la fila i després la columna
				o.(*fyne.Container).Objects = []fyne.CanvasObject{
					widget.NewLabel(app.Registres[i.Row][i.Col].(string)),
				}
			}
		})

	//Establim el ample de les diferents celdes
	colWidths := []float32{50, 100, 100, 100, 100, 100, 110}
	//Executem una estructura for per aplicar cada un de els amples amb el metode SetColumnWidth
	for i := 0; i < len(colWidths); i++ {
		t.SetColumnWidth(i, colWidths[i])
	}

	return t
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
