// Data analysis module
package	walvis

import (
	"os"

	"github.com/go-gota/gota/dataframe"
	log "github.com/sirupsen/logrus"
)

func ReadCSVToDf(csvPath string) dataframe.DataFrame {
	csvfile, err := os.Open(csvPath)
	if err != nil {
		log.Fatal(err)
		return dataframe.DataFrame{}
	}

	log.Debug("read csv to df: csvPath=", csvPath)
	df := dataframe.ReadCSV(csvfile,
		dataframe.WithDelimiter(';'),
		dataframe.HasHeader(true))
	return df
}

// Analytic functions
func CalcTotalExpense(df dataframe.DataFrame) float64 {
	return 0.0
}

func AverageExpensePerCat(df dataframe.DataFrame) dataframe.DataFrame {
	return dataframe.DataFrame{}
}