package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/astaxie/beego"
)

func Metodo_get(nombre_servicio, parametro string) ([]byte, error) {
	url := beego.AppConfig.String(nombre_servicio) + parametro
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body, nil
}

func ProcesarJsonArreglos(datos []byte) ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	err := json.Unmarshal(datos, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func Metodo_post(nombre_servicio string, data []byte) ([]byte, error) {
	url := beego.AppConfig.String(nombre_servicio)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body, nil
}

func ProcesarJson(datos []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	err5 := json.Unmarshal(datos, &result)
	if err5 != nil {
		log.Fatal(err5)
		return nil, err5
	}
	return result, nil
}

func Metodo_put(nombre_servicio string, id string, data []byte) ([]byte, error) {
	//obtener la URL base desde la configuracion de Beego
	baseURL := beego.AppConfig.String(nombre_servicio)

	//construir una url final con el ID
	url := fmt.Sprintf("%s/%s", baseURL, id)

	//crear la solicitud put
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	// Establecer el encabezado  COntent-Type
	req.Header.Set("Content-Type", "application/json")

	// Enviar solicitud
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Leer la respuesta
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body, nil
}