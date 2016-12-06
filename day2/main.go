package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func move(position int, line []byte) int {
	for _, character := range line {
		switch character {
		case 'U':
			if position > 3 {
				position -= 3
			}
		case 'D':
			if position < 7 {
				position += 3
			}
		case 'L':
			if position%3 != 1 {
				position--
			}
		case 'R':
			if position%3 != 0 {
				position++
			}
		case '\n':
			// Do nothing
		default:
			panic(fmt.Sprintf("Unrecognized character: %c", character))
		}
	}
	return position
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	position := 5
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		position = move(position, line)
		fmt.Print(position)
	}
	fmt.Print("\n")
}
