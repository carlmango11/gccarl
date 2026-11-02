package preprocessor

import (
	"github.com/carlmango11/gccarl/gccarl/tokens"
)

func NewReader(includePaths []string, filename string) (tokens.Reader, error) {
	return newReader(filename, includePaths, false, make(map[string]*Macro))
}
