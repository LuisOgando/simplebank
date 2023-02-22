package util

const (
	USD = "USD"
	DOP = "DOP"
	CAD = "CAD"
	EUR = "EUR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, DOP, CAD, EUR:
		return true
	}
	return false
}
