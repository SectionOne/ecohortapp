package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func (app *Config) pronosticTab() *fyne.Container {
	grafic := app.obtenirGrafic()                //Invoquem el mètode obtenirGrafic() i el guardem en la variable grafic
	graficContainer := container.NewVBox(grafic) //Creem un nou contenidor Vertical a on afegim la variable grafic
	app.PronosticGraficContainer = graficContainer

	//Retornem el contenidor
	return graficContainer
}

func (app *Config) obtenirGrafic() *canvas.Image {
	//Definim la url de carrega del gràfic
	apiURL := fmt.Sprint("https://my.meteoblue.com/visimage/meteogram_web_hd?look=KILOMETER_PER_HOUR%2CCELSIUS%2CMILLIMETER&apikey=5838a18e295d&temperature=C&windspeed=kmh&precipitationamount=mm&winddirection=3char&city=Abrera&iso2=es&lat=41.5168&lon=1.901&asl=111&tz=Europe%2FMadrid&lang=es&sig=b353aab637f77ab97ae54cbd760554f2")
	var img *canvas.Image

	//Descarreguem la imatge mitjantçant la invocació del mètode descarregarArxiu() transmetent els seus dos parametres la url i el nom del arxiu que desitgem adjudicar-li.
	err := app.descarregarArxiu(apiURL, "pronostic.png")
	//Comprobem si és produeix algun error com per exemple, el servidor esta desconectat, hi ha algun error de xarxa, etç...
	if err != nil {
		//En aquest cas emprarem la imatge del paquet per defecte
		img = canvas.NewImageFromResource(resourceNodisponiblePng)
	} else {
		//Generem la imatge
		img = canvas.NewImageFromFile("pronostic.png")
	}

	//Ara establim el tamany minim de la imatge emprant un mètode de Fyne
	img.SetMinSize(fyne.Size{
		Width:  770,
		Height: 480,
	})

	//Determinem com la imatge omplira el canvas
	img.FillMode = canvas.ImageFillOriginal

	//Retornem la imatge
	return img
}

func (app *Config) descarregarArxiu(URL string, nomArxiu string) error {
	//Obtenim la resposta en bytes desde la crida a una url
	response, err := app.HTTPClient.Get(URL)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return errors.New("rebem un codi de resposta erronia quan descarreguem la imatge")
	}

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	//Decodifiquem la imatge en bytes per poder tractarla
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return err
	}

	//Obtenim la sortida de l'arxiu
	out, err := os.Create(fmt.Sprintf("./%s", nomArxiu))
	if err != nil {
		return err
	}

	//Codifiquem a png transmetent els parametres de la ruta a on es creara l'arxiu i el contingut del binari
	err = png.Encode(out, img)
	if err != nil {
		return err
	}

	return nil
}
