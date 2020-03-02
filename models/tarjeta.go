package models

var tarjetaSchemeSQLITE string = `CREATE TABLE IF NOT EXISTS tarjetas(
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    id_cuenta INTEGER NOT NULL,
    id_cliente INTEGER NOT NULL,
    numero_tarjeta TEXT NOT NULL UNIQUE,
    nip TEXT NOT NULL,
    fecha_vencimiento TEXT,
    numero_seguridad TEXT,
    habilitado INTEGER NOT NULL,
    fecha_creacion TEXT NOT NULL
);`

//Cliente clase del modelo clientes para la creacion de objetos
type Tarjeta struct {
    ID                  int     `json:"id"`
    IDCuenta            int     `json:"id_cuenta"`
    IDCliente           int     `json:"id_cliente"`
    NumeroTarjeta       string  `json:"numero_tarjeta"`
    Nip                 string  `json:"nip"`
    FechaVencimiento    string  `json:"fecha_vencimiento"`
    NumeroSeguridad     string  `json:"numero_seguridad"`
    habilitado          int
    fechaCreacion       string
}

type Tarjetas []Tarjeta
//NuevoEmpleado Funcion que asigna valores al empleado para ser registrado
func NuevaTarjeta(idCuenta, idCliente int, numeroTarjeta, nip, fechaVenvicimiento, 
    numeroSeguridad string) *Tarjeta {
    tarjeta := &Tarjeta{
        IDCuenta:           idCuenta,
        IDCliente:          idCliente,
        NumeroTarjeta:      numeroTarjeta,
        Nip:                nip,
        FechaVencimiento:   fechaVenvicimiento,
        NumeroSeguridad:    numeroSeguridad,
        habilitado:         1,
        fechaCreacion:      ObtenerFechaHoraActualString(),
    }
    return tarjeta
}

//CrearTarjeta metodo para la creacion de una nueva tarjeta
func CrearTarjeta(idCuenta, idCliente int, numeroTarjeta, nip, fechaVenvicimiento, 
    numeroSeguridad string) (*Tarjeta, error) {
    tarjeta := NuevaTarjeta(idCuenta, idCliente, numeroTarjeta, nip, fechaVenvicimiento, numeroSeguridad)
    err := tarjeta.Guardar()
    return tarjeta, err
}

//getTarjeta Solicita la obtencion de la informacion de una tarjeta
func getTarjeta(query string, condicion interface{}) (*Tarjeta, error) {
    t := &Tarjeta{}
    rows, err := Query(query, condicion)
    for rows.Next() {
        rows.Scan(&t.ID, &t.IDCuenta, &t.IDCliente, &t.NumeroTarjeta, &t.Nip, &t.FechaVencimiento, 
            &t.NumeroSeguridad, &t.habilitado, &t.fechaCreacion)
    }
    return t, err
}
//GetTarjetaByID Solicita la obtencion de la informacion de una tarjeta por su ID
func GetTarjetaByID(id int) (*Tarjeta, error) {
    query := "SELECT id, id_cuenta, id_cliente, numero_tarjeta, nip, fecha_vencimiento, numero_seguridad, habilitado, fecha_creacion FROM tarjetas WHERE id=?"
    return getTarjeta(query, id)
}

//GetTarjetaByNumeroTarjeta Solicita la obtencion de la informacion de una tarjeta por su numero de tarjeta
func GetTarjetaByNumeroTarjeta(numeroTarjeta string) (*Tarjeta, error) {
    query := "SELECT id, id_cuenta, id_cliente, numero_tarjeta, nip, fecha_vencimiento, numero_seguridad, habilitado, fecha_creacion FROM tarjetas WHERE numero_tarjeta=?"
    return getTarjeta(query, numeroTarjeta)
}

//GetTarjetas Solicita la obtencion de la lista de tarjetas activas
func GetTarjetas() (Tarjetas, error){
    var tarjetas Tarjetas
    query := "SELECT id, id_cuenta, id_cliente, numero_tarjeta, nip, fecha_vencimiento, numero_seguridad, habilitado, fecha_creacion FROM tarjetas"
    rows, err := Query(query)
    for rows.Next() {
        var t Tarjeta
        rows.Scan(&t.ID, &t.IDCuenta, &t.IDCliente, &t.NumeroTarjeta, &t.Nip, &t.FechaVencimiento, 
            &t.NumeroSeguridad, &t.habilitado, &t.fechaCreacion)
        tarjetas = append(tarjetas, t)
    }
    return tarjetas, err
}

