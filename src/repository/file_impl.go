package repository

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aikizoku/rundoc/src/log"
)

type file struct {
}

func (r *file) GetNameList(dirPath string) ([]string, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Errorf(err, "ディレクトリ作成に失敗: %s", dirPath)
		return []string{}, err
	}
	var names []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		names = append(names, file.Name())
	}
	return names, nil
}

func (r *file) Exist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func (r *file) WriteDir(path string) error {
	if r.Exist(path) {
		return nil
	}
	err := os.Mkdir(path, 0755)
	if err != nil {
		log.Errorf(err, "ディレクトリ作成に失敗: %s", path)
		return err
	}
	return nil
}

func (r *file) Write(path string, body string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Errorf(err, "ファイルオープンに失敗: %s", path)
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, body)
	if err != nil {
		log.Errorf(err, "ファイル書き込みに失敗: %s, %s", path)
		return err
	}
	return nil
}

func (r *file) Remove(path string) error {
	if err := os.Remove(path); err != nil {
		log.Errorf(err, "ファイル削除に失敗: %s", path)
		return err
	}
	return nil
}

// NewFile ... リポジトリを作成する
func NewFile() File {
	return &file{}
}
