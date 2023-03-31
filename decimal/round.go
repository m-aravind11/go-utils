package decimal

import "github.com/shopspring/decimal"

func Round(val float64, places int32) float64 {
	return decimal.NewFromFloat(val).Round(places).InexactFloat64()
}

func RoundUp(val float64, places int32) float64 {
	return decimal.NewFromFloat(val).RoundUp(places).InexactFloat64()
}

func RoundDown(val float64, places int32) float64 {
	return decimal.NewFromFloat(val).RoundDown(places).InexactFloat64()
}
