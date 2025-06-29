package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatRupiah(t *testing.T) {
	testCases := []struct {
		desc     string
		amount   float64
		expected string
	}{
		{
			desc:     "should format positive amount with thousands separator",
			amount:   10000,
			expected: "Rp 10.000,-",
		},
		{
			desc:     "should format zero amount",
			amount:   0,
			expected: "Rp 0,-",
		},
		{
			desc:     "should format large amount with multiple separators",
			amount:   1000000,
			expected: "Rp 1.000.000,-",
		},
		{
			desc:     "should format small amount without separator",
			amount:   100,
			expected: "Rp 100,-",
		},
		{
			desc:     "should format decimal amount and round down",
			amount:   10000.99,
			expected: "Rp 10.000,-",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := FormatRupiah(tc.amount)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestFormatRupiahWithoutDecimal(t *testing.T) {
	testCases := []struct {
		desc     string
		amount   float64
		expected string
	}{
		{
			desc:     "should format positive amount with thousands separator without decimal",
			amount:   10000,
			expected: "Rp 10.000",
		},
		{
			desc:     "should format zero amount without decimal",
			amount:   0,
			expected: "Rp 0",
		},
		{
			desc:     "should format large amount with multiple separators without decimal",
			amount:   1000000,
			expected: "Rp 1.000.000",
		},
		{
			desc:     "should format small amount without separator and decimal",
			amount:   100,
			expected: "Rp 100",
		},
		{
			desc:     "should format decimal amount and round down without decimal",
			amount:   10000.99,
			expected: "Rp 10.000",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := FormatRupiahWithoutDecimal(tc.amount)
			assert.Equal(t, tc.expected, result)
		})
	}
}