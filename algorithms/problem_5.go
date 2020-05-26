package algorithms

type palindromeFinder interface {
	longestPalindrome(s string) string
}

func isPalindrome(s string) bool {
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

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Brute force version, check all the substrings from a given string, and find the longest palinderome from them //
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type bruteforce struct {
}

func (this *bruteforce) longestPalindrome(s string) string {
	var result string

	for subs := range this.substrings(s) {
		if isPalindrome(subs) && len(subs) > len(result) {
			result = subs
		}
	}

	return result
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

//////////////////////////////////
// DP version of implementation //
//////////////////////////////////
type dp struct {
	memo map[int]map[int]string
}

func (this *dp) longestPalindrome(s string) string {
	return this.longestPalindromeBetween(s, 0, len(s))
}

func (this *dp) longestPalindromeBetween(s string, i, j int) string {
	if j <= i {
		return ""
	}
	// fmt.Printf("Checkint %d, %d ", i, j)

	if cache, found := this.findPalindrome(s, i, j); found {
		// fmt.Printf("Cache Hit\n")
		return cache
	}
	// fmt.Printf("Cache Miss\n")

	if isPalindrome(s[i:j]) {
		this.memorize(i, j, s[i:j])
		return s[i:j]
	}

	leftPalindrome, rightPalindrome := this.longestPalindromeBetween(s, i, j-1), this.longestPalindromeBetween(s, i+1, j)

	if len(leftPalindrome) > len(rightPalindrome) {
		this.memorize(i, j, leftPalindrome)
		return leftPalindrome
	}
	this.memorize(i, j, rightPalindrome)
	return rightPalindrome
}

func (this *dp) memorize(i, j int, palindrome string) {
	m, ok := this.memo[i]
	if ok {
		m[j] = palindrome
	} else {
		this.memo[i] = map[int]string{
			j: palindrome,
		}
	}
}

func (this *dp) findPalindrome(s string, i, j int) (string, bool) {
	m, ok := this.memo[i]
	if ok {
		result, ok := m[j]
		return result, ok
	}

	return "", false
}

func longestPalindrome(s string) string {
	// return (&bruteforce{}).longestPalindrome(s)
	// return (&expandCenter{}).longestPalindrome(s)
	return (&dp{map[int]map[int]string{}}).longestPalindrome(s)
}
