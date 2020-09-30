package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//debug
	//os.Args = []string{"kindmaker", "uoop.txt"}

	if len(os.Args) != 2 {
		fmt.Println("USAGE : kindleconv textfile.txt")
		return
	}

	filenamewithext := os.Args[1]
	filename := strings.Replace(filenamewithext, ".txt", "", 1)

	strResult := ReadTextFromFile(filenamewithext)

	strResult = MakePrettyTexts(strResult)

	m := MakeMobiMetadata(filename)

	content := []byte(strResult)

	// It has only 1 chapter.
	m.NewChapter(filename, content)

	// Output MOBI File
	m.Write()
}
