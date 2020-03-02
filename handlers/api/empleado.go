package api

import (
	"encoding/json"
	"net/http"
	//"strconv"

	"../../models"
	//"github.com/gorilla/mux"
)

//Obtiene la lista de empleados del modelo
func GetEmpleados(w http.ResponseWriter, r *http.Request) {
	models.SendData(w, models.GetEmpleados())
}

//Crea un empleado mandandole la informacion al modelo
func CreateEmpleado(w http.ResponseWriter, r *http.Request) {
	var empleado models.Empleado
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&empleado); err != nil {
		models.SendUnprocessableEntity(w)
	} else {
		empleado.Save()
		models.SendData(w, empleado)
	}
}
