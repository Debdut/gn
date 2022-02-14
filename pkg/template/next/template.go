package next

import (
	_ "embed"
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

//go:embed template/page.tmpl
var pageTmpl string

//go:embed template/api.tmpl
var apiTmpl string

func GenPage(page Page, writer io.Writer) error {
	template, err := template.New("Page").Parse(pageTmpl)
	if err == nil {
		err = template.Execute(writer, page)
	}
	return err
}

func GenApi(api Api, writer io.Writer) error {
	template, err := template.New("Api").Parse(apiTmpl)
	if err == nil {
		err = template.Execute(writer, api)
	}
	return err
}
