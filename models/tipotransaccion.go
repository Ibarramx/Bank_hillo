package models

var tipoTransaccionSchemeSQLITE string = `CREATE TABLE IF NOT EXISTS tipos_transaccion(
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    tipo_transaccion TEXT NOT NULL
);`

//TipoTransaccion clase del modelo TipoTransaccion para la creacion de objetos
type TipoTransaccion struct {
    ID                      int     `json:"id"`
    NombreTipoTransaccion   string  `json:"tipo_transaccion"`
}

type TiposTransaccion []TipoTransaccion

//NuevoTipoTransaccion Asigna valor al nuevo tipo de transaccion
func NuevoTipoTransaccion (nombreTipoTransaccion string) *TipoTransaccion {
    tipoTransaccion := &TipoTransaccion{
        NombreTipoTransaccion:  nombreTipoTransaccion,
    }
    return tipoTransaccion
}

//CrearTipoTransaccion Valida y solicita el registro de el nuevo tipo de transaccion
func CrearTipoTransaccion (nombreTipoTransaccion string) (*TipoTransaccion, error) {
    tipoTransaccion := NuevoTipoTransaccion(nombreTipoTransaccion)
    err := tipoTransaccion.Guardar()
    return tipoTransaccion, err
}

//GetTipoTransaccion obtiene un tipo de transaccion
func GetTipoTransaccion(query string, condicion interface{}) (*TipoTransaccion, error) {
    tipoTransaccion := &TipoTransaccion{}
    rows, err := Query(query, condicion)
    for rows.Next() {
        rows.Scan(&tipoTransaccion.ID, &tipoTransaccion.NombreTipoTransaccion)
    }
    return tipoTransaccion, err
}

//GetTipoTransaccionByID Obtiene un id de la base de datos por su ID
func GetTipoTransaccionByID (id int) (*TipoTransaccion, error) {
    query := "SELECT id, tipo_transaccion FROM tipos_transaccion WHERE id = ?"
    return GetTipoTransaccion(query, id)
}

//GetTiposTransaccion Obtiene todos los tipos de transacciones que hay de la base de datos
func GetTiposTransaccion() (TiposTransaccion, error) {
    var tiposTransaccion TiposTransaccion
    query := "SELECT id, tipo_transaccion FROM tipos_transaccion"
    rows, err := Query(query)
    if err != nil {
        return nil, err
    }
    for rows.Next() {
        tipoTransaccion := TipoTransaccion{}
        rows.Scan(&tipoTransaccion.ID, &tipoTransaccion.NombreTipoTransaccion )
        tiposTransaccion = append(tiposTransaccion, tipoTransaccion)
    }
    return tiposTransaccion, err
}

//Guardar Valida y solicita guardar un tipo de transaccion nuevo
func (tipoTransaccion *TipoTransaccion) Guardar() error {
    if tipoTransaccion.ID == 0 {
        return tipoTransaccion.registrar()
    }
    return tipoTransaccion.actualizar()
}

//registrar da de alta un nuevo tipo de transaccion en la base de datos
func (tipoTransaccion *TipoTransaccion) registrar() error {
    query := "INSERT INTO tipos_transaccion (tipo_transaccion) VALUES(?)"
    tipoTransaccionID, err := InsertData(query, tipoTransaccion.NombreTipoTransaccion)
    tipoTransaccion.ID = int(tipoTransaccionID)
    return err
}

//actualizar actualiza un registro de la base de datos
func (tipoTransaccion *TipoTransaccion) actualizar() error {
    query := "UPDATE tipos_transaccion SET tipo_transaccion=? WHERE id=?"
    _, err := Exec(query, tipoTransaccion.NombreTipoTransaccion, tipoTransaccion.ID)
    return err
} 

//Eliminar elimina un tipo de transaccion de la base de datos
func (tipoTransaccion *TipoTransaccion) Eliminar() error {
    query := "DELETE FROM tipos_transaccion WHERE id=?"
    _, err := Exec(query, tipoTransaccion.ID)
    return err
}
