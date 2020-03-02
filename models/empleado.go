package models

//Empleado clase del modelo empleados para la creacion de objetos
type Empleado struct {
    ID                 int       `json:"id"`
    IDTipoEmpleado     int       `json:"id_tipo_empleado"`
    Nombre             string    `json:"nombre"`
    ApellidoPaterno    string    `json:"apellido_paterno"`
    ApellidoMaterno    string    `json:"apellido_materno"` 
    Username           string    `json:"username"`
    Password           string    `json:"password"`
    habilitado         int
    fechaCreacion      string
}

var empleadosSchemeSQLITE string = `CREATE TABLE IF NOT EXISTS empleados(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
    id_tipo_empleado INTEGER NOT NULL,
	nombre TEXT NOT NULL,
	apellido_paterno TEXT NOT NULL,
	apellido_materno TEXT NOT NULL,
	username TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	habilitado INTEGER NOT NULL,
	fecha_creacion TEXT NOT NULL
);`

type Empleados []Empleado

//NuevoEmpleado Funcion que asigna valores al empleado para ser registrado
func NuevoEmpleado(nombre, apellido_paterno, apellido_materno, username, password string) *Empleado {
    empleado := &Empleado{
        Nombre:	          nombre,
        ApellidoPaterno:  apellido_paterno,
        ApellidoMaterno:  apellido_materno,
        Username:         username,
        Password:         password,
        habilitado:       1,
        fechaCreacion:    ObtenerFechaHoraActualString(),
    }
    return empleado
}

//CrearEmpleado metodo para la creacion de un nuevo empleado
func CrearEmpleado(nombre, apellido_paterno, apellido_materno, username, password string) (*Empleado, error) {
	empleado := NuevoEmpleado(nombre, apellido_paterno, apellido_materno, username, password)
	err := empleado.Save()
	return empleado, err
}

//GetEmpleado Solicita la obtencion de la informacion de un empleado
func GetEmpleado(sql string, condicion interface{}) (*Empleado, error) {
    empleado := &Empleado{}
    rows, err := Query(sql, condicion)
    for rows.Next() {
        rows.Scan(&empleado.ID, &empleado.IDTipoEmpleado ,&empleado.Nombre, &empleado.ApellidoPaterno, &empleado.ApellidoMaterno, 
            &empleado.Username, &empleado.Password, &empleado.habilitado, &empleado.fechaCreacion)
    }
    return empleado, err
}

//GetEmpleadoByID Solicita la obtencion de la informacion de un empleado por su ID
func GetEmpleadoByID(id int) (*Empleado, error) {
    sql := "SELECT id, id_tipo_empleado, nombre, apellido_paterno, apellido_materno, username, password, habilitado, fecha_creacion FROM empleados WHERE habilitado=1 AND id=?"
    return GetEmpleado(sql, id)
}

//GetEmpleadoByUsername Solicita la obtencion de un empleado por su nombre de usuario
func GetEmpleadoByUsername(username string) (*Empleado, error) {
    sql := "SELECT id, id_tipo_empleado, nombre, apellido_paterno, apellido_materno, username, password, habilitado, fecha_creacion FROM empleados WHERE habilitado=1 AND username=?"
    return GetEmpleado(sql, username)
}

//GetEmpleados Solicita la obtencion de la lista de empleados
func GetEmpleados() Empleados {
    var empleados Empleados
    sql := "SELECT id, id_tipo_empleado, nombre, apellido_paterno, apellido_materno, username, password, habilitado, fecha_creacion FROM empleados WHERE habilitado=1"
    rows, _ := Query(sql)
    for rows.Next() {
        var empleado Empleado
        rows.Scan(&empleado.ID, &empleado.IDTipoEmpleado, &empleado.Nombre, &empleado.ApellidoPaterno, 
            &empleado.ApellidoMaterno, &empleado.Username, &empleado.Password, &empleado.habilitado, &empleado.fechaCreacion)
        empleados = append(empleados, empleado)
    }
    return empleados
}

/*func LoginEmpleado(username, password string) (*Empleado, error) {
	empleado,_ := GetEmpleadoByUsername(username)
	if empleado.Password != password {
		return &Empleado{}, errors.New("Usuario o contraseña no coinciden")
	}
	return empleado, nil
}*/

//Save Solicita la transaccion para guardar a un nuevo empleado en la base de datos
func (empleado *Empleado) Save() error {
	if empleado.ID == 0 {
		return empleado.registrar()
	}
	return empleado.actualizar()
}

//registrar Registra la transaccion para dar de alta un nuevo empleado en la base de datos
func (empleado *Empleado) registrar() error {
    empleado.habilitado=1
    empleado.fechaCreacion=ObtenerFechaHoraActualString()
	sql := "INSERT INTO empleados(id_tipo_empleado, nombre, apellido_paterno, apellido_materno, username, password, habilitado, fecha_creacion) VALUES(?,?,?,?,?,?,?,?);"
	empleadoID, err := InsertData(sql, empleado.IDTipoEmpleado , empleado.Nombre, empleado.ApellidoPaterno, 
        empleado.ApellidoMaterno, empleado.Username, empleado.Password, empleado.habilitado, empleado.fechaCreacion)
	empleado.ID = int(empleadoID)
	return err
}

//actualizar Actualiza la informacion de un empleado
func (empleado *Empleado) actualizar() error {
	sql := "UPDATE empleados SET id_tipo_empleado=?, nombre=?, apellido_paterno=?, apellido_materno=?, username=?, password=?, habilitado=? WHERE id=?"
	_, err := Exec(sql, empleado.IDTipoEmpleado, empleado.Nombre, empleado.ApellidoPaterno, empleado.ApellidoMaterno, 
        empleado.Username, empleado.Password, empleado.habilitado, empleado.ID)
	return err
}

//EliminarLog Elimina al empleado de manera logica de la base de datos
func (empleado *Empleado) EliminarLog() error {
    empleado.habilitado=0
    return empleado.actualizar()
}

//Eliminar Elimina de forma permanente a un empleado de la base de datos
func (empleado *Empleado) Eliminar() error {
	sql := "DELETE FROM empleados WHERE id=?"
	_, err := Exec(sql, empleado.ID)
	return err
}

//SetPassword Envia la contraseña
func (empleado *Empleado) SetPassword(password string) {
	empleado.Password = password
    empleado.Save()
}
