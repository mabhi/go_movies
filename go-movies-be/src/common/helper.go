package common

import (
	"encoding/csv"
	"os"
	"time"
)

// ReadCsv accepts a file and returns its content as a multi-dimentional type
// with lines and each column. Only parses to string type.
func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

//Utility to parse time stamps as per the supplied layout format
func ParseTime(ts, layout string) (time.Time, error) {
	return time.Parse(layout, ts)
}
