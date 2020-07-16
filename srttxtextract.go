package srttxtextract

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
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

func IsIndexLine(line string) bool {
	matched, _ := regexp.MatchString(`[0-9]+`, line)
	return matched
}

func IsEmptyLine(line string) bool {
	return strings.TrimSpace(line) == ""
}

func ReadSrtFile(filePath string) (err error) {
	fmt.Println(filePath)
	lines, err := FileToLines(filePath)
	if err != nil {
		return
	}

	state := StartLine
	for _, line := range lines {
		if state == StartLine && IsIndexLine(line) {
			state = IndexLine
		}
		if state == EmptyLine && IsIndexLine(line) {
			state = IndexLine
		}
	}

	return
}

func ReadSrtDir(dir string) (err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	for _, file := range files {
		err = ReadSrtFile(path.Join(dir, file.Name()))
		if err != nil {
			return
		}
	}
	return
}
