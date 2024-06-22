package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	items := []string{
		"pacman", "arifo", "teris", "frogger",
		"pacman", "arfio", "terifs", "frogger",
		"pacman", "ario", "terfis", "frogger",
		"pacmfn",
	}

	const pageSize = 6

	l := len(items)
	for from := 0; from < l; from += pageSize {
		to := from + pageSize
		if to > l {
			to = l
		}

	currentPage := items[from:to]
	head :=fmt.Sprintf("page #%d", (from / pageSize))
	s.Show(head, currentPage)
	}

}
