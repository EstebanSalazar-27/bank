package fileOps

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, filename string) {
	stringValue := fmt.Sprint(value)
	os.WriteFile(filename, []byte(stringValue), 0644)
}

func ReadFloatFromFile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 1000, errors.New("Error to find the  file.")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("Error converting the value to float64")
	}
	return value, nil
}
