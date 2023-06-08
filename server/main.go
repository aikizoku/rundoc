package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/unrolled/render"
)

func main() {
	r := chi.NewRouter()
	r.Route("/sample", func(r chi.Router) {
		r.Get("/", formHandler)
		r.Post("/", jsonHandler)
		r.Put("/", jsonHandler)
		r.Delete("/", formHandler)
		r.Get("/error", errorHandler)
	})
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}

type Param struct {
	Text        string   `form:"text"         json:"text"`
	Number      int      `form:"number"       json:"number"`
	ArrayText   []string `form:"array_text"   json:"array_text"`
	ArrayNumber []int    `form:"array_number" json:"array_number"`
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	param := &Param{}
	param.Text = getForm(r, "text")
	param.Number = getFormByNumber(r, "number")
	param.ArrayText = getFormByArray(r, "array_text")
	param.ArrayNumber = getFormByArrayNumber(r, "array_number")
	renderJSON(w, param)
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	param := &Param{}
	getJSON(r, param)
	renderJSON(w, param)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte("<html>\n<body>test error</body>\n</html>"))
	if err != nil {
		panic(err)
	}
}

func getForm(r *http.Request, key string) string {
	return r.FormValue(key)
}

func getFormByNumber(r *http.Request, key string) int {
	v := r.FormValue(key)
	i, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return i
}

func getFormByArray(r *http.Request, key string) []string {
	sKey := fmt.Sprintf("%s[]", key)
	qs := r.URL.RawQuery
	vs := []string{}
	var err error
	for _, q := range strings.Split(qs, "&") {
		kv := strings.Split(q, "=")
		if len(kv) < 2 {
			continue
		}
		k := kv[0]
		k, err = url.QueryUnescape(k)
		if err != nil {
			continue
		}
		if k != sKey {
			continue
		}
		v := kv[1]
		v, err = url.QueryUnescape(v)
		if err != nil {
			continue
		}
		vs = append(vs, v)
	}
	return vs
}

func getFormByArrayNumber(r *http.Request, key string) []int {
	sKey := fmt.Sprintf("%s[]", key)
	qs := r.URL.RawQuery
	vs := []int{}
	var err error
	for _, q := range strings.Split(qs, "&") {
		kv := strings.Split(q, "=")
		if len(kv) < 2 {
			continue
		}
		k := kv[0]
		k, err = url.QueryUnescape(k)
		if err != nil {
			continue
		}
		if k != sKey {
			continue
		}
		v := kv[1]
		v, err = url.QueryUnescape(v)
		if err != nil {
			continue
		}
		n, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		vs = append(vs, n)
	}
	return vs
}

func getJSON(r *http.Request, dst *Param) {
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(dst)
	if err != nil {
		panic(err)
	}
}

func renderJSON(w http.ResponseWriter, param *Param) {
	ren := render.New()
	err := ren.JSON(w, http.StatusOK, param)
	if err != nil {
		panic(err)
	}
}
