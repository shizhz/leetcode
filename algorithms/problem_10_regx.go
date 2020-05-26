package algorithms

import (
	"fmt"
)

type matchType int

const (
	matchAny matchType = iota
	matchPrevious
	matchExactly
	matchFinished
)

type regexMatcher struct {
	pattern             string
	matchResult         []rune
	currentPatternIndex int
	rollbackIndex       int
}

func buildPattern(pattern string) string {
	p := []rune{}

	for _, c := range pattern {
		// Make any continious *'s to one
		if c == '*' && len(p) > 0 && p[len(p)-1] == c {
			continue
		}
		p = append(p, c)
	}

	return string(p)
}

func newMatcher(pattern string) *regexMatcher {
	return &regexMatcher{
		pattern: buildPattern(pattern),
	}
}

func (this *regexMatcher) nextMatchType() (matchType, error) {
	mt, err := this.matchTypeAt(this.currentPatternIndex)
	this.currentPatternIndex += 1

	return mt, err
}

func (this *regexMatcher) matchTypeAt(index int) (matchType, error) {
	if index < 0 {
		return 0, fmt.Errorf("Invalid pattern index: %d ", index)
	}
	if index >= len(this.pattern) {
		return matchFinished, nil
	}
	switch this.pattern[index] {
	case '.':
		return matchAny, nil
	case '*':
		return matchPrevious, nil
	default:
		return matchExactly, nil
	}
}

func isMatch(s string, pattern string) bool {
	for i := 0; i < len(s); i++ {

	}

	return false
}
