package unique_string

import "testing"

func TestIsUniqueString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		b    bool
	}{
		{"#1", "abc", true},
		{"#2", "sdfsdrewrwer", false},
		{"#3", "uirjflsm", true},
		{"#4", "uirjflsm", true},
		{"#5", "​123", true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := isUniqueString(test.s); got != test.b {
				t.Errorf("%s =%v, want %v", test.s, got, test.b)
			}
		})

	}
}
