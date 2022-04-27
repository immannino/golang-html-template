package html

import (
	"bytes"
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var t *template.Template
var debug string

//go:embed templates/*.html
var fs embed.FS

func init() {
	t = template.Must(template.ParseFS(fs, "templates/*.html"))
	debug = os.Getenv("DEBUG")
}

func HTML(w http.ResponseWriter, httpStatus int, templateName string, data interface{}) {
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	if len(debug) > 0 {
		d := template.Must(template.ParseGlob(filepath.Join("./html/templates", "*.html")))
		tmp := d.Lookup(templateName)
		err := tmp.Execute(w, data)

		if err != nil {
			log.Println("Error executing template :", err)
		}

		return
	}

	tmp := t.Lookup(templateName)
	err := tmp.Execute(w, data)

	if err != nil {
		log.Println("Error executing template :", err)
	}
}

func HTMLBytes(name string, data interface{}) []byte {
	var tpl bytes.Buffer

	if len(debug) > 0 {
		d := template.Must(template.ParseGlob(filepath.Join("./html/templates", "*.html")))
		tmp := d.Lookup(name)
		err := tmp.Execute(&tpl, data)

		if err != nil {
			log.Println("Error executing template :", err)
		}

		return tpl.Bytes()
	}

	tmp := t.Lookup(name)
	err := tmp.Execute(&tpl, data)

	if err != nil {
		log.Println("Error executing template :", err)
	}

	return tpl.Bytes()
}
