package prose

import (
	"fmt"
	"testing"
)

// To create a file of test cases, the file name must be the suffix _test.go. i.e.: nameOfFile_test.go
// You're not required to make your tests part of the same package as the code you're testing, but if you want to
// access unexported types or function from the package, you'll need to.
// The testing functions need to be the prefix Test followed the name what you want to test ou the function name.
// i.e.: TestMyFunction
// All testing function has a parameter of type *testing.T

// To run test files with Go test tool:
// Run 'go test' followed by the import path of the package that contains your tests.
// go test ./... -v
// or: go test ./ch14/prose -v
// -v is for verbose
// It will run every function contained in those files whose name starts with Test.

func TestJoinWithCommas(t *testing.T) {
	// table driven tests
	// Many testing and one structure
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "One element",
			args: []string{"apple"},
			want: "apple",
		},
		{
			name: "Two elements",
			args: []string{"apple", "banana"},
			want: "apple and banana",
		},
		{
			name: "Three elements",
			args: []string{"apple", "banana", "pineapple"},
			want: "apple, banana, and pineapple",
		},
		{
			name: "Four elements",
			args: []string{"apple", "banana", "pineapple", "grapefruit"},
			want: "apple, banana, pineapple, and grapefruit",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := JoinWithCommas(test.args); got != test.want {
				// call a method on the testing.T to fail the test
				t.Errorf(errorString(test.args, got, test.want))
			}
		})
	}
}

// this is a 'test helper function'.
// Since this function doesn't start with Test, the 'go test tool' will ignore it. So, it's not treated as a test.
func errorString(list []string, got, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = %s, want = %s", list, got, want)
}
