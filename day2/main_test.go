package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMove(t *testing.T) {
	keypad := generate_standard_keypad(1)
	assert.Equal(t, keypad.move([]byte{'U'}).number, byte('1'), "Moving up from top position shouldn't move")
	keypad = generate_standard_keypad(4)
	assert.Equal(t, keypad.move([]byte{'U'}).number, byte('1'), "Moving up from non-top position should move")
}

type TestInput struct {
	directions []byte
	expect     byte
}

func TestInstruction(t *testing.T) {
	inputs := []TestInput{
		{
			directions: []byte("ULL"),
			expect:     '1',
		},
		{
			directions: []byte("RRDDD"),
			expect:     '9',
		},
		{
			directions: []byte("LURDL"),
			expect:     '8',
		},
		{
			directions: []byte("UUUUD"),
			expect:     '5',
		},
	}

	keypad := generate_standard_keypad(5)
	for _, input := range inputs {
		keypad = keypad.move(input.directions)
		assert.Equal(t, keypad.number, input.expect, "keypad.move(%q)", input.directions)
	}
}

func TestDiamondKeymap(t *testing.T) {
	inputs := []TestInput{
		{
			directions: []byte("ULL"),
			expect:     '5',
		},
		{
			directions: []byte("RRDDD"),
			expect:     'D',
		},
		{
			directions: []byte("LURDL"),
			expect:     'B',
		},
		{
			directions: []byte("UUUUD"),
			expect:     '3',
		},
	}

	keypad := generate_diamond_keypad(5)
	for _, input := range inputs {
		keypad = keypad.move(input.directions)
		assert.Equal(t, keypad.number, input.expect, "keypad.move(%q)", input.directions)
	}
}
