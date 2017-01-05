package palindrome

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

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

func randomNonPalindrome(rng *rand.Rand) string {
	for {
		n := rng.Intn(25) // random length up to 24
		runes := make([]rune, n)
		for i := 0; i < n; i++ {
			r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
			runes[i] = r
		}

		palindrome := true
		var letters []rune
		for _, r := range runes {
			if unicode.IsLetter(r) {
				letters = append(letters, unicode.ToLower(r))
			}
		}
		for i := range letters {
			if letters[i] != letters[len(letters)-1-i] {
				palindrome = false
				break
			}
		}

		if !palindrome {
			return string(runes)
		}
	}
}

func TestRandomNonPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomNonPalindrome(rng)
		if IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true", p)
			for _, v := range p {
				t.Errorf("%v", v)
			}
		}
	}
}
