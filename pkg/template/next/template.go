package next

import (
	"io"
	"text/template"
)

type Page struct {
	Page string
	TS   bool
}

type Api struct {
	Api string
	TS  bool
}

var nextTemplateRoot string = "template/next/"

func GenPage(page Page, writer io.Writer) error {
	template, err := template.ParseFiles(nextTemplateRoot + "page.tmpl")
	if err == nil {
		err = template.Execute(writer, page)
	}
	return err
}

func GenApi(api Api, writer io.Writer) error {
	template, err := template.ParseFiles(nextTemplateRoot + "api.tmpl")
	if err == nil {
		err = template.Execute(writer, api)
	}
	return err
}
