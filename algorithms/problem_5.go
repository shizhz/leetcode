package algorithms

type palindromeFinder interface {
	longestPalindrome(s string) string
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Brute force version, check all the substrings from a given string, and find the longest palinderome from them //
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type bruteforce struct {
}

func (this *bruteforce) longestPalindrome(s string) string {
	var result string

	for subs := range this.substrings(s) {
		if this.isPalindrome(subs) && len(subs) > len(result) {
			result = subs
		}
	}

	return result
}

func (this *bruteforce) isPalindrome(s string) bool {
	for i := 0; i < len(s); i++ {
		j := len(s) - i - 1
		if j <= i {
			return true
		}

		if s[i] != s[j] {
			return false
		}
	}

	// We treat empty string as true
	return true
}

func (this *bruteforce) substrings(s string) chan string {
	result := make(chan string)

	go func() {
		defer close(result)

		for i := 0; i < len(s); i++ {
			for j := i + 1; j <= len(s); j++ {
				result <- s[i:j]
			}
		}
	}()

	return result
}

////////////////////////////////////////////////////////////////////////////
// Find all palindrome from all possible index and select the longest one //
////////////////////////////////////////////////////////////////////////////
type expandCenter struct {
}

func (this *expandCenter) extractPalindrome(s string, j, k int) string {

	for {
		if j < 0 || k >= len(s) || s[j] != s[k] {
			j = j + 1
			break
		}
		j, k = j-1, k+1
	}
	if j < 0 && k >= len(s) {
		return s
	}
	if j < 0 {
		return s[:k]
	}
	if k >= len(s) {
		return s[j:]
	}

	return s[j:k]
}

func (this *expandCenter) longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	var result string = s[:1] // Always use the first letter as the default palindrome

	for i := 0; i < len(s); i++ {
		if p := this.extractPalindrome(s, i, i); len(p) > len(result) {
			result = p
		}

		if p := this.extractPalindrome(s, i, i+1); len(p) > len(result) {
			result = p
		}
	}

	return result
}

func longestPalindrome(s string) string {
	// return (&bruteforce{}).longestPalindrome(s)
	return (&expandCenter{}).longestPalindrome(s)
}
