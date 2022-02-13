package dir

import (
	"errors"
	"os"
	"path"
	"strings"
)

var configs = []string{"next.config.js", "package.json"}
var pageDirs = []string{"pages", "src/pages"}

type NextDirs struct {
	root string
	page string
	api  string
}

func GetNextDirs() NextDirs {
	root, _ := GetNextRoot()
	page, _ := GetNextPageRoot(root)
	api, _ := GetNextApiRoot(page)

	return NextDirs{root, page, api}
}

func GetNextApiRoot(pageRoot string) (string, error) {
	apiDir := join("api", pageRoot)
	if exists(apiDir) {
		return apiDir, nil
	}
	return apiDir, errors.New("api dir not found")
}

func GetNextPageRoot(root string) (string, error) {
	for i := 0; i < len(pageDirs); i++ {
		pageDir := join(pageDirs[i], root)
		if exists(pageDir) {
			return pageDir, nil
		}
	}

	return join(pageDirs[0], root), errors.New("pages root not found")
}

func GetNextRoot() (string, error) {
	dir := cwd()

	for pathLen(dir) > 2 {
		for _, config := range configs {
			configPath := join(config, dir)
			if exists(configPath) {
				return dir, nil
			}
		}

		dir = path.Dir(dir)
	}

	return cwd(), errors.New("root not found")
}

func pathLen(path string) int {
	return strings.Count(path, string(os.PathSeparator))
}

func join(fileName string, dir string) string {
	return dir + string(os.PathSeparator) + fileName
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func cwd() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return dir
}
