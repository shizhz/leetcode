package algorithms

import (
	"fmt"
)

const asterisk byte = '*'
const dot byte = '.'

type matchType int

const (
	matchAny matchType = iota
	matchPreviousRepeatly
	matchExactly
	matchFinished
)

type regexMatcher struct {
	pattern             string
	currentPatternIndex int
	rollbackIndex       int
}

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

func newMatcher(pattern string) (*regexMatcher, error) {
	if len(pattern) == 0 || pattern[0] == asterisk {
		return nil, fmt.Errorf("Invalid pattern:%s", pattern)
	}
	return &regexMatcher{
		pattern:       buildPattern(pattern),
		rollbackIndex: -1,
	}, nil
}

func (this *regexMatcher) currentActualMatchType() matchType {
	return this.matchTypeAt(this.currentPatternIndex)
}

func (this *regexMatcher) currentEffectMatchType() matchType {
	mt := this.matchTypeAt(this.currentPatternIndex)

	if mt == matchPreviousRepeatly {
		return this.matchTypeAt(this.currentPatternIndex - 1)
	}

	return mt
}

func (this *regexMatcher) matchTypeAt(index int) matchType {
	if index >= len(this.pattern) {
		return matchFinished
	}
	switch this.pattern[index] {
	case dot:
		return matchAny
	case asterisk:
		return matchPreviousRepeatly
	default:
		return matchExactly
	}
}

func (this *regexMatcher) setRollbackIndex(index int) {
	this.rollbackIndex = index
}

func (this *regexMatcher) match(c byte) bool {
	var m bool
	switch this.currentEffectMatchType() {
	case matchExactly:
		m = this.pattern[this.currentPatternIndex] == c
	case matchAny:
		m = true
	}

	if this.currentActualMatchType() != matchPreviousRepeatly {
		this.currentPatternIndex += 1
	}

	return m
}

func (this *regexMatcher) isRollbackable() bool {
	return this.rollbackIndex >= 0
}

func isMatch(s string, pattern string) bool {
	matcher, err := newMatcher(pattern)

	if err != nil {
		return false
	}

	for i := 0; i < len(s); {
		if matcher.match(s[i]) {
			// TODO; handle rollback, this is NOT correct
			if matcher.currentActualMatchType() == matchPreviousRepeatly {
				matcher.setRollbackIndex(i + 1)
			}
			i++
		} else if matcher.isRollbackable() {
			matcher.currentPatternIndex += 1
			i = matcher.rollbackIndex
		} else {
			return false
		}
	}

	return true
}
