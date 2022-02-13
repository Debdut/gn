package dir

import (
	"errors"
	"os"
	"path"
	"strings"
)

var (
	nextConfig string = "next.config.js"
	pkg        string = "package.json"
)

type Dir struct {
	path string `default:""`
	conf string `default:""`
}

func GetNextRoot() (Dir, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	for pathLen(currentDir) > 2 {
		nextConfigPath := getFilePath(nextConfig, currentDir)
		_, err := os.Stat(nextConfigPath)
		if err == nil {
			return Dir{path: currentDir, conf: nextConfig}, nil
		}

		pkgPath := getFilePath(pkg, currentDir)
		_, err = os.Stat(pkgPath)
		if err == nil {
			return Dir{path: currentDir, conf: pkg}, nil
		}

		currentDir = path.Dir(currentDir)
	}

	return Dir{}, errors.New("root not found")
}

func pathLen(path string) int {
	return strings.Count(path, string(os.PathSeparator))
}

func getFilePath(fileName string, dir string) string {
	return dir + string(os.PathSeparator) + fileName
}
