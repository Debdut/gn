package next

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"

	template "github.com/debdut/gn/pkg/template/next"
	"github.com/debdut/gn/pkg/util"
)

var configs = []string{
	"next.config.js",
	"package.json",
	"tsconfig.json",
}
var nextConfigs []string = configs[:1]
var pageDirs = []string{"pages", "src/pages"}

type NextDirs struct {
	Root string
	Page string
	Api  string
}

func WriteApiTemplate(name string, dir string, ts bool) error {
	getData := func(name string, ts bool) interface{} {
		api := template.Api{
			Api: strings.Title(name),
			TS:  ts,
		}

		return api
	}

	getExtension := func(ts bool) string {
		fileName := name + "."
		if ts {
			fileName += "ts"
		} else {
			fileName += "js"
		}

		return fileName
	}

	genTemplate := func(data interface{}, file *os.File) error {
		return template.GenApi(data.(template.Api), file)
	}

	return WriteTemplate(name, dir, ts, getData, getExtension, genTemplate)
}

func WritePageTemplate(name string, dir string, ts bool) error {
	getData := func(name string, ts bool) interface{} {
		page := template.Page{
			Page: strings.Title(name),
			TS:   ts,
		}

		return page
	}

	getExtension := func(ts bool) string {
		fileName := name + "."
		if ts {
			fileName += "tsx"
		} else {
			fileName += "jsx"
		}

		return fileName
	}

	genTemplate := func(data interface{}, file *os.File) error {
		return template.GenPage(data.(template.Page), file)
	}

	return WriteTemplate(name, dir, ts, getData, getExtension, genTemplate)
}

func WriteTemplate(
	name string,
	dir string,
	ts bool,
	getData func(string, bool) interface{},
	getExtension func(bool) string,
	genTemplate func(interface{}, *os.File) error,
) error {

	TS := IsTypescript() || ts
	fileName := getExtension(TS)
	data := getData(name, TS)
	filePath := filepath.Join(dir, fileName)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	err = genTemplate(data, file)
	return err
}

func IsTypescript() bool {
	root, _ := GetNextRoot()
	tsConf := filepath.Join(root, "tsconfig.json")
	return util.Exists(tsConf)
}

func GetConfigs() []string {
	var confs []string
	root, _ := GetNextRoot()

	for _, conf := range configs {
		configPath := filepath.Join(root, conf)
		if util.Exists(configPath) {
			confs = append(confs, conf)
		}
	}

	return confs
}

func GetNextDirs() NextDirs {
	root, _ := GetNextRoot()
	page, _ := GetNextPageRoot(root)
	api, _ := GetNextApiRoot(page)

	return NextDirs{Root: root, Page: page, Api: api}
}

func GetNextApiRoot(pageRoot string) (string, error) {
	apiDir := filepath.Join(pageRoot, "api")
	if util.Exists(apiDir) {
		return apiDir, nil
	}
	return apiDir, errors.New("api dir not found")
}

func GetNextPageRoot(root string) (string, error) {
	for i := 0; i < len(pageDirs); i++ {
		pageDir := filepath.Join(root, pageDirs[i])
		if util.Exists(pageDir) {
			return pageDir, nil
		}
	}

	return filepath.Join(root, pageDirs[0]), errors.New("pages root not found")
}

func GetNextRoot() (string, error) {
	dir := util.Cwd()

	for util.PathLen(dir) > 2 {
		for _, config := range nextConfigs {
			configPath := filepath.Join(dir, config)
			if util.Exists(configPath) {
				return dir, nil
			}
		}

		dir = path.Dir(dir)
	}

	return util.Cwd(), errors.New("root not found")
}
