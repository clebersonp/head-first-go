package prose

import "strings"

// Having automated tests is like having your code inspected for bugs automatically every time you make a change!

func JoinWithCommas(phrases []string) string {
	result := strings.Join(phrases[:len(phrases)-1], ", ")
	result += ", and "
	result += phrases[len(phrases)-1]
	return result
}
