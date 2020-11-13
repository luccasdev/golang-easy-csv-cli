package main

import (
	"bufio"
	"easy-csv-go/download"
	"easy-csv-go/read"
	"easy-csv-go/utils"
	"easy-csv-go/write"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	action := flag.String("action", "", "Action [write,read,download] (Required)")
	url := flag.String("url", "", "Url to download CSV - example: 'http://url.com/file.csv'")
	downloadFilePath := flag.String("download-path", "", "File path with file name to save CSV file - example: 'files/file' ")
	readFilePath := flag.String("read-path", "", "File path to read CSV file - example: 'files/file.csv' ")

	flag.Parse()

	if *action == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	switch *action {
	case "write":
		writeCSV()
	case "read":
		csvLines, err := read.Lines(*readFilePath)
		utils.CheckError("Error on read file.", err)
		log.Println(csvLines)
	case "download":
		if *url == "" {
			flag.PrintDefaults()
			utils.InvalidAction("URL cannot be empty.")
		}
		err := download.FromURL(*url, *downloadFilePath)
		utils.CheckError("Error on download file.", err)
	default:
		flag.PrintDefaults()
		utils.InvalidAction("Invalid command.")
		os.Exit(1)
	}

}

func writeCSV() {
	var numberOfRows int
	var numberOfColumns int
	scanner := bufio.NewScanner(os.Stdin)

	println("Enter The Number Of Columns")
	fmt.Scan(&numberOfColumns)

	println("Enter The Number Of Rows")
	fmt.Scan(&numberOfRows)

	var header []string

	rows := make([][]string, 0)

	for i := 0; i < numberOfRows; i++ {
		tmp := make([]string, 0)
		for j := 0; j < numberOfColumns; j++ {
			tmp = append(tmp, "")
		}
		rows = append(rows, tmp)
	}

	for i := 0; i < numberOfColumns; i++ {
		var headerInput string

		fmt.Printf("Enter The %dº Header : ", i+1)

		for scanner.Scan() {
			headerInput = scanner.Text()
			break
		}
		header = append(header, headerInput)
	}

	for j := 0; j < numberOfRows; j++ {
		for i := 0; i < numberOfColumns; i++ {
			fmt.Printf("Enter The Cell of Line %dº and Column %dº : ", j+1, i+1)
			for scanner.Scan() {
				rows[j][i] = scanner.Text()
				break
			}
		}
	}
	err := write.FileCSV("files/sample-write", rows, header)
	utils.CheckError("Error on download file.", err)

}
