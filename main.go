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
	// TODO move the panic into ReadCsv
	articles, readCsvError := ReadCsv()
	if readCsvError != nil {
		panic(readCsvError)
	}

	// TODO extract into separate method
	var filteredArticles []*Article
	for _, article := range articles {
		if article.Location == "inbox" {
			filteredArticles = append(filteredArticles, article)
		}
	}

	// TODO move the panic into WriteCSV
	writeCsvError := WriteCsv(filteredArticles)
	if writeCsvError != nil {
		panic(readCsvError)
	}

}

// TODO What is the *. Pointers??
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

	// TODO: remove this output
	for _, article := range articles {
		log.Printf("Title: %s, URL: %s", article.Title, article.URL)
	}

	return articles, nil
}

func WriteCsv(articles []*Article) error {
	resultFile, resultFileError := os.OpenFile("result.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if resultFileError != nil {
		return resultFileError
	}
	defer resultFile.Close()

	if marshalFileError := gocsv.MarshalFile(&articles, resultFile); marshalFileError != nil {
		return marshalFileError
	}

	return nil
}
