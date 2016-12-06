package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMove(t *testing.T) {
	assert.Equal(t, move(1, []byte{'U'}), 1, "Moving up from top position shouldn't move")
	assert.Equal(t, move(4, []byte{'U'}), 1, "Moving up from non-top position should move")
}

func TestInstruction(t *testing.T) {
	type TestInput struct {
		start      int
		directions []byte
		expect     int
	}

	inputs := []TestInput{
		{
			start:      5,
			directions: []byte("ULL"),
			expect:     1,
		},
		{
			start:      1,
			directions: []byte("RRDDD"),
			expect:     9,
		},
		{
			start:      9,
			directions: []byte("LURDL"),
			expect:     8,
		},
		{
			start:      8,
			directions: []byte("UUUUD"),
			expect:     5,
		},
	}

	for _, input := range inputs {
		assert.Equal(t, move(input.start, input.directions), input.expect, "move(%d, %q)", input.start, input.directions)
	}
}
