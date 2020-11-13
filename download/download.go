package download

import (
	"errors"
	"io"
	"net/http"
	"os"
)

const (
	StatusError   = "received non 200 response code"
	ExtensionFile = ".csv"
)

func FromURL(url string, filePath string) error {

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New(StatusError)
	}

	if len(filePath) == 0 {
		filePath = "files/file"
	}
	out, err := os.Create(filePath + ExtensionFile)

	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, response.Body)
	return err
}
