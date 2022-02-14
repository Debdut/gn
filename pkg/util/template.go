package util

import (
	"io"
	"text/template"
)

func GenTemplate(data interface{}, writer io.Writer, tmpl string, tmplName string) error {
	template, err := template.New(tmplName).Parse(tmpl)
	if err == nil {
		err = template.Execute(writer, data)
	}
	return err
}
