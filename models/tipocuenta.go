package models

var tipoCuentaSchemeSQLITE string = `CREATE TABLE IF NOT EXISTS tipos_cuenta(
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    tipo_cuenta TEXT NOT NULL
);`

//TipoCuenta clase del modelo TipoCuenta para la creacion de objetos
type TipoCuenta struct {
    ID                 int      `json:"id"`
    NombreTipoCuenta   string   `json:"tipo_cuenta"`
}

type TiposCuenta []TipoCuenta

//NuevoTipoCuenta Funcion que asigna valores a un tipo de cuenta nuevo para ser registrado
func NuevoTipoCuenta(nombreTipoCuente string) *TipoCuenta {
    tipoCuenta := &TipoCuenta{ 
        NombreTipoCuenta:   nombreTipoCuente,
    }
    return tipoCuenta
}
//CrearTipoCuenta metodo para la creacion de un nuevo tipo de cuenta
func CrearTipoCuenta(nombreTipoCuente string) (*TipoCuenta, error) {
    tipoCuenta := NuevoTipoCuenta(nombreTipoCuente)
    err := tipoCuenta.Guardar()
    return tipoCuenta, err
}
//getTipoCuenta Solicita la obtencion de la informacion de un tipo de cuenta
func getTipoCuenta(query string, condicion interface{}) (*TipoCuenta, error) {
    tipoCuenta := &TipoCuenta{}
    rows, err := Query(query, condicion)
    for rows.Next() {
        rows.Scan(&tipoCuenta.ID, &tipoCuenta.NombreTipoCuenta)
    }
    return tipoCuenta, err
}
//GetEmpleadoByID Solicita la obtencion de la informacion de un tipo de cuenta por su ID
func GetTipoCuentaByID(id int) (*TipoCuenta, error) {
    query := "SELECT id, tipo_cuenta FROM tipos_cuenta WHERE id=?"
    return getTipoCuenta(query, id)
}

//GetEmpleados Solicita la obtencion de la lista de tipos de cuentas
func GetTiposCuenta() (TiposCuenta, error) {
    var tiposCuenta TiposCuenta
    query := "SELECT id, tipo_cuenta FROM tipos_cuenta"
    rows, err := Query(query)
    for rows.Next() {
        var tipoCuenta TipoCuenta
        rows.Scan(&tipoCuenta.ID, &tipoCuenta.NombreTipoCuenta)
        tiposCuenta = append(tiposCuenta, tipoCuenta)
    }
    return tiposCuenta, err
}
//Guardar Solicita la transaccion para guardar a un nuevo tipo de cuenta en la base de datos
func (tipoCuenta *TipoCuenta) Guardar() error {
    if tipoCuenta.ID == 0 {
        return tipoCuenta.registrar()
    }
    return tipoCuenta.actualizar()
}
//registrar Registra la transaccion para dar de alta un nuevo tipo de cuenta en la base de datos
func (tipoCuenta *TipoCuenta) registrar() error {
    query := "INSERT INTO tipos_cuenta(tipo_cuenta) VALUES(?);"
    tipoCuentaID, err :=  InsertData(query, tipoCuenta.NombreTipoCuenta)
    tipoCuenta.ID = int(tipoCuentaID)
    return err
}
//actualizar Actualiza la informacion de un tipo de cuenta
func (tipoCuenta *TipoCuenta) actualizar() error {
    query := "UPDATE tipos_cuenta SET tipo_cuenta=? WHERE id=?"
    _, err := Exec(query, tipoCuenta.NombreTipoCuenta, tipoCuenta.ID)
    return err
}
//Eliminar Elimina de forma permanente a un tipo de cuenta de la base de datos
func (tipoCuenta *TipoCuenta) Eliminar() error {
    query := "DELETE FROM tipos_cuenta WHERE id=?"
    _, err := Exec(query, tipoCuenta.ID)
    return err
}
