// Package sanitized_anchor_name provides a func to create sanitized anchor names.
package sanitized_anchor_name

import (
	"fmt"
	"unicode"
)

// Create returns a sanitized anchor name for the given text.
func Create(text string) string {
	var anchorName []rune
	for _, r := range []rune(text) {
		switch {
		case r == ' ':
			anchorName = append(anchorName, '-')
		case unicode.IsLetter(r) || unicode.IsNumber(r):
			anchorName = append(anchorName, unicode.ToLower(r))
		}
	}
	return string(anchorName)
}

// A Context specifies the supporting context for a generating anchor names.
// All generated anchor names within the same context are guaranteed to be unique.
type Context struct {
	m map[string]uint
}

func NewContext() *Context {
	return &Context{m: make(map[string]uint)}
}

func (ctxt *Context) Create(text string) string {
	// TODO: Clean up this code.
	var anchorName string

	originalAnchorName := Create(text)
	anchorName = originalAnchorName

TryAgain:
	if ctxt.m[originalAnchorName] != 0 {
		anchorName = originalAnchorName + fmt.Sprintf("-%d", ctxt.m[originalAnchorName]+1)
	}

	if ctxt.m[anchorName] != 0 {
		ctxt.m[originalAnchorName] = ctxt.m[originalAnchorName] + 1
		goto TryAgain
	} else {
		ctxt.m[originalAnchorName] = ctxt.m[originalAnchorName] + 1
	}
	if anchorName != originalAnchorName {
		ctxt.m[anchorName] = ctxt.m[anchorName] + 1
	}

	return anchorName
}
