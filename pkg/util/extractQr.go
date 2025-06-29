package util

import (
	"errors"
	"fmt"
	"strconv"
)

type TLV struct {
	Tag   string `json:"tag"`
	Len   int    `json:"length"`
	Value string `json:"value"`
	Child []TLV  `json:"child,omitempty"`
}

func ParseQR(data string) ([]TLV, error) {
	var result []TLV
	for i := 0; i < len(data); {
		if i+4 > len(data) {
			return nil, errors.New("invalid QRIS content: incomplete TLV header")
		}

		tag := data[i : i+2]
		lengthStr := data[i+2 : i+4]
		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			return nil, fmt.Errorf("invalid length at pos %d: %v", i, err)
		}
		if i+4+length > len(data) {
			return nil, fmt.Errorf("invalid value length at tag %s", tag)
		}

		value := data[i+4 : i+4+length]
		child := []TLV{}

		// Tags that may contain nested TLVs
		if tag == "26" || tag == "51" || tag == "62" || tag == "27" {
			child, err = ParseQR(value)
			if err != nil {
				return nil, fmt.Errorf("error parsing nested TLV at tag %s: %v", tag, err)
			}
		}

		result = append(result, TLV{
			Tag:   tag,
			Len:   length,
			Value: value,
			Child: child,
		})

		i += 4 + length
	}
	return result, nil
}

func FindTLV(tlvs []TLV, tag string) *TLV {
	for _, tlv := range tlvs {
		if tlv.Tag == tag {
			return &tlv
		}
	}
	return nil
}

func FindSubTLV(tlvs []TLV, tag, subtag string) *TLV {
	tlv := FindTLV(tlvs, tag)
	if tlv != nil {
		return FindTLV(tlv.Child, subtag)
	}
	return nil
}

func PadRight(str string, length int) string {
	for len(str) < length {
		str += " "
	}
	return str
}
