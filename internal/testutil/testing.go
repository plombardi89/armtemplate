package testutil

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func Data(t *testing.T, file string) string {
	b, err := ioutil.ReadFile(filepath.Join("testdata", file))
	if err != nil {
		t.Fatal(err)
	}

	return string(b)
}

func Jsonify(t *testing.T, v interface{}) string {
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	return string(b)
}
