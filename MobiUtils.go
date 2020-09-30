package main

import (
	"time"

	"github.com/766b/mobi"
)

// MakeMobiMetadata : make metadata for mobi file
func MakeMobiMetadata(filename string) (writer *mobi.MobiWriter) {
	m, err := mobi.NewWriter(filename + ".mobi")
	if err != nil {
		panic(err)
	}

	m.Title(filename)
	m.Compression(mobi.CompressionNone) // LZ77 compression is also possible using  mobi.CompressionPalmDoc

	// Add cover image
	//m.AddCover("data/cover.jpg", "data/thumbnail.jpg")

	// Meta data
	currentTime := time.Now()
	today := currentTime.Format("2006-01-02")

	m.NewExthRecord(mobi.EXTH_DOCTYPE, "EBOK")
	m.NewExthRecord(mobi.EXTH_AUTHOR, today)

	return m
}
