package models

import "errors"

//Cuenta clase del modelo ceunta para la creacion de nuevas cuentas
type Cuenta struct {
    ID                 int     `json:"id"`
    NumeroDeCuenta     string  `json:"numero_de_cuenta"`
    Saldo              float32 `json:"saldo"`
    IDCliente          int     `json:"id_cliente"`
    IDTipoDeCuenta     int     `json:"id_tipo_de_cuenta"`
    habilitado         int
    fechaCreacion      string
}

type Cuentas []Cuenta

var cuentaSchemeSQLITE string = `CREATE TABLE IF NOT EXISTS cuentas(
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    numero_de_cuenta TEXT NOT NULL UNIQUE,
    saldo REAL DEFAULT 0.0,
    id_cliente TEXT,
    id_tipo_de_cuenta INTEGER,
    habilitado INTEGER DEFAULT 0,
    fecha_creacion TEXT);`


//nuevaCuenta funcion para dar valores al objeto cuenta para que sea creado
func nuevaCuenta(numeroDeCuenta string, idCliente, idTipoDeCuenta int) *Cuenta {
    cuenta := &Cuenta {
        NumeroDeCuenta: numeroDeCuenta,
        Saldo:          0.0,
        IDCliente:      idCliente,
        IDTipoDeCuenta: idTipoDeCuenta,
        habilitado:     0,
        fechaCreacion:  ObtenerFechaHoraActualString(),
    }
    return cuenta
}

//AltaCuenta solicita la creacion del producto a un metodo
func AltaCuenta(numeroDeCuenta string, idCliente, idTipoDeCuenta int) (*Cuenta, error) {
    cuenta := nuevaCuenta(numeroDeCuenta, idCliente, idTipoDeCuenta )
    err := cuenta.Guardar()
    return cuenta, err
}

//getCuenta Solicita una cuenta al arreglo para mandarla a la vista
func getCuenta(sqlQuery string, condicion interface{}) (*Cuenta, error) {
	cuenta := &Cuenta{}
	rows, err := Query(sqlQuery, condicion)
	for rows.Next() {
		rows.Scan(&cuenta.ID, &cuenta.NumeroDeCuenta, &cuenta.Saldo, &cuenta.IDCliente, &cuenta.IDTipoDeCuenta, 
            &cuenta.habilitado, &cuenta.fechaCreacion)
	}
	return cuenta, err
}

//GetCuentaByID Solicita a la base de datos una cuenta dando de referencia el ID
func GetCuentaByID(id int) (*Cuenta, error) {
	query := "SELECT id, numero_de_cuenta, saldo, id_cliente, id_tipo_de_cuenta, habilitado, fecha_creacion FROM cuentas WHERE id=?"
	return getCuenta(query, id)
}

//GetCuentaByNumeroCuenta Solicita a la base de datos una cuenta dando de 
//referencia el numero de cuenta
func GetCuentaByNumeroCuenta(numeroDeCuenta string) (*Cuenta, error) {
	query := "SELECT id, numero_de_cuenta, saldo, id_cliente, id_tipo_de_cuenta, habilitado, fecha_creacion FROM cuentas WHERE numero_de_cuenta=?"
	return getCuenta(query, numeroDeCuenta)
}

//GetCuentaByNumeroTarjeta Solicita ala base de datos una cuenta dando de
//referencia el numero de tarjeta
func GetCuentaByNumeroTarjeta(numeroDeTarjeta string) (*Cuenta, error) {
	query := "SELECT id, numero_de_cuenta, saldo, id_cliente, id_tipo_de_cuenta, habilitado, fecha_creacion FROM cuentas WHERE id in (SELECT id_cuenta FROM tarjetas WHERE numero_tarjeta=?)"
	return getCuenta(query, numeroDeTarjeta)
}

//GetCuentas Solicita todas las cuentas de la base de datos
func GetCuentas() (Cuentas, error) {
	var cuentas Cuentas
	query := "SELECT id, numero_de_cuenta, saldo, id_cliente, id_tipo_de_cuenta, habilitado, fecha_creacion FROM cuentas WHERE habilitado=1"
	rows, err := Query(query)
	for rows.Next() {
		cuenta := Cuenta{}
		rows.Scan(&cuenta.ID, &cuenta.NumeroDeCuenta, &cuenta.Saldo, &cuenta.IDCliente, &cuenta.IDTipoDeCuenta, 
            &cuenta.habilitado, &cuenta.fechaCreacion)
        cuentas = append(cuentas, cuenta)
	}
	return cuentas, err
}

//Depositar Valida la cuenta para que un deposito sea efectuado
//y solicita a otro metodo el deposito
func (cuenta *Cuenta) Depositar(monto float32) error {
	cuenta.Saldo += monto
	err := cuenta.Guardar()
	if err == nil {
		cuenta.activarCuenta()
	}
	return err
}

//Retirar Esta funcion valida la solicitud de retiro de efectivo de una cuenta
func (cuenta *Cuenta) Retirar(monto float32) error {
	if cuenta.Saldo >= monto {
		cuenta.Saldo -= monto
		return cuenta.Guardar()
	} else {
		return errors.New("Saldo insuficiente")
	}
}

//Transferir Esta funcion valida y solicita la transferencia de dinero de cuenta a cuenta
func (cuenta *Cuenta) Transferir(numeroCuentaDestino string, monto float32) error {
	cuentaDestino := &Cuenta{}
	err := errors.New("")

	if cuentaDestino, err = GetCuentaByNumeroCuenta(numeroCuentaDestino); err != nil {
		return err
	}
	if err = cuenta.Retirar(monto); err != nil {
		return err
	}
	err = cuentaDestino.Depositar(monto)
	return err
}

//SolicitarSaldo Este metodo solicita a la base de datos el saldo de la cuenta
func (cuenta *Cuenta) SolicitarSaldo() (float32, error) {
	var saldo float32
	query := "SELECT saldo FROM cuentas WHERE numero_de_cuenta = ?"
	rows, err := Query(query, cuenta.NumeroDeCuenta)
	if err != nil {
		return saldo, err
	}
	for rows.Next() {
		rows.Scan(&saldo)
	}
	return saldo, nil
}

//Guardar Valida la cuenta y solicita su registro en la base de datos
func (cuenta *Cuenta) Guardar() error {
	if cuenta.ID == 0 {
		return cuenta.registrar()
	} 

	return cuenta.actualizar()
}

//registrar Realiza el registro de la cuenta en la base de datos
func (cuenta *Cuenta) registrar() error {
	cuenta.NumeroDeCuenta="5050"+RandomDigits(16)
	cuenta.fechaCreacion=ObtenerFechaHoraActualString()
	query := "INSERT INTO cuentas(numero_de_cuenta, saldo, id_cliente, id_tipo_de_cuenta, habilitado, fecha_creacion) VALUES(?,?,?,?,?,?);"
	cuentaID, err := InsertData(query, cuenta.NumeroDeCuenta, cuenta.Saldo, cuenta.IDCliente, cuenta.IDTipoDeCuenta, 
        cuenta.habilitado, cuenta.fechaCreacion)
	cuenta.ID = int(cuentaID)
	return err
}

//ActualizarCuenta habilita el estado de la cuenta para poder hacer transacciones
func (cuenta *Cuenta) ActualizarCuenta() error {
	cuenta.habilitado = 1
	return cuenta.actualizar()
}

//actualizar metodo que cambia algun dato de la cuenta 
func (cuenta *Cuenta) actualizar() error {
	query := "UPDATE cuentas SET numero_de_cuenta=?, saldo=?, id_cliente=?, id_tipo_de_cuenta=?, habilitado=?, fecha_creacion=? WHERE id=?;"
	_, err := Exec(query, cuenta.NumeroDeCuenta, cuenta.Saldo, cuenta.IDCliente, cuenta.IDTipoDeCuenta, cuenta.habilitado, 
		cuenta.fechaCreacion, cuenta.ID)
	return err
}

//estaActivada Valida el estado de la cuenta
func (cuenta *Cuenta) estaActivada() bool {	
	if cuenta.habilitado == 1 {
		return true
	}
	return false
}

//activarCuenta Actualiza el estado de la cuenta en la base de datos 
func (cuenta *Cuenta) activarCuenta() error {
	var err error

	if !cuenta.estaActivada() {
		query := "UPDATE cuentas SET habilitado=1 WHERE id=?"
		_, err = Exec(query, cuenta.ID)
		if err == nil {
			cuenta.habilitado = 1
		}
	}
	return err
}

//GetFechaCreacion obtiene la fecha de creacion de la cuenta
func (cuenta *Cuenta) GetFechaCreacion() string {
	return cuenta.fechaCreacion
}

//SetFechaCreacion Envia la fecha para la creacion de la cuenta
func (cuenta *Cuenta) SetFechaCreacion(fechaCreacion string) {
	cuenta.fechaCreacion = fechaCreacion
}
