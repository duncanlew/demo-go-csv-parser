package main

import (
	"github.com/gocarina/gocsv"
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
	articles := ReadCsv()
	filteredArticles := GetInboxArticles(articles)
	WriteCsv(filteredArticles)
}

func ReadCsv() []*Article {
	csvFile, csvFileError := os.OpenFile("example.csv", os.O_RDWR, os.ModePerm)
	if csvFileError != nil {
		panic(csvFileError)
	}
	defer csvFile.Close()

	var articles []*Article
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
