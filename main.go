package main

import (
	"github.com/gocarina/gocsv"
	"log"
	"os"
)

type Article struct {
	Title           string `csv:"Title"`
	URL             string `csv:"URL"`
	DocumentTags    string `csv:"Document tags"`
	SavedDate       string `csv:"Saved date"`
	ReadingProgress string `csv:"Reading progress"`
	Location        string `csv:"Location"`
	Seen            string `csv:"Seen"`
}

func main() {
	csvFile, err := os.OpenFile("example.csv", os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer csvFile.Close()

	var articles []*Article

	if err := gocsv.UnmarshalFile(csvFile, &articles); err != nil {
		panic(err)
	}

	for _, article := range articles {
		log.Printf("Title: %s, URL: %s", article.Title, article.URL)
	}

	// TODO: next step, create articlesThatAreUnread
}
