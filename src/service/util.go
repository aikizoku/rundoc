package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/aikizoku/rundoc/src/log"
	_ "github.com/aikizoku/rundoc/statik" // バイナリ化したファイル
	"github.com/rakyll/statik/fs"
)

func convertPrettyJSON(body []byte) (string, error) {
	out := new(bytes.Buffer)
	err := json.Indent(out, body, "", "    ")
	if err != nil {
		log.Errorf(err, "jsonのparseに失敗: %s", string(body))
		return "", err
	}
	return out.String(), nil
}

func getBinFileData(name string) ([]byte, error) {
	st, err := fs.New()
	if err != nil {
		log.Errorf(err, "組み込みファイル初期化に失敗")
		return nil, err
	}

	f, err := st.Open("/" + name)
	if err != nil {
		log.Errorf(err, "組み込みファイルオープンに失敗: %s", name)
		return nil, err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Errorf(err, "組み込みファイル読み込みに失敗: %s", name)
		return nil, err
	}

	return b, nil
}
