package render

import (
	"go_module/pkg/models"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, url string,td *models.TemplateData) {
	tmpl, _ := template.ParseFiles("./templates/"+url,"./templates/base.layout.tmpl")

	tmpl.Execute(w,td)
}