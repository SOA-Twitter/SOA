package data

import (
	"encoding/json"
	"io"
	"net/http"
)

func DecodeBody(r io.Reader) (*Tweet, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var rt Tweet
	if err := dec.Decode(&rt); err != nil {
		return nil, err
	}
	return &rt, nil
}
func RenderJson(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
func ExecuteQuery(query string, values ...interface{}) error {
	if err := Session.Query(query).Bind(values...).Exec(); err != nil {
		return err
	}
	return nil
}
