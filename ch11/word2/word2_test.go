package word2

import (
	"math/rand"
	"testing"
	"time"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		name string
		input string
		want  bool
	}{
		{"1","", true},
		{"2","a", true},
		{"3","aa", true},
		{"4","ab", false},
		{"5","kayak", true},
		{"6","detartrated", true},
		{"7","A man, a plan, a canal: Panama", true},
		{"8","Evil I did dwell; lewd did I live.", true},
		{"9","Able was I ere I saw Elba", true},
		{"10","été", true},
		{"11","Et se resservir, ivresse reste.", true},
		{"12","palindrome", false}, // non-palindrome
		{"13","desserts", false},   // semi-palindrome
	}
	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			if got := IsPalindrome(c.input); got != c.want {
				t.Errorf("IsPalindrome(%q) = %v", c.input, got)
			}
		})
	}
}
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}
