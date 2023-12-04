package currency

// main - dollar

// dollar to another currency
type CurrencyRate map[string]float64

var Rates = CurrencyRate{
	"uah": 37.6,
	"pln": 3.1,
	"usd": 1,
}

func ExchangeCurrency(currFrom string, currTo string) float64 {
	currFromRate := Rates[currFrom]
	currToRate := Rates[currTo]
	return currFromRate / currToRate
}