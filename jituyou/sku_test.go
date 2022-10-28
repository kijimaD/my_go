package sku

import (
	"fmt"
	"testing"
)

func TestSKU(t *testing.T) {
	skuCD := SKUCode("00001MMBL")

	tests := []struct {
		input          string
		expectedString string
	}{
		{skuCD.ItemCD(), "00001"},
		{skuCD.SizeCD(), "MM"},
		{skuCD.ColorCD(), "BL"},
	}

	for _, tt := range tests {
		if tt.input != tt.expectedString {
			fmt.Println("is not match")
		}
	}
}
