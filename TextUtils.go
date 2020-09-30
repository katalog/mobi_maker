package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"

	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

// ReadTextFromFile : text -> string with enc conv
func ReadTextFromFile(filenamewithext string) string {
	// See exth.go for additional EXTH record IDs
	content, err := ioutil.ReadFile(filenamewithext)
	if err != nil {
		log.Fatal(err)
	}

	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest(content)
	if err == nil {
		fmt.Printf("%s = charset is %s, language is %s\n", filenamewithext, result.Charset, result.Language)
	} else {
		panic(err)
	}

	var bufs bytes.Buffer

	// convert euckr -> utf8
	if result.Language == "ko" {
		wr := transform.NewWriter(&bufs, korean.EUCKR.NewDecoder())
		wr.Write([]byte(content))
		wr.Close()

		// check convert result
		result, err = detector.DetectBest(bufs.Bytes())
		if err == nil {
			fmt.Printf("Please check again %s mobi file!!!\n", filenamewithext)
			fmt.Printf("Convert Successfully.\n")
		} else {
			panic(err)
		}

		// debug
		//ioutil.WriteFile(filenamewithext+"11", bufs.Bytes(), os.FileMode(0666))
		//fmt.Println(bufs.String())

		return bufs.String()
	}

	content, err = ioutil.ReadFile(filenamewithext)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

// MakePrettyTexts : input newline if text is too long. and remove too many whitespace characters.
func MakePrettyTexts(s2Strings string) string {
	s2Strings = regexp.MustCompile("[\t\v\f\r]").ReplaceAllString(s2Strings, "")
	s2Strings = regexp.MustCompile("[\\pC]").ReplaceAllString(s2Strings, "\n")
	s2Strings = regexp.MustCompile("\"[\\s]+(\\pL)").ReplaceAllString(s2Strings, "\"\n$1")
	s2Strings = regexp.MustCompile("][\\s]+(\\pL)").ReplaceAllString(s2Strings, "]\n$1")
	s2Strings = regexp.MustCompile("”[\\s]+(\\pL)").ReplaceAllString(s2Strings, "”\n$1")
	s2Strings = regexp.MustCompile("\\.[\\s]+\"").ReplaceAllString(s2Strings, ".\n\"")
	s2Strings = regexp.MustCompile("\\.[\\s]+\\[").ReplaceAllString(s2Strings, ".\n[")
	s2Strings = regexp.MustCompile("\\.[\\s]+”").ReplaceAllString(s2Strings, ".\n”")
	s2Strings = regexp.MustCompile("\\. ").ReplaceAllString(s2Strings, ".\n")
	s2Strings = regexp.MustCompile("\n\\s+").ReplaceAllString(s2Strings, "\n")

	// Convert /n -> <br>
	s2Strings = "<br><br>" + s2Strings
	s2Strings = regexp.MustCompile("\n").ReplaceAllString(s2Strings, "<br><br>")

	return s2Strings
}
