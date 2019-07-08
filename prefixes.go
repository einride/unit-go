package unit

import (
	"strconv"
	"strings"
	"unicode/utf8"

	"golang.org/x/xerrors"
)

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

func formatPrefix(p float64) string {
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

func parsePrefix(p rune) (float64, bool) {
	switch p {
	case 'k':
		return Kilo, true
	case 'c':
		return Centi, true
	case 'm':
		return Milli, true
	case 'µ':
		return Micro, true
	case 'n':
		return Nano, true
	default:
		return 1, false
	}
}

func parse(str string, units map[string]float64) (float64, error) {
	rest := str
	// parse unit
	var okUnit bool
	var symbol string
	var unit float64
	for symbol, unit = range units {
		if !strings.HasSuffix(rest, symbol) {
			continue
		}
		okUnit = true
		break
	}
	if !okUnit {
		return 0, xerrors.Errorf("parse '%s': unknown unit", str)
	}
	rest = rest[:len(rest)-len(symbol)]
	if len(rest) == 0 {
		return 0, xerrors.Errorf("parse '%s': not a number", str)
	}
	// parse prefix, if any
	lastRune, lastRuneSize := utf8.DecodeLastRuneInString(rest)
	prefix, hasPrefix := parsePrefix(lastRune)
	if hasPrefix {
		rest = rest[:len(rest)-lastRuneSize]
	}
	if len(rest) == 0 {
		return 0, xerrors.Errorf("parse '%s': not a number", str)
	}
	// parse magnitude
	magnitude, err := strconv.ParseFloat(rest, 64)
	if err != nil {
		return 0, xerrors.Errorf("parse '%s': %w", str, err)
	}
	return magnitude * prefix * unit, nil
}

func format(v float64, symbol string) string {
	var sign string
	if v < 0 {
		sign = "-"
		v *= -1
	}
	prefix := getPrefix(v)
	return sign + strconv.FormatFloat(v/prefix, 'g', -1, 64) + formatPrefix(prefix) + symbol
}
