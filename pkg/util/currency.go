package util

import (
	"strings"

	"github.com/dustin/go-humanize"
)

// FormatRupiah to convert number 10000 to Rp 10.000,-
func FormatRupiah(amount float64) string {
	humanizeValue := humanize.CommafWithDigits(amount, 0)
	stringValue := strings.Replace(humanizeValue, ",", ".", -1)
	return "Rp " + stringValue + ",-"
}

// FormatRupiah to convert number 10000 to Rp 10.000
func FormatRupiahWithoutDecimal(amount float64) string {
	humanizeValue := humanize.CommafWithDigits(amount, 0)
	return "Rp " + strings.Replace(humanizeValue, ",", ".", -1)
}
