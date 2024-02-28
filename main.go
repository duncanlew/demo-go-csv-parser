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
	articles, error := ReadCsv()
	if error != nil {
		panic(error)
	}

	var filteredArticles []*Article
	for _, article := range articles {
		if article.Location == "inbox" {
			filteredArticles = append(filteredArticles, article)
		}
	}

	resultFile, resultFileError := os.OpenFile("result.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if resultFileError != nil {
		panic(resultFileError)
	}
	defer resultFile.Close()

	if err := gocsv.MarshalFile(&filteredArticles, resultFile); resultFileError != nil {
		panic(err)
	}
}

func ReadCsv() ([]*Article, error) {
	csvFile, csvFileError := os.OpenFile("example.csv", os.O_RDWR, os.ModePerm)
	if csvFileError != nil {
		return nil, csvFileError
	}
	defer csvFile.Close()

	var articles []*Article
	if unmarshalError := gocsv.UnmarshalFile(csvFile, &articles); unmarshalError != nil {
		return nil, unmarshalError
	}

	for _, article := range articles {
		log.Printf("Title: %s, URL: %s", article.Title, article.URL)
	}

	return articles, nil
}
