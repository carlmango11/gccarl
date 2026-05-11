package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseIdentifier_Valid(t *testing.T) {
	sc := &Scanner{
		text: " h_1x after",
	}

	iden, err := sc.ParseIdentifier()
	assert.NoError(t, err)
	assert.Equal(t, Identifier("h_1x"), iden)
	assert.Equal(t, 5, sc.i)
}

func TestParseIdentifier_Invalid(t *testing.T) {
	sc := &Scanner{
		text: "1omg",
	}

	iden, err := sc.ParseIdentifier()
	assert.Error(t, err)
	assert.Empty(t, iden)
}

func TestParseNumber_Valid(t *testing.T) {
	sc := &Scanner{
		text: "123",
	}

	n, err := sc.ParseNumber()
	assert.NoError(t, err)
	assert.Equal(t, 123., n)
}
