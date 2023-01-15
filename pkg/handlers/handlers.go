package handlers

import (
	"fmt"
	"go_module/pkg/models"
	"go_module/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w ,"home.page.tmpl",&models.TemplateData{})

}
func About(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w,"about.page.tmpl",&models.TemplateData{})
}
func GetData(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w,"form.page.tmpl",&models.TemplateData{})
}

func PostData(w http.ResponseWriter, r *http.Request) {
	first:=r.FormValue("first")
	last:=r.FormValue("last")
	email:=r.FormValue("email")

	fmt.Println(first )
	fmt.Println(last )
	fmt.Println(email )

	stringMap:=make(map[string]string)
	stringMap["first"]=first
	stringMap["last"]=last
	stringMap["email"]=email
	render.RenderTemplate(w,"output.page.tmpl",&models.TemplateData{
		StringMap: stringMap,
	})
}