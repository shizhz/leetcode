package algorithms

import (
	"strings"
)

func zigzag(s string, numRows int) string {
	layout := [][]rune{}
	down := true
	row := 0

	if numRows > len(s) {
		numRows = len(s)
	}

	for i := 0; i < numRows; i++ {
		layout = append(layout, []rune{})
	}

	for _, c := range s {
		r := layout[row]
		if len(r) == 0 {
			r = []rune{c}
		} else {
			r = append(r, c)
		}

		layout[row] = r

		if numRows == 1 {
			continue
		}

		if down {
			if row < numRows-1 {
				row += 1
			} else {
				row -= 1
				down = false
			}
		} else {
			if row > 0 {
				row -= 1
			} else {
				row += 1
				down = true
			}
		}
	}

	rows := []string{}

	for i := 0; i < len(layout); i++ {
		rows = append(rows, string(layout[i]))
	}

	return strings.Join(rows, "")
}
