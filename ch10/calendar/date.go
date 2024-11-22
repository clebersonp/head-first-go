package calendar

import "errors"

type Date struct {
	// reduce the visibility of fields to protect them from invalid data
	// This will enforce that the field must be set using a setter method
	year  int
	month int
	day   int
}

// Year is a getter method and return the value of Date's year field
// Using a pointer receiver type for consistency with the setter and others methods.
// Same name as the field (but capitalized so it's exported).
// It's similar to getters and setters methods in java. The difference is the convention names for getters methods.
func (d *Date) Year() int {
	return d.year
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	// Go automatically converts to Date value, so we don't need to use '* operator''
	d.year = year
	return nil
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

func (d *Date) Day() int {
	return d.day
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}
