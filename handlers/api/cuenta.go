package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../../models"
	"github.com/gorilla/mux"
	"log"
)

//Solicita al modelo las cuentas para mostrarlas en las vistas
func GetCuentas(w http.ResponseWriter, r *http.Request) {
	cuentas,_ := models.GetCuentas()
	models.SendData(w, cuentas)
}

//Solicita una cuenta especifica al modelo
func GetCuenta(w http.ResponseWriter, r *http.Request) {
	if cuenta, err := getCuentaByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		if cuenta.ID == 0 {
			models.SendNotFound(w)
			return
		}
		models.SendData(w, cuenta)
	}
}

//Envia al modelo la informacion para la creacion de una cuenta
func CreateCuenta(w http.ResponseWriter, r *http.Request) {
	var cuenta models.Cuenta
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&cuenta); err != nil {
		models.SendUnprocessableEntity(w)
	} else {
		cuenta.Guardar()
		models.SendData(w, cuenta)
	}
}

//Envia la informacion al modelo para actualizar una cuenta
func UpdateCuenta (w http.ResponseWriter, r *http.Request) {
	cuenta, err := getCuentaByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	}

	var cuentaResponse models.Cuenta
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&cuentaResponse); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	cuentaResponse.ID = cuenta.ID
	cuentaResponse.SetFechaCreacion(cuenta.GetFechaCreacion())
	if err := cuentaResponse.ActualizarCuenta(); err != nil {
		log.Println(err)
	}

	models.SendData(w, cuentaResponse)
}

//Solicita al modelo una cuenta por repuesta de una peticion con argumentos
func getCuentaByRequest(r *http.Request) (*models.Cuenta, error) {
	vars := mux.Vars(r)
	cuentaID, _ := strconv.Atoi(vars["id"])

	cuenta, err := models.GetCuentaByID(cuentaID)
	return cuenta, err
}
