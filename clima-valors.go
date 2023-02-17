package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	_"time"
)

type PreUrl struct {
	Url string      `json:"datos"` //Definim que el camp Url sera de tipus igual que el item "datos" del Json obtingut
	Client *http.Client //Aquest camp ens servirà per quan realitzem testing, tindrem enmagatzemat aqui el tipus de crida que s'ha realitzat
}

//Desenvolupem un main temporal per provar la correcta recepció de les dades
func main() {
	result, _ := GetPreUrl() //Definim una funció per obtenir la Url per la petició de dades climatologiques
	fmt.Println(result) //Imprimim el resultat Ex. https://opendata.aemet.es/opendata/sh/4aa1d5d1
}

func GetPreUrl() (string, error) {
	//Definim la variable url amb el endpoint corresponent a la Predicció Especifica d'un Municipi. En aquest cas Abrera codi 08001
	url := "https://opendata.aemet.es/opendata/api/prediccion/especifica/municipio/diaria/08001/?api_key=xxx" //S'ha de substituir les 3 x per l'Apikey
	
	//Preparem la petició emprant el package http
	req, _ := http.NewRequest("GET", url, nil) //A on li indiquem el metode get, la url de la petició i el tercer parametyre com a nil

	//Afegim una capçelera per que no cachegi la petició
	req.Header.Add("cache-control", "no-cache")
	//Realitzem la petició emprant el metode Do i transmetent la variable req com a parametre, que conté la petició en si.
	res, err := http.DefaultClient.Do(req)
	//Controlem si és produeix un error i corresponentment err es diferent de nil
	if err != nil {
		log.Println("error contactant amb aemet.es", err)
		return "", err //Retornem un error controlat
	}

	defer res.Body.Close() //Diferim la resposta
	body, err := ioutil.ReadAll(res.Body) //Llegim el cos de la resposta de la peticio
	if err != nil {
		log.Println("error llegint el json", err)
		return "", err //Retornem un error controlat
	}

	preUrl := PreUrl{} //Creem un estruct buit
	err = json.Unmarshal(body, &preUrl)     //Enmagatzema el valor del body en l'estruct PreUrl
	if err != nil {
		log.Println("error unmarshalling", err)
		return "", err
	}
	
	return preUrl.Url, err //Retornem els valors
}