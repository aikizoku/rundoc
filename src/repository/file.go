package repository

// File ... ファイル操作に関するリポジトリ
type File interface {
	GetNameList(dir string) []string
	Exist(path string, name string) bool
	Write(path string, name string, body string)
	Remove(path string, name string)
}
