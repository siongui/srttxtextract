package main

import (
	. "github.com/siongui/godom"
	"github.com/siongui/srttxtextract"
)

func main() {
	input := Document.QuerySelector("#input")
	output := Document.QuerySelector("#output")

	Document.GetElementById("extract").AddEventListener("click", func(e Event) {

		s := input.Value()
		lines, err := srttxtextract.StringToLines(s)
		if err != nil {
			output.SetValue(err.Error())
		}

		texts := srttxtextract.SrtFileLinesToTexts(lines)
		output.SetValue(texts)
	})
}
