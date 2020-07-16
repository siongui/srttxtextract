package srttxtextract

import (
	"testing"
)

func TestFileToLines(t *testing.T) {
	err := ReadSrtDir("testdata/")
	if err != nil {
		t.Error(err)
		return
	}
}
