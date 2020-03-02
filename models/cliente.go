package models

//Sentencia que crea tabla de clientes en caso de no existir en la base de datos
var clienteSchemeSQLITE string = `CREATE TABLE IF NOT EXISTS clientes(
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    nombre TEXT NOT NULL,
    apellido_paterno TEXT,
    apellido_materno TEXT,
    clave TEXT UNIQUE,
    habilitado INTEGER NOT NULL,
    fecha_creacion TEXT NOT NULL
);`

//Cliente clase del modelo clientes para la creacion de objetos
type Cliente struct {
    ID              int     `json:"id"`
    Nombre          string  `json:"nombre"`
    ApellidoPaterno string  `json:"apellido_paterno"`
    ApellidoMaterno string  `json:"apellido_materno"`
    Clave           string  `json:"clave"`
    habilitado      int
    fechaCreacion   string
}

type Clientes []Cliente

//NuevoCliente Funcion para la creacion de un nuevo objeto cliente para ser registrado
func NuevoCliente(nombre, apellido_paterno, apellido_materno, clave string) *Cliente {
    cliente := &Cliente{
        Nombre:             nombre,
        ApellidoPaterno:    apellido_paterno,
        ApellidoMaterno:    apellido_materno,
        Clave:              clave,
        habilitado:         1,
        fechaCreacion:      ObtenerFechaHoraActualString(),
    }
    return cliente
}

//CrearCliente Funcion para realizar la transaccion en la base de datos para la creacion de un cliente
func CrearCliente(nombre, apellido_paterno, apellido_materno, clave string) (*Cliente, error) {
    cliente := NuevoCliente(nombre, apellido_paterno, apellido_materno, clave)
    err := cliente.Guardar()
    return cliente, err
}

//GetClientes Transaccion en la base de datos que devuelve todos los clientes de la base de datos que esten habiles
func GetClientes() (Clientes, error) {
    var clientes Clientes
    query := "SELECT id, nombre, apellido_paterno, apellido_materno, clave, habilitado, fecha_creacion FROM clientes WHERE habilitado=1"
    rows, err := Query(query)
    for rows.Next(){
        cliente := Cliente{}
        rows.Scan(&cliente.ID, &cliente.Nombre, &cliente.ApellidoPaterno, &cliente.ApellidoMaterno, &cliente.Clave,
            &cliente.habilitado, &cliente.fechaCreacion)
        clientes = append(clientes, cliente)
    }
    return clientes, err
}

//getCliente Envia la coleccion de clientes obtenidos de la transaccion al controlador
func getCliente(query string, condicion interface{}) (*Cliente, error) {
    cliente := Cliente{}
    rows, err := Query(query, condicion)
    for rows.Next() {
        rows.Scan(&cliente.ID, &cliente.Nombre, &cliente.ApellidoPaterno, &cliente.ApellidoMaterno, &cliente.Clave,
            &cliente.habilitado, &cliente.fechaCreacion)
    }
    return &cliente, err
}

//GetClienteByID Realiza una transaccion en la base de datos solicitando un cliente por su ID
func GetClienteByID(id int) (*Cliente, error) {
    query := "SELECT id, nombre, apellido_paterno, apellido_materno, clave, habilitado, fecha_creacion FROM clientes WHERE habilitado=1 AND id=?"
    return getCliente(query, id)
}

//GetClienteByClave Realiza una transaccion en la base de datos solicitando un cliente por su clave
func GetClienteByClave(clave string) (*Cliente, error) {
    query := "SELECT id, nombre, apellido_paterno, apellido_materno, clave, habilitado, fecha_creacion FROM clientes WHERE habilitado=1 AND clave=?"
    return getCliente(query, clave)
}

//GetClienteByNumeroTarjeta Realiza una transaccion para obtener los datos de un cliente por su numero de tarjeta
func GetClienteByNumeroTarjeta(tarjeta string) (*Cliente, error) {
    query := "SELECT id, nombre, apellido_paterno, apellido_materno, clave, habilitado, fecha_creacion FROM clientes WHERE habilitado=1 AND id in (SELECT id_cliente FROM tarjetas WHERE numero_tarjeta=?)"
    return getCliente(query, tarjeta)
}

//Guardar Funcion para guardar cambios en un cliente o registrarlos
func (cliente *Cliente) Guardar() error {
    if cliente.ID == 0 {
        return cliente.registrar()
    }
    return cliente.actualizar()
}

//registrar Realiza la transaccion en la base de datos para registrar un cliente nuevo
func (cliente *Cliente) registrar() error {
    cliente.habilitado=1
    cliente.fechaCreacion=ObtenerFechaHoraActualString()
    query := "INSERT INTO clientes(nombre, apellido_paterno, apellido_materno, clave, habilitado, fecha_creacion) VALUES(?,?,?,?,?,?);"
    clienteID, err := InsertData(query, cliente.Nombre, cliente.ApellidoPaterno, cliente.ApellidoMaterno, cliente.Clave,
        cliente.habilitado, cliente.fechaCreacion)
    cliente.ID = int(clienteID)
    return err
}

//actualizar Realiza una transaccion en la base de datos para actualizar la informacion de un cliente
func (cliente *Cliente) actualizar() error {
    query := "UPDATE clientes SET nombre=?, apellido_paterno=?, apellido_materno=?, clave=?, habilitado=? WHERE id=?"
    _, err := Exec(query, cliente.Nombre, cliente.ApellidoPaterno, cliente.ApellidoMaterno, cliente.habilitado, cliente.ID )
    return err
}

//EliminarLog Transaccion para eliminar a un cliente de manera logica de la base de datos
func (cliente *Cliente) EliminarLog() error {
    cliente.habilitado=0
    return cliente.actualizar()
}

//Eliminar Transaccion para eliminar de forma definitiva un cliente de la base de datos
func (cliente *Cliente) Eliminar() error {
    query := "DELETE FROM clientes WHERE id=?"
    _, err := Exec(query, cliente.ID)
    return err
}
