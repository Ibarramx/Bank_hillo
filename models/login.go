package models

var loginSchemeSQLITE string = `CREATE TABLE IF NOT EXISTS logins(
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    id_empleado INTEGER NOT NULL,
    token_string TEXT NOT NULL UNIQUE,
    fecha_registro TEXT NOT NULL,
    sesion_activa INTEGER NOT NULL
);`

//Login clase del modelo Login para la creacion de objetos
type Login struct {
    ID              int     `json:"id"`
    IDEmpleado      int     `json:"id_empleado"`
    TokenString     string  `json:"token_string"`
    FechaRegistro   string  `json:"fecha_registro"`
    SesionActiva    int     `json:"sesion_activa"`
}

//TokenResponse clase del modelo TokenResponse para la creacion de objetos
type TokenResponse struct {
    TokenString     string  `json:"token_string"`
    IDEmpleado      int     `json:"id_empleado"`
    Username        string  `json:"username"`
}

//nuevoLogin Funcion que asigna valores a un nuevo login para ser registrado
func nuevoLogin(idEmpleado int, tokenString string) *Login {
    login := &Login{
        IDEmpleado:     idEmpleado,
        TokenString:    tokenString,
        FechaRegistro:  ObtenerFechaHoraActualString(),
        SesionActiva:   1,
    }
    return login
}

//CrearLogin metodo para la creacion de un nuevo login
func CrearLogin(idEmpleado int) (*Login, error) {
    tokenString, _ := RandomHex(20)
    token := nuevoLogin(idEmpleado, tokenString)
    err := token.Guardar()
    return token, err
}

//getLogin Solicita la obtencion de la informacion de un login para manejar las sesiones
func getLogin(sqlQuery string, condicion interface{}) (*Login, error) {
    login := &Login{}
    rows, err := Query(sqlQuery, condicion)
    for rows.Next() {
        rows.Scan(&login.ID, &login.IDEmpleado, &login.TokenString, &login.FechaRegistro, &login.SesionActiva)
    }
    return login, err
}

//GetLoginByToken Se solicita por token para validar si esta activo
func GetLoginByToken(token string) (*Login, error) {
    sqlQuery := "SELECT id, id_empleado, token_string, fecha_registro, sesion_activa FROM logins WHERE token_string=?;"
    return getLogin(sqlQuery, token)
}

//Guardar Solicita la transaccion para guardar a una nueva sesion en la base de datos
func (login *Login) Guardar() error {
    if login.ID == 0 {
        return login.registrar()
    }
    return login.actualizar()
}
//registrar Registra la sesion para manejar la activdad de un usuario
func (login *Login) registrar() error {
    sqlQuery := "INSERT INTO logins(id_empleado, token_string, fecha_registro, sesion_activa) VALUES(?,?,?,?);"
    loginID, err := InsertData(sqlQuery, login.IDEmpleado, login.TokenString, login.FechaRegistro, login.SesionActiva)
    login.ID = int(loginID)
    return err
}
//actualizar Actualiza las sesiones
func (login *Login) actualizar() error {
    sqlQuery := "UPDATE logins SET id_empleado=?, token_string=?, fecha_registro=?, sesion_activa=? WHERE id=?;"
    _, err := Exec(sqlQuery, login.IDEmpleado, login.TokenString, login.FechaRegistro, login.SesionActiva, login.ID)
    return err
}
//Eliminar Elimina de forma permanente a una sesion de la base de datos
func (login *Login) Eliminar() error {
    sqlQuery := "DELETE FROM logins WHERE id=?;"
    _, err := Exec(sqlQuery, login.ID)
    return err
}