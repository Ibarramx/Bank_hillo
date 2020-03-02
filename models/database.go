package models

import (
	"database/sql"
	"fmt"
	"log"

	_"github.com/mattn/go-sqlite3"
)

var db *sql.DB

//init Inicializa la coneccion con la base de datos y crea las tablas
func init() {
	CreateConnection()
	CreateTables()
}

//CreateConnection crea la conexion con la base de datos
func CreateConnection() {

	if GetConnection() != nil {
		return
	}

	if connection, err := sql.Open("sqlite3", "./banco.db"); err != nil {
		panic(err)
	} else {
		db = connection
	}
}

//CreateTables Realiza transacciones para la creacion de las tablas de la base de datos
func CreateTables() {
	createTable("clientes", clienteSchemeSQLITE)
	createTable("tipos_cuenta", tipoCuentaSchemeSQLITE)
	createTable("cuentas", cuentaSchemeSQLITE)
	createTable("empleados", empleadosSchemeSQLITE)
	createTable("tipos_transaccion", tipoTransaccionSchemeSQLITE)
	createTable("transacciones", transaccionSchemeSQLITE)
	createTable("tarjetas", tarjetaSchemeSQLITE)
	createTable("logins", loginSchemeSQLITE)
}

/*func createTable(tableName, scheme string) {
	if !existsTable(tableName) {
		Exec(scheme)
	} else {
		truncateTable(tableName)
	}
}*/

//createTable ejecta las transacciones para la creacion de las tablas
func createTable(tableName, scheme string) {
	Exec(scheme)	
}

//truncateTable Elimina una tabla de la base de datos
func truncateTable(tableName string) {
	sql := fmt.Sprintf("DELETE FROM %s", tableName)
	Exec(sql)
}

//existsTable Valida si la tabla existe
func existsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, _ := Query(sql)
	return rows.Next()
}

//Exec Ejecuta las transacciones en la base de datos
func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	/*if err != nil && !debug {
		log.Println(err)
	}*/
	if err != nil {
		log.Println(err)
	}
	return result, err
}

//Query Crea el query para las transacciones
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	/*if err != nil && !debug {
		log.Println(err)
	}*/
	if err != nil{
		log.Println(err)
	}
	return rows, err
}

//InsertData Metodo para realizar inserciones a la base de datos
func InsertData(query string, args ...interface{}) (int64, error) {
	result, err := Exec(query, args...)
	if err != nil {
		return int64(0), err
	}
	id, err := result.LastInsertId()
	return id, err
}

//GetConnection Obtiene la coneccion para realizar transacciones
func GetConnection() *sql.DB {
	return db
}

//Ping 
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

//CloseConnection Cierra la conexion con la base de datos
func CloseConnection() {
	db.Close()
}
 