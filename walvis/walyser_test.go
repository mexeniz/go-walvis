package walvis

import (
	// "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCSVToDf(t *testing.T) {
	csvFiles := ListCsvFiles("/workspace/data")
	for _, csvPath := range csvFiles {
		var df = ReadCSVToDf(csvPath)
		assert.NotNil(t, df)
		// fmt.Println("Df:", df)
	}
}