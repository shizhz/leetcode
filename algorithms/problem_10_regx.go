package algorithms

import (
	"fmt"
)

const asterisk byte = '*'
const dot byte = '.'

func buildPattern(pattern string) string {
	p := []byte{}

	for _, c := range pattern {
		// Make any continious *'s to one
		if byte(c) == asterisk && len(p) > 0 && p[len(p)-1] == byte(c) {
			continue
		}
		p = append(p, byte(c))
	}

	return string(p)
}

// Dynamic Programming problem
//
// Divive the matching problem into sub ones:
// 1. Detect whether the first char of string matches the first char of the pattern
// 2. If the first char matches, then go on and check if substring and sub-pattern match
//
// When to end:
// - either s or p is empty
//
// How to divide substring:
// - Easy, just s[1:]
// How to divide sub-patterns:
// - If the second char of p is not *, then just p[1:]
// - If the second char of p is *, then we got two sub patterns here: p and p[2:]
//
// Use a map to save all intermediate matching results as cache for later use to avoid duplicate matching
//
// Note: The first char of p MUST NOT be asterisk

var regexMatchCache map[string]bool = map[string]bool{}

func match(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	key := fmt.Sprintf("%s-%s", s, p)

	if m, found := regexMatchCache[key]; found {
		// fmt.Printf("Cache hit: %s -> %v\n", key, m)
		return m
	}

	pc := p[0]

	if pc == asterisk {
		fmt.Println("The first char of pattern can not be asterisk")
		regexMatchCache[key] = false
		return false
	}

	asteriskFollows := len(p) > 1 && p[1] == asterisk

	if len(s) == 0 {
		result := asteriskFollows && match(s, p[2:])

		regexMatchCache[key] = result
		return result
	}

	sc := s[0]

	firstCharMatch := sc == pc || pc == dot
	var result bool

	if !firstCharMatch {
		if asteriskFollows {
			result = match(s, p[2:])
		}
	} else {
		if asteriskFollows {
			result = match(s, p[2:]) || match(s[1:], p[2:]) || match(s[1:], p)
		} else {
			result = match(s[1:], p[1:])
		}
	}

	regexMatchCache[key] = result

	return result
}

func isMatch(s string, p string) bool {
	pattern := buildPattern(p)

	return match(s, pattern)
}
