package unit

import "strconv"

// Prefixes.
const (
	Kilo  = 1e3  // k
	Centi = 1e-2 // c
	Milli = 1e-3 // m
	Micro = 1e-6 // µ
	Nano  = 1e-9 // n
)

func getPrefix(v float64) float64 {
	switch {
	case v >= Kilo:
		return Kilo
	case v >= 0.1 || v == 0:
		return 1
	case v >= Centi:
		return Centi
	case v >= Milli:
		return Milli
	case v >= Micro:
		return Micro
	default:
		return Nano
	}
}

func prefixString(p float64) string {
	switch p {
	case Kilo:
		return "k"
	case Centi:
		return "c"
	case Milli:
		return "m"
	case Micro:
		return "µ"
	case Nano:
		return "n"
	default:
		return ""
	}
}

func formatWithPrefixAndSymbol(v float64, symbol string) string {
	prefix := getPrefix(v)
	return strconv.FormatFloat(v/prefix, 'f', -1, 64) + prefixString(prefix) + symbol
}
