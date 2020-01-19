package counter

import (
	"bufio"
	"strings"
)

// WordCounter returns the seperated words count.
type WordCounter int

// LineCounter retunrs the line numbers of given input.
type LineCounter int

func (w *WordCounter) Write(p []byte) (n int, err error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		*w++
	}

	return int(*w), nil
}

func (l *LineCounter) Write(p []byte) (int, error) {

	scanner := bufio.NewScanner(strings.NewReader(string(p)))

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		*l++
	}

	return int(*l), nil

}
