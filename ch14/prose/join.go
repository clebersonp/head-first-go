package prose

import "strings"

// Having automated tests is like having your code inspected for bugs automatically every time you make a change!

func JoinWithCommas(phrases []string) string {
	if len(phrases) == 0 {
		return ""
	}
	if len(phrases) == 1 {
		return phrases[0]
	}
	if len(phrases) == 2 {
		return phrases[0] + " and " + phrases[1]
	}
	result := strings.Join(phrases[:len(phrases)-1], ", ")
	result += ", and "
	result += phrases[len(phrases)-1]
	return result
}
