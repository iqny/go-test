package reverse_string

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		reverse string
	}{
		{"#1", "123456789", "9876543210"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, _ := reverseString(test.s)
			if got != test.reverse {
				t.Errorf("%s ,got=%v", test.s, got)
			}
		})
	}
}
