package walvis

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

func ListCsvFiles(dirPath string) []string {
	var csvPaths []string
	// List all files at dirPath
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
		return csvPaths
	}

	// Filter only csv files by an extension
	for _, file := range files {
		fileName := file.Name()
		if strings.HasSuffix(fileName, ".csv") {
			csvPath := filepath.Join(dirPath, fileName)
			fmt.Println(csvPath)
			csvPaths = append(csvPaths, csvPath)
		}
	}
	return csvPaths
}
