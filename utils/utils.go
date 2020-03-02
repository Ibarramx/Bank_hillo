package utils

import (
	"net/http"
	"strings"
)

//enableCors Habilita el mecanisco CORS
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

//HideCard Metodo el cual oculta los digitos de la tarjeta,
// dejando los ultimos 4 visibles

func HideCard(tarjeta string) string {
	s := strings.Split(tarjeta, "")
	hidden:="************"
	for i:=12; i<len(s); i++{
		hidden+=s[i]
	}
	return hidden
}
