package repository

import (
	"os"
	"encoding/csv"
)

// ReadCsv is temporarily public, later on will be a private function 
// used by the functions in the repository interface
func ReadCsv(filename string) (records [][]string, err error) {
	csvfile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(csvfile)
	data, err := reader.ReadAll()
	return data, err
}


