package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../../models"
	"github.com/gorilla/mux"
)

//Trae del modelo los tipos de cuenta existentes
func GetTiposCuenta(w http.ResponseWriter, r *http.Request) {
	tiposCuenta,_ := models.GetTiposCuenta()
	models.SendData(w, tiposCuenta)
}

//Solicita al modelo un tipo de cuenta
func GetTipoCuenta(w http.ResponseWriter, r *http.Request) {
	if tipoCuenta, err := getTipoCuentaByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		if tipoCuenta.ID == 0 {
			models.SendNotFound(w)
			return
		}
		models.SendData(w, tipoCuenta)
	}
}

//Solicita al modelo la creacion de un tipo de cuenta
func CreateTipoCuenta(w http.ResponseWriter, r *http.Request) {
	var tipoCuenta models.TipoCuenta
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&tipoCuenta); err != nil {
		models.SendUnprocessableEntity(w)
	} else {
		tipoCuenta.Guardar()
		models.SendData(w, tipoCuenta)
	}
}

//UpdateCiudad method
/*func UpdateCiudad(w http.ResponseWriter, r *http.Request) {
	ciudad, err := getCiudadByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	}

	var ciudadResponse models.Ciudad
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&ciudadResponse); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	ciudadResponse.ID = ciudad.ID
	ciudadResponse.Save()
	models.SendData(w, ciudadResponse)
}

//DeleteCiudad method
func DeleteCiudad(w http.ResponseWriter, r *http.Request) {
	if ciudad, err := getCiudadByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		ciudad.Delete()
		models.SendNoContent(w)
	}
}*/

//Solicita al modelo una cuenta pasandole un argumento
func getTipoCuentaByRequest(r *http.Request) (*models.TipoCuenta, error) {
	vars := mux.Vars(r)
	tipoCuentaID, _ := strconv.Atoi(vars["id"])

	tipoCuenta, err := models.GetTipoCuentaByID(tipoCuentaID)
	return tipoCuenta, err
}
