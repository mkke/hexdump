package hexdump

import (
	"fmt"
	"strings"

	"github.com/mkke/color"
)

func toChar(b byte) string {
	if b < 32 || b > 126 {
		return "."
	}
	return string([]byte{b})
}

// Dump2 returns a string that contains a side-by-side hex dump of the two given data slices.
func Dump2(a, b []byte) string {
	sprintf := func(format string, emphasize bool, args ...any) string {
		return fmt.Sprintf(format, args...)
	}
	if !color.NoColor {
		col := color.New(color.FgYellow)
		sprintf = func(format string, emphasize bool, args ...any) string {
			if emphasize {
				return col.Sprintf(format, args...)
			} else {
				return fmt.Sprintf(format, args...)
			}
		}
	}

	sb := strings.Builder{}
	for lineAddr := 0; lineAddr < max(len(a), len(b)); lineAddr += 16 {
		lineAHex := ""
		lineAText := ""
		lineBHex := ""
		lineBText := ""

		for colIdx := range 16 {
			addr := lineAddr + colIdx
			aOk := addr < len(a)
			bOk := addr < len(b)

			equal := aOk && bOk && a[addr] == b[addr]
			if aOk {
				lineAHex += sprintf(" %02x", !equal, a[addr])
				lineAText += sprintf("%s", !equal, toChar(a[addr]))
			} else {
				lineAHex += "   "
				lineAText += " "
			}

			if bOk {
				lineBHex += sprintf(" %02x", !equal, b[addr])
				lineBText += sprintf("%s", !equal, toChar(b[addr]))
			} else {
				lineBHex += "   "
				lineBText += " "
			}
			if colIdx == 7 {
				lineAHex += " "
				lineBHex += " "
			}
		}
		sb.WriteString(fmt.Sprintf("%08x %s |%s |%s| |%s|\n", lineAddr, lineAHex, lineBHex, lineAText, lineBText))
	}

	return sb.String()
}
