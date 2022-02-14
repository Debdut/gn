package next

import (
	_ "embed"
	"io"

	"github.com/debdut/gn/pkg/util"
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
	return util.GenTemplate(page, writer, pageTmpl, "Page")
}

func GenApi(api Api, writer io.Writer) error {
	return util.GenTemplate(api, writer, apiTmpl, "Api")
}
