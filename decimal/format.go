package decimal

import (
	"strconv"
)

func RoundAndFormat(val interface{}, places int32) (string, error) {
	rounded, err := Round(val, places)
	if err != nil {
		return "", err
	}
	return strconv.FormatFloat(rounded, 'f', int(places), 64), nil
}

func RoundUpAndFormat(val interface{}, places int32) (string, error) {
	rounded, err := RoundUp(val, places)
	if err != nil {
		return "", err
	}
	return strconv.FormatFloat(rounded, 'f', int(places), 64), nil
}

func RoundDownAndFormat(val interface{}, places int32) (string, error) {
	rounded, err := RoundDown(val, places)
	if err != nil {
		return "", err
	}
	return strconv.FormatFloat(rounded, 'f', int(places), 64), nil
}
