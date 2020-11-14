# Easy CSV (Writer, Reader, Downloader, Converter) CLI

### Important warning
- Project under construction
- Still needs improvement

### About
Golang CLI tool to manipule (write, read, download, *convert) CSV file.

Here is a list of the available commands:

* **action**                  - Action to manipulate CSV file [write | read | download]
* **url**                     - URL to download CSV - example: 'http://url.com/file.csv'
* **download-path**           - File path with file name to save CSV file - example: 'files/file' 
* **read-path**               - File path to read CSV file - example: 'files/file.csv'

## Usage

- Download and save in specific folder
```bash
go run main.go --action download --url "http://url.com/file.csv" --download-path "files/file"
```

- Write
```bash
go run main.go --action write
```

- Read and print CSV File
```bash
go run main.go --action read --read-path "files/file.csv"
```
### TODO:

* Create converter CSV to JSON
* Create decent CLI interaction and flow
* Performance tests
