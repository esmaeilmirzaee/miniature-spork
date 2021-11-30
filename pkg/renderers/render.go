package renderers

import (
	"github.com/esmaeilmirzaee/random-time-sleeper/helpers"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, templateName string) {
	templateFile, err := template.ParseFiles("./templates/"+templateName)
	if err != nil {
		helpers.HandleErrors(err)
		return
	}
	err = templateFile.Execute(w, nil)
}