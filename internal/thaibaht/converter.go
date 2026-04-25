package thaibaht

import (
	"strings"

	"github.com/shopspring/decimal"
)

var digits = []string{
	"", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า",
}

var positions = []string{
	"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน", "ล้าน",
}

// convertGroup converts a number 0–999999 into Thai text.
// It handles recursive million grouping so very large numbers work correctly.
func convertGroup(n int64) string {
	if n == 0 {
		return ""
	}

	// Handle millions recursively
	if n >= 1_000_000 {
		millions := n / 1_000_000
		remainder := n % 1_000_000
		result := convertGroup(millions) + "ล้าน"
		if remainder > 0 {
			result += convertGroup(remainder)
		}
		return result
	}

	var parts []string
	place := 0
	tmp := n

	// Build digit-position pairs from least significant
	var digitPlaces [][2]int64
	for tmp > 0 {
		digitPlaces = append(digitPlaces, [2]int64{tmp % 10, int64(place)})
		tmp /= 10
		place++
	}

	// Reverse so most significant comes first
	for i, j := 0, len(digitPlaces)-1; i < j; i, j = i+1, j-1 {
		digitPlaces[i], digitPlaces[j] = digitPlaces[j], digitPlaces[i]
	}

	for _, dp := range digitPlaces {
		d, p := dp[0], dp[1]
		if d == 0 {
			continue
		}
		// Special case: "สิบเอ็ด" for 1 in tens place when not the leading digit
		if p == 1 && d == 1 {
			parts = append(parts, "สิบ")
			continue
		}
		// Special case: "เอ็ด" for 1 in tens place (11, 21, 31…)
		if p == 0 && d == 1 && n > 10 {
			parts = append(parts, "เอ็ด")
			continue
		}
		parts = append(parts, digits[d]+positions[p])
	}

	return strings.Join(parts, "")
}

// Convert takes a decimal value and returns Thai baht text.
// Example: 1234 → "หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน"
func Convert(amount decimal.Decimal) string {
	// Split into integer and fractional parts
	intPart := amount.Floor()
	fracPart := amount.Sub(intPart).Mul(decimal.NewFromInt(100)).Round(0)

	intVal := intPart.IntPart()
	fracVal := fracPart.IntPart()

	var sb strings.Builder

	if intVal == 0 {
		sb.WriteString("ศูนย์")
	} else {
		sb.WriteString(convertGroup(intVal))
	}
	sb.WriteString("บาท")

	if fracVal == 0 {
		sb.WriteString("ถ้วน")
	} else {
		sb.WriteString(convertGroup(fracVal))
		sb.WriteString("สตางค์")
	}

	return sb.String()
}
