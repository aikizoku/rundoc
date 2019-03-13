package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	_ "github.com/aikizoku/rundoc/statik" // バイナリ化したファイル
	"github.com/rakyll/statik/fs"
)

func convertPrettyJSON(body []byte) string {
	out := new(bytes.Buffer)
	err := json.Indent(out, body, "", "    ")
	if err != nil {
		panic(err)
	}
	return out.String()
}

func getBinFileData(name string) []byte {
	st, err := fs.New()
	if err != nil {
		panic(err)
	}

	f, err := st.Open("/" + name)
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return b
}
