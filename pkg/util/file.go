package util

import (
	"os"
	"strings"
)

func PathLen(path string) int {
	return strings.Count(path, string(os.PathSeparator))
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func Cwd() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return dir
}
