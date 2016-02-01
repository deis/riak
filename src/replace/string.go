package replace

import (
	"fmt"
	"strings"
)

type Replacement struct {
	Old string
	New string
}

func FmtReplacement(old string, newFmt string, vals ...interface{}) Replacement {
	return Replacement{Old: old, New: fmt.Sprintf(newFmt, vals...)}
}

func String(orig string, repls ...Replacement) string {
	for _, repl := range repls {
		orig = strings.Replace(orig, repl.Old, repl.New, -1)
	}
	return orig
}
