package walvis

import (
	"fmt"
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestListCsvFiles(t *testing.T) {
	csvFiles := ListCsvFiles("/workspace/data")
	assert.NotNil(t, csvFiles)
	fmt.Println(csvFiles)
}
