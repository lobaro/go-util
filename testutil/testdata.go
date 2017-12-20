package testutil

import (
	"io/ioutil"
	"path/filepath"
	"os"
)

func hasSubdir(path, name string) bool {
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, fi := range dirs {
		if fi.Name() == name {
			return true
		}
	}
	return false
}

func TestdataDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for !hasSubdir(pwd, "testdata") {
		prev := pwd
		pwd = filepath.Clean(filepath.Join(pwd, ".."))
		if pwd == prev {
			panic("no testdata directory found")
			break
		}
	}

	return filepath.Clean(filepath.Join(pwd, "testdata"))
}

func Testdata(path string) string {
	return filepath.Join(TestdataDir(), path)
}