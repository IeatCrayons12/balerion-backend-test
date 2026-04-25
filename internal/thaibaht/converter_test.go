package thaibaht

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		input    float64
		expected string
	}{
		// From assignment spec
		{1234, "หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน"},
		{33333.75, "สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์"},

		// Edge cases
		{0, "ศูนย์บาทถ้วน"},
		{1, "หนึ่งบาทถ้วน"},
		{10, "สิบบาทถ้วน"},
		{11, "สิบเอ็ดบาทถ้วน"},
		{21, "ยี่สิบเอ็ดบาทถ้วน"},
		{100, "หนึ่งร้อยบาทถ้วน"},
		{1000000, "หนึ่งล้านบาทถ้วน"},
		{0.25, "ศูนย์บาทยี่สิบห้าสตางค์"},
		{1000000.50, "หนึ่งล้านบาทห้าสิบสตางค์"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			input := decimal.NewFromFloat(tt.input)
			got := Convert(input)
			if got != tt.expected {
				t.Errorf("Convert(%v)\n  got:  %q\n  want: %q", tt.input, got, tt.expected)
			}
		})
	}
}
