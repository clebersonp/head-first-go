package magazine

// Subscriber is the underlying type of struct with these fields
// Subscriber is capitalized and exported. So it can be used from other packages
// If we want to access the fields, we need to export the defined type (struct) and its fields
type Subscriber struct {
	// We can mixe exported and unexported fields.
	// Unexported fields are fields that begin with a lowercase letter.
	// Sometimes, we want to encapsulate data that we don't want to expose to the outside world.
	Name   string  // name is Capitalized and exported
	Rate   float64 // Rate is exported
	Active bool    // Active is exported
	// Adding a struct as a field on another type
	HomeAddress Address
}

type Employee struct {
	Name   string
	Salary float64
	// Adding a struct as a field on another type
	HomeAddress Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
