package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var floats []float64

	for _, string := range strings {
		floeatPrice, err := strconv.ParseFloat(string, 64)
		if err != nil {
			return nil, errors.New("Converting prices failed")
		}
		floats = append(floats, floeatPrice)
	}

	return floats, nil
}
