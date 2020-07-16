package srttxtextract

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestFileToLines(t *testing.T) {
	files, err := ioutil.ReadDir("testdata/")
	if err != nil {
		t.Error(err)
		return
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
