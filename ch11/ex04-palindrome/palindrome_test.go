package palindrome

import (
	"math/rand"
	"testing"
	"time"
)

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}

	nNonLetter := rng.Intn(5)
	index := []int{}
	if n > 0 {
		for i := 0; i < nNonLetter; i++ {
			index = append(index, rng.Intn(n))
		}
	}

	for _, i := range index {
		// select rune
		var r rune
		switch i % 3 {
		case 0:
			r = ' '
		case 1:
			r = '.'
		case 2:
			r = ','
		}

		// insert
		runes = append(runes[:i], append([]rune{r}, runes[i:]...)...)
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
