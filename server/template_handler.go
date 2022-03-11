package server

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type templateHandlerFactory struct {
	tmpl          *template.Template
	MongoDBClient mongoDBClient
}

func (rt Router) NewTemplateHandlerFactory(templateDirPath string) templateHandlerFactory {
	var tmpl *template.Template
	tmpl = template.Must(tmpl.ParseGlob(filepath.Join(templateDirPath, "*.gohtml")))
	template.Must(tmpl.ParseGlob(filepath.Join(templateDirPath, "partials/*.gohtml")))

	return templateHandlerFactory{tmpl, rt.MongoDBClient}
}

type templateData struct {
	PageName string
}

func (t templateHandlerFactory) DefaultHandler(templateName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.NotFound(w, r)
			return
		}

		pageName := "Default Page"

		data := &templateData{
			PageName: pageName,
		}

		err := t.tmpl.ExecuteTemplate(w, templateName, data)
		if err != nil {
			println(err.Error())
			http.NotFound(w, r)
			return
		}

	}
}
