package preprocessor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	expected := ` 4 ifndef   FOUR   4 `

	actual, err := New([]string{"testdata/testlib"}).Process("testdata/test1.h")
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
