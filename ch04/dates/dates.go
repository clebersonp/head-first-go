package dates

// DaysInWeek is a constant
// "const" keyword followed by the name of the constant and the value of the constant
// The type can be omitted, and it will be inferred from the value being assigned
// Attempting to assign a new value to a constant will result in a compile-time error
// It's much more typical to declare constants at the package level, even though they can also be declared locally
// As with variables and functions, a constant whose name begins with a capital letter is exported
const DaysInWeek = 7

func WeeksToDays(weeks int) int {
	return weeks * DaysInWeek
}

func DaysToWeeks(days int) float64 {
	return float64(days) / float64(DaysInWeek)
}
