package main

import (
	"easy-csv-go/download"
	"easy-csv-go/read"
	"easy-csv-go/write"
	"testing"
)

func TestDownloadFromURL(t *testing.T) {
	err := download.FromURL("https://www.sample-videos.com/csv/Sample-Spreadsheet-10-rows.csv", "../files/sample-download")

	if err != nil {
		t.Errorf("Error on download file: %d", err)
	}
}

func TestWriteAndSaveFile(t *testing.T) {
	header := []string{"name", "lastname", "city"}

	rows := [][]string{{"Lucas", "Souza", "SP"}, {"Lucas", "Souza", "SP"}}

	err := write.FileCSV("../files/sample-write", rows, header)

	if err != nil {
		t.Errorf("Error on write file: %d", err)
	}
}

func TestReadFile(t *testing.T) {
	csvFile, err := read.File("../files/sample-write.csv")

	if csvFile == nil {
		t.Errorf("Error on read file: %d", err)
	}
}

func TestReadLines(t *testing.T) {
	csvLines, err := read.Lines("../files/sample-write.csv")

	if err != nil {
		t.Errorf("Error on lines file: %d", err)
	}
	for _, line := range csvLines {
		t.Log(line)
	}
}
