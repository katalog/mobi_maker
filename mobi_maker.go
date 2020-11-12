package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func convert(filenamewithext string) {
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

func main() {
	//debug
	//os.Args = []string{"kindmaker", "uoop.txt"}

	if len(os.Args) != 2 {
		fmt.Println("USAGE : mobi_maker textfile.txt")
		fmt.Println("or      mobi_maker -a")
		fmt.Println("-a option : convert all txt file in same folder")
		return
	}

	lstFilename := []string{}

	if os.Args[1] == "-a" {
		files, err := ioutil.ReadDir("./")
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			if strings.Contains(f.Name(), ".txt") {
				fmt.Println(f.Name())
				lstFilename = append(lstFilename, f.Name())
			}
		}
	} else {
		filenamewithext := os.Args[1]
		lstFilename = append(lstFilename, filenamewithext)
	}

	for _, fname := range lstFilename {
		convert(fname)
	}
}
