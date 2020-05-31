package armtemplate_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func testdata(t *testing.T, file string) string {
	b, err := ioutil.ReadFile(filepath.Join("testdata", file))
	if err != nil {
		t.Fatal(err)
	}

	return string(b)
}
