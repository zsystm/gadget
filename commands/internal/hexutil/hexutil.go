package hexutil

import "strings"

// StripHexPrefix removes a leading 0x/0X prefix if present.
func StripHexPrefix(input string) string {
	if strings.HasPrefix(input, "0x") || strings.HasPrefix(input, "0X") {
		return input[2:]
	}
	return input
}
