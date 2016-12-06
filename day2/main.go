package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Keypad struct {
	number byte
	steps  [4]*Keypad
}

func generate_standard_keypad(position int) *Keypad {
	var entries [9]Keypad
	for i := range entries {
		entry := &entries[i]
		entry.number = byte(i) + '1'
		for j := range entry.steps {
			entry.steps[j] = entry
		}
		if i >= 3 {
			entry.steps[0] = &entries[i-3]
		}
		if i < len(entries)-3 {
			entry.steps[1] = &entries[i+3]
		}
		if i%3 != 0 {
			entry.steps[2] = &entries[i-1]
		}
		if i%3 != 2 {
			entry.steps[3] = &entries[i+1]
		}
	}
	return &entries[position-1]
}

func (k *Keypad) move_by_direction(direction byte) *Keypad {
	switch direction {
	case 'U':
		return k.steps[0]
	case 'D':
		return k.steps[1]
	case 'L':
		return k.steps[2]
	case 'R':
		return k.steps[3]
	}
	return nil
}

func (k *Keypad) move(line []byte) *Keypad {
	current := k
	for _, character := range line {
		current = current.move_by_direction(character)
	}
	return current
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	keypad := generate_standard_keypad(5)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		keypad = keypad.move(line)
		fmt.Print(keypad.number)
	}
	fmt.Print("\n")
}
