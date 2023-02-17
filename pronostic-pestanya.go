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
)

func (app *Config) pronosticTab() *fyne.Container {
	return nil
}

func (app *Config) obtenirGrafic() *canvas.Image {
	return nil
}

func (app *Config) descarregarArxiu(URL, nomArxiu string) error {
	//Obtenim la resposta en bytes desde la crida a una url
	response, err := app.HTTPClient.Get(URL)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return errors.New("Rebem un codi de resposta erronia quan descarreguem la imatge.")
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
