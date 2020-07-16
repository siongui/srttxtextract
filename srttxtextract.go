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

func StringToLines(s string) (lines []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

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
	/*
		if matched {
			fmt.Println(line)
		}
	*/
	return matched
}

func IsEmptyLine(line string) bool {
	return strings.TrimSpace(line) == ""
}

// state machine:
// Start -> Index -> Time -> Text -> Empty -> Index -> ...
func NextState(state LineType, line string) LineType {
	if state == StartLine && IsIndexLine(line) {
		return IndexLine
	}
	if state == EmptyLine && IsIndexLine(line) {
		return IndexLine
	}
	if state == IndexLine {
		return TimeLine
	}
	if state == TimeLine {
		return TextLine
	}
	if state == TextLine && IsEmptyLine(line) {
		return EmptyLine
	}
	return state
}

func SrtFileLinesToTexts(lines []string) (texts string) {
	state := StartLine
	for _, line := range lines {
		//fmt.Println("current state:", state)
		state = NextState(state, line)
		if state == TextLine {
			texts = texts + line + "\n"
		}
	}
	return
}

func ReadSrtFileTexts(filePath string) (texts string, err error) {
	fmt.Println(filePath)
	lines, err := FileToLines(filePath)
	if err != nil {
		return
	}
	return SrtFileLinesToTexts(lines), err
}

func ReadSrtDir(dir string) (err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	for _, file := range files {
		texts, err := ReadSrtFileTexts(path.Join(dir, file.Name()))
		if err != nil {
			return err
		}
		fmt.Println(texts)
	}
	return
}
