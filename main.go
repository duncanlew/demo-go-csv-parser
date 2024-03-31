package main

import (
	"os"

	"github.com/gocarina/gocsv"
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
	articles := ReadCsv()
	filteredArticles := GetInboxArticles(articles)
	WriteCsv(filteredArticles)
}

func ReadCsv() []*Article {
	// Try to open the example.csv file in read-write mode.
	csvFile, csvFileError := os.OpenFile("example.csv", os.O_RDWR, os.ModePerm)
	// If an error occurs during os.OpenFIle, panic and halt execution.
	if csvFileError != nil {
		panic(csvFileError)
	}
	// Ensure the file is closed once the function returns
	defer csvFile.Close()

	var articles []*Article
	// Parse the CSV data into the articles array. If an error occurs, panic.
	if unmarshalError := gocsv.UnmarshalFile(csvFile, &articles); unmarshalError != nil {
		panic(unmarshalError)
	}

	return articles
}

func WriteCsv(articles []*Article) {
	resultFile, resultFileError := os.OpenFile("result.csv", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if resultFileError != nil {
		panic(resultFileError)
	}
	defer resultFile.Close()

	if marshalFileError := gocsv.MarshalFile(&articles, resultFile); marshalFileError != nil {
		panic(marshalFileError)
	}
}

func GetInboxArticles(articles []*Article) []*Article {
	var filteredArticles []*Article
	for _, article := range articles {
		if article.Location == "inbox" {
			filteredArticles = append(filteredArticles, article)
		}
	}
	return filteredArticles
}