//GetTarjetasByIDCuenta Solicita la obtencion de la informacion de una tarjeta por su numero de cuenta
func GetTarjetasByIDCuenta(idCuenta int) (Tarjetas, error) {
    var tarjetas Tarjetas
    query := "SELECT id, id_cuenta, id_cliente, numero_tarjeta, nip, fecha_vencimiento, numero_seguridad, habilitado, fecha_creacion FROM tarjetas WHERE id_cuenta=?"
    rows, err := Query(query, idCuenta)
    for rows.Next() {
        var t Tarjeta
        rows.Scan(&t.ID, &t.IDCuenta, &t.IDCliente, &t.NumeroTarjeta, &t.Nip, &t.FechaVencimiento, 
            &t.NumeroSeguridad, &t.habilitado, &t.fechaCreacion)
        tarjetas = append(tarjetas, t)
    }
    return tarjetas, err
}

//ValidTarjeta Valida la informacion de la tarjeta para la realizacion de transacciones
func ValidTarjeta(numeroTarjeta, fechaVencimiento, cvv string) bool {
    tarjeta, _ := GetTarjetaByNumeroTarjeta(numeroTarjeta)

    if tarjeta.ID != 0 {
        if tarjeta.FechaVencimiento == fechaVencimiento && tarjeta.NumeroSeguridad == cvv {
            return true
        }
    }
    return false
}

//Guardar Valida el registro de la tarjeta para retornar los cambios
func (tarjeta *Tarjeta) Guardar() error {
    if tarjeta.ID == 0 {
        return tarjeta.registrar()
    }
    return tarjeta.actualizar()
}

//registrar Metodo que registra en la base de datos las tarjetas asignandole un cliente
func (tarjeta *Tarjeta) registrar () error {
    if tarjeta.NumeroTarjeta == "" || len(tarjeta.NumeroTarjeta) != 16{
        tarjeta.NumeroTarjeta="5050"+RandomDigits(12)
    }

    tarjeta.Nip=RandomDigits(4)

    if tarjeta.FechaVencimiento == "" {
        tarjeta.FechaVencimiento=GetFechaVencimientoString()
    }

    if tarjeta.NumeroSeguridad == "" {
        tarjeta.NumeroSeguridad=RandomDigits(3)
    }
    
    tarjeta.habilitado=1
    tarjeta.fechaCreacion=ObtenerFechaHoraActualString()
    query := "INSERT INTO tarjetas(id_cuenta, id_cliente, numero_tarjeta, nip, fecha_vencimiento, numero_seguridad, habilitado, fecha_creacion) VALUES(?,?,?,?,?,?,?,?);"
    tarjetaID, err := InsertData(query, tarjeta.IDCuenta, tarjeta.IDCliente, tarjeta.NumeroTarjeta, 
        tarjeta.Nip, tarjeta.FechaVencimiento, tarjeta.NumeroSeguridad, tarjeta.habilitado, tarjeta.fechaCreacion)
    tarjeta.ID = int(tarjetaID)
    return err
}

//actualizar actualiza la tarjeta del cliente 
func (tarjeta *Tarjeta) actualizar() error {
    tarjeta.habilitado = 1
    query := "UPDATE tarjetas SET id_cuenta=?, id_cliente=?, numero_tarjeta=?, nip=?, fecha_vencimiento=?, numero_seguridad=?, habilitado=? WHERE id=?"
    _, err := Exec(query, tarjeta.IDCuenta, tarjeta.IDCliente, tarjeta.NumeroTarjeta, tarjeta.Nip,
        tarjeta.FechaVencimiento, tarjeta.NumeroSeguridad, tarjeta.habilitado, tarjeta.ID)
    return err
}

//Eliminar Elimina la tarjeta deslindandola del cliente
func (tarjeta *Tarjeta) Eliminar() error {
    query := "DELETE FROM tarjetas WHERE id=?"
    _, err := Exec(query, tarjeta.ID)
    return err
}

//GetFechaCreacion Obtiene la fecha para la validacion de la tarjeta
func (tarjeta *Tarjeta) GetFechaCreacion() string {
    return tarjeta.fechaCreacion
}

//SetFechaCreacion Guarda la fecha de creacion de la tarjeta
func (tarjeta *Tarjeta) SetFechaCreacion(fecha string) {
    tarjeta.fechaCreacion = fecha
}