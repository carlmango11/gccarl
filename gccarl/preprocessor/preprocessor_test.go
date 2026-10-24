package preprocessor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	const input = `void main() { // main func
	int x = 1; // comment // another
}`

	const expected = `void main() { 
	int x = 1; 
}`

	actual := Process([]byte(input))
	assert.Equal(t, expected, actual)
}
