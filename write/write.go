package write

import (
	"encoding/csv"
	"os"
)

const (
	ExtensionFile = ".csv"
)

func FileCSV(filePath string, rows [][]string, header []string) error {

	var data [][]string

	data = append(data, header)
	for i := 0; i < len(rows); i++ {
		data = append(data, rows[i])
	}

	file, err := os.Create(filePath + ExtensionFile)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.WriteAll(data)
}
