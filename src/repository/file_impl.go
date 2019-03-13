package repository

import (
	"fmt"
	"io/ioutil"
	"os"
)

type file struct {
}

func (r *file) GetNameList(dirPath string) []string {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}
	var names []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		names = append(names, file.Name())
	}
	return names
}

func (r *file) Exist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func (r *file) WriteDir(path string) {
	if r.Exist(path) {
		return
	}
	err := os.Mkdir(path, 0755)
	if err != nil {
		panic(err)
	}
}

func (r *file) Write(path string, body string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, body)
	if err != nil {
		panic(err)
	}
}

func (r *file) Remove(path string) {
	if err := os.Remove(path); err != nil {
		panic(err)
	}
}

// NewFile ... リポジトリを作成する
func NewFile() File {
	return &file{}
}
