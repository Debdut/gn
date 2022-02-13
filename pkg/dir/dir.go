package dir

import (
	"errors"
	"os"
	"path"
	"strings"
)

var configs = []string{"next.config.js", "package.json"}
var pageDirs = []string{"pages", "src/pages"}

func GetNextPageRoot() (string, error) {
	root, err := GetNextRoot()
	if err != nil {
		return root, err
	}

	for i := 0; i < len(pageDirs); i++ {
		pageDir := join(pageDirs[i], root)
		if exists(pageDir) {
			return pageDir, nil
		}
	}

	return "", errors.New("pages root not found")
}

func GetNextRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for pathLen(dir) > 2 {
		for _, config := range configs {
			configPath := join(config, dir)
			if exists(configPath) {
				return dir, nil
			}
		}

		dir = path.Dir(dir)
	}

	return "", errors.New("root not found")
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
