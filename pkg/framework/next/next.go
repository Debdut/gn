package next

import (
	"errors"
	"os"
	"path"
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

func WritePageTemplate(name string, dir string, ts bool) error {
	TS := IsTypescript() || ts
	fileName := name + "."
	if TS {
		fileName += "tsx"
	} else {
		fileName += "jsx"
	}

	page := template.Page{
		Page: strings.Title(name),
		TS:   TS,
	}
	filePath := util.Join(fileName, dir)

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	err = template.GenPage(page, file)

	return err
}

func IsTypescript() bool {
	root, _ := GetNextRoot()
	tsConf := util.Join("tsconfig.json", root)
	return util.Exists(tsConf)
}

func GetConfigs() []string {
	var confs []string
	root, _ := GetNextRoot()

	for _, conf := range configs {
		configPath := util.Join(conf, root)
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
	apiDir := util.Join("api", pageRoot)
	if util.Exists(apiDir) {
		return apiDir, nil
	}
	return apiDir, errors.New("api dir not found")
}

func GetNextPageRoot(root string) (string, error) {
	for i := 0; i < len(pageDirs); i++ {
		pageDir := util.Join(pageDirs[i], root)
		if util.Exists(pageDir) {
			return pageDir, nil
		}
	}

	return util.Join(pageDirs[0], root), errors.New("pages root not found")
}

func GetNextRoot() (string, error) {
	dir := util.Cwd()

	for util.PathLen(dir) > 2 {
		for _, config := range nextConfigs {
			configPath := util.Join(config, dir)
			if util.Exists(configPath) {
				return dir, nil
			}
		}

		dir = path.Dir(dir)
	}

	return util.Cwd(), errors.New("root not found")
}
