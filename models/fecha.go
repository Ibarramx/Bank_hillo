package models

import (
	"time"
	"fmt"
)

//ObtenerFechaHoraActualString Obtiene la fecha del sistema en string
func ObtenerFechaHoraActualString() string {
	t := time.Now()

	fechaHora := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", 
		t.Year(), t.Month(), t.Day(), 
		t.Hour(), t.Minute(), t.Second())

	return fechaHora
}

//GetFechaVencimientoString obtiene fecha de vencimiento de tarjeta
func GetFechaVencimientoString() string {
	t := time.Now()
	mes := fmt.Sprintf("%02d",t.Month())
	year := fmt.Sprintf("%02d",t.Year()-1998)
	fecha := mes+"/"+year
	return fecha
}
