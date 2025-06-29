package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseQRIS(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected []TLV
		wantErr  bool
	}{
		{
			desc:  "valid QRIS content",
			input: "00020101021126660014ID.LINKAJA.WWW011893600911000000000802152103124400000080303UMI51440014ID.CO.QRIS.WWW0215ID20210652077750303UMI5204839853033605802ID5922YAY BAKTI KAMAJAYA IND6006SLEMAN61055528162070703A016304FA4D",
			expected: []TLV{
				{
					Tag:   "00",
					Len:   2,
					Value: "01",
					Child: []TLV{},
				},
				{
					Tag:   "01",
					Len:   2,
					Value: "11",
					Child: []TLV{},
				},
				{
					Tag:   "26",
					Len:   66,
					Value: "0014ID.LINKAJA.WWW011893600911000000000802152103124400000080303UMI",
					Child: []TLV{
						{
							Tag:   "00",
							Len:   14,
							Value: "ID.LINKAJA.WWW",
							Child: []TLV{},
						},
						{
							Tag:   "01",
							Len:   18,
							Value: "936009110000000008",
							Child: []TLV{},
						},
						{
							Tag:   "02",
							Len:   15,
							Value: "210312440000008",
							Child: []TLV{},
						},
						{
							Tag:   "03",
							Len:   03,
							Value: "UMI",
							Child: []TLV{},
						},
					},
				},
				{
					Tag:   "51",
					Len:   44,
					Value: "0014ID.CO.QRIS.WWW0215ID20210652077750303UMI",
					Child: []TLV{
						{
							Tag:   "00",
							Len:   14,
							Value: "ID.CO.QRIS.WWW",
							Child: []TLV{},
						},
						{
							Tag:   "02",
							Len:   15,
							Value: "ID2021065207775",
							Child: []TLV{},
						},
						{
							Tag:   "03",
							Len:   03,
							Value: "UMI",
							Child: []TLV{},
						},
					},
				},
				{
					Tag:   "52",
					Len:   04,
					Value: "8398",
					Child: []TLV{},
				},
				{
					Tag:   "53",
					Len:   03,
					Value: "360",
					Child: []TLV{},
				},
				{
					Tag:   "58",
					Len:   02,
					Value: "ID",
					Child: []TLV{},
				},
				{
					Tag:   "59",
					Len:   22,
					Value: "YAY BAKTI KAMAJAYA IND",
					Child: []TLV{},
				},
				{
					Tag:   "60",
					Len:   06,
					Value: "SLEMAN",
					Child: []TLV{},
				},
				{
					Tag:   "61",
					Len:   05,
					Value: "55281",
					Child: []TLV{},
				},
				{
					Tag:   "62",
					Len:   07,
					Value: "0703A01",
					Child: []TLV{
						{
							Tag:   "07",
							Len:   03,
							Value: "A01",
							Child: []TLV{},
						},
					},
				},
				{
					Tag:   "63",
					Len:   04,
					Value: "FA4D",
					Child: []TLV{},
				},
			},
			wantErr: false,
		},
		{
			desc:     "empty input",
			input:    "",
			expected: nil,
			wantErr:  false,
		},
		{
			desc:    "incomplete TLV header",
			input:   "01",
			wantErr: true,
		},
		{
			desc:    "invalid length format",
			input:   "0101xx",
			wantErr: true,
		},
		{
			desc:    "value length exceeds input",
			input:   "010105A",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result, err := ParseQR(tc.input)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}
