package repository

// File ... ファイル操作に関するリポジトリ
type File interface {
	GetNameList(dirPath string) ([]string, error)
	Exist(path string) bool
	WriteDir(path string) error
	Write(path string, body string) error
	Remove(path string) error
}
