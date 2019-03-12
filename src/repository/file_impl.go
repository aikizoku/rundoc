package repository

import (
	"fmt"
	"io/ioutil"
	"os"
)

type file struct {
}

func (r *file) GetNameList(dir string) []string {
	files, err := ioutil.ReadDir(dir)
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

func (r *file) Exist(path string, name string) bool {
	_, err := os.Stat(path + name)
	return err == nil
}

func (r *file) Write(path string, name string, body string) {
	file, err := os.OpenFile(path+name, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, body)
	if err != nil {
		panic(err)
	}
}

func (r *file) Remove(path string, name string) {
	if err := os.Remove(path + name); err != nil {
		panic(err)
	}
}

// NewFile ... リポジトリを作成する
func NewFile() File {
	return &file{}
}
