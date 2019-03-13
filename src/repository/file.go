package repository

// File ... ファイル操作に関するリポジトリ
type File interface {
	GetNameList(dirPath string) []string
	Exist(path string) bool
	WriteDir(path string)
	Write(path string, body string)
	Remove(path string)
}
