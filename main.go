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
	// Parse the CSV data into the articles slice. If an error occurs, panic.
	if unmarshalError := gocsv.UnmarshalFile(csvFile, &articles); unmarshalError != nil {
		panic(unmarshalError)
	}

	return articles
}

func GetInboxArticles(articles []*Article) []*Article {
	// Initialize an empty slice to store inbox articles
	var inboxArticles []*Article

	// Iterate through each article in the provided slice.
	for _, article := range articles {
		// Check if the article's Location is equal to inbox
		if article.Location == "inbox" {
			// If the article's location is inbox, add it to the inboxArticles slice
			inboxArticles = append(inboxArticles, article)
		}
	}

	return inboxArticles
}

func WriteCsv(articles []*Article) {
	// Open result.csv for writing; create it if it doesn't exist, or overwrite it if it already exists.
	resultFile, resultFileError := os.OpenFile("result.csv", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)

	// Check for errors when opening or creating the file. If there's an error, panic.
	if resultFileError != nil {
		panic(resultFileError)
	}
	defer resultFile.Close()

	// Marshal the articles into the CSV format and write them to the result.csv file
	if marshalFileError := gocsv.MarshalFile(&articles, resultFile); marshalFileError != nil {
		panic(marshalFileError)
	}
}
