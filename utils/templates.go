package utils

import(
  "log"
  "net/http"
  "html/template"

)

//Variables Para el manejo de vistas
var templates = template.Must(template.New("t").ParseGlob(TemplatesDir()))
var errorTemplate = template.Must(template.ParseFiles(ErrorTemplateDir()))
  
//RenderErrorTemplate Metodo que solicita y renderisa la vista de error,
//si se presento alguno
func RenderErrorTemplate(w http.ResponseWriter, status int){
  w.WriteHeader(status)
  errorTemplate.Execute(w, nil)
}

//RenderTemplate Metodo que solicita y renderisa la vista solicitada
func RenderTemplate(w http.ResponseWriter, name string, data interface{}){
  w.Header().Set("Content-Type", "text/html")
  
  err := templates.ExecuteTemplate(w, name, data)
  
  if err != nil{
    log.Println(err)
    RenderErrorTemplate(w, http.StatusInternalServerError)
  }
}

//TemplatesDir Busca la vista solicitada y la retorna al metodo
func TemplatesDir() string{
  //return fmt.Sprintf("%s/templates/**/*.html", server.templateDir)
  return "templates/**/*.html"
}

//ErrorTemplateDir Agarra la vista Error y la envia al metodo
func ErrorTemplateDir() string{
  //return fmt.Sprintf("%s/templates/error.html", server.templateDir)
  return "templates/error.html"
}
