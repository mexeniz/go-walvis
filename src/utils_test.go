package walvis

import (
	"fmt"
	"testing")

func TestListCsvFiles(t *testing.T) {
    csvFiles := ListCsvFiles("/workspace/data")
    fmt.Println(csvFiles)
}