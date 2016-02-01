package replace

import (
	"testing"

	"github.com/arschles/assert"
)

func TestFmtReplacement(t *testing.T) {
	repl := FmtReplacement("abc", "def%s%s", "g", "h")
	assert.Equal(t, repl.Old, "abc", "OLD param")
	assert.Equal(t, repl.New, "defgh", "NEW param")
}

func TestString(t *testing.T) {
	repl := FmtReplacement("abc", "def%s%s", "g", "h")
	origStr := "abcdefg"
	newStr := String(origStr, repl)
	assert.Equal(t, newStr, "defghdefg", "new string")
}
