package routes

import (
	"../handlers/api"
	"../handlers/app"
	"github.com/gorilla/mux"
)

//Endpoints Maneja los Endpoints para las rutas de la api
func Endpoints(mux *mux.Router) {
	clientesEndpoints(mux)
	tiposCuentaEndpoints(mux)
	cuentasEndpoints(mux)
	tarjetasEndpoints(mux)
	empleadosEndpoints(mux)
	tiposTransaccionesEndpoints(mux)
	transaccionEndpoints(mux)

	application(mux)
}

//clientesEndPoint Maneja las rutas de cliente segun su metodo
func clientesEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/clientes/", api.GetClientes).Methods("GET")
	mux.HandleFunc("/api/clientes/{id:[0-9]+}", api.GetCliente).Methods("GET")
	mux.HandleFunc("/api/clientes/", api.CreateCliente).Methods("POST")
}

//cuentasEndPoint Maneja las rutas de cuentas segun su metodo
func cuentasEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/cuentas/", api.GetCuentas).Methods("GET")
	mux.HandleFunc("/api/cuentas/{id:[0-9]+}", api.GetCuenta).Methods("GET")
	mux.HandleFunc("/api/cuentas/", api.CreateCuenta).Methods("POST")
	mux.HandleFunc("/api/cuentas/{id:[0-9]+}", api.UpdateCuenta).Methods("PUT")
}

//tarjetasEndPoint Maneja las rutas de tarjetas segun su metodo
func tarjetasEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/tarjetas/", api.GetTarjetas).Methods("GET")
	mux.HandleFunc("/api/tarjetas/{id:[0-9]+}", api.GetTarjeta).Methods("GET")
	mux.HandleFunc("/api/tarjetas/", api.CreateTarjeta).Methods("POST")
	mux.HandleFunc("/api/tarjetas/{id:[0-9]+}", api.UpdateTarjeta).Methods("PUT")
}

//tiposCuentaEndpoints Maneja las rutas de tipos de cuenta segun su metodo
func tiposCuentaEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/tipos_cuenta/", api.GetTiposCuenta).Methods("GET")
	mux.HandleFunc("/api/tipos_cuenta/{id:[0-9]+}", api.GetTipoCuenta).Methods("GET")
	mux.HandleFunc("/api/tipos_cuenta/", api.CreateTipoCuenta).Methods("POST")
}

//tiposCuentaEndpoints Maneja las rutas de los tupos de transacciones segun su metodo
func tiposTransaccionesEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/tipos_transaccion/", api.GetTiposTransaccion).Methods("GET")
	mux.HandleFunc("/api/tipos_transaccion/{id:[0-9]+}", api.GetTipoTransaccion).Methods("GET")
	mux.HandleFunc("/api/tipos_transaccion/", api.CreateTipoTransaccion).Methods("POST")
}

//empleadosEndpoints Maneja las rutas de emmpleados segun su metodo
func empleadosEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/empleados/", api.GetEmpleados).Methods("GET")
	//mux.HandleFunc("/api/v1/status/{id:[0-9]+}", handlers.GetStatus).Methods("GET")
	mux.HandleFunc("/api/empleados/", api.CreateEmpleado).Methods("POST")
	//mux.HandleFunc("/api/v1/status/{id:[0-9]+}", handlers.UpdateStatus).Methods("PUT")
	//mux.HandleFunc("/api/v1/status/{id:[0-9]+}", handlers.DeleteStatus).Methods("DELETE")
}

//transaccionEndpoints Maneja las rutas de las transacciones segun su metodo
func transaccionEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/transacciones/", api.GetTransacciones).Methods("GET")
	mux.HandleFunc("/api/transacciones/transferencias/", api.DoTransferencia).Methods("POST")
	mux.HandleFunc("/api/transacciones/depositos/", api.DoDeposito).Methods("POST")
}

//application maneja las rutas Endpoint para cargar las vistas
func application(mux *mux.Router) {
	mux.HandleFunc("/", app.Index)
	mux.HandleFunc("/login/", app.Login)
	mux.HandleFunc("/cliente/", app.Cliente)
	mux.HandleFunc("/cajero/", app.Cajero)
	mux.HandleFunc("/admin/", app.Admin)
}
