package srttxtextract

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type LineType int

const (
	StartLine LineType = iota
	IndexLine
	TimeLine
	TextLine
	EmptyLine
)

func FileToLines(filePath string) (lines []string, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

func PrintLines(filePath string) (err error) {
	fmt.Println(filePath)
	return
}

func ReadSrtDir(dir string) (err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	for _, file := range files {
		err = PrintLines(path.Join(dir, file.Name()))
		if err != nil {
			return
		}
	}
	return
}
