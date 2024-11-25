package prose

import "strings"

// Having automated tests is like having your code inspected for bugs automatically every time you make a change!

func JoinWithCommas(phrases []string) string {
	if len(phrases) == 1 {
		return phrases[0]
	}
	result := strings.Join(phrases[:len(phrases)-1], ", ")
	if len(phrases) > 2 {
		result += ","
	}
	result += " and "
	result += phrases[len(phrases)-1]
	return result
}
