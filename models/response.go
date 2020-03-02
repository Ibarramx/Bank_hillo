package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	writer      http.ResponseWriter
}

//CreateDefaultResponse method
func CreateDefaultResponse(w http.ResponseWriter) Response {
	return Response{Status: http.StatusOK, writer: w, contentType: "application/json"}
}

func SendPaymentRequired(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.PaymentRequired()
	response.Send()
}

func (r *Response) PaymentRequired() {
	r.Status = http.StatusPaymentRequired
	r.Message = "payment couldn't be processed"
}

//SendNotFound method
func SendNotFound(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NotFound()
	response.Send()
}

//NotFound method
func (r *Response) NotFound() {
	r.Status = http.StatusNotFound
	r.Message = "Resource not found."
}

//SendUnprocessableEntity method
func SendUnprocessableEntity(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.UnprocessableEntity()
	response.Send()
}

//UnprocessableEntity method
func (r *Response) UnprocessableEntity() {
	r.Status = http.StatusUnprocessableEntity
	r.Message = "Unprocesable Entity"
}

//Send method
func (r *Response) Send() {
	r.writer.Header().Set("Content-Type", r.contentType)
	r.writer.WriteHeader(r.Status)

	output, _ := json.Marshal(&r)
	fmt.Fprintf(r.writer, string(output))
}

//SendData method
func SendData(w http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(w)
	response.Data = data
	response.Send()
}

//SendNoContent method
func SendNoContent(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NoContent()
	response.Send()
}

//NoContent method
func (r *Response) NoContent() {
	r.Status = http.StatusNoContent
	r.Message = "No Content."
}
