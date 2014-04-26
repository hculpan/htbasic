package scanner

import (
	"bufio"
	"os"
)

type Source struct {
	filename string
	currLine int
	lines    []string
}

func (source *Source) Open(filename string) error {
	lines, err := readLines(filename)
	if err != nil {
		return err
	}
	source.filename = filename
	source.currLine = 0
	source.lines = lines

	return nil
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func (source *Source) CurrentLine() string {
	return source.lines[source.currLine]
}

func (source *Source) NextLine() string {
	if source.currLine >= len(source.lines) {
		return ""
	} else {
		result := source.CurrentLine()
		source.currLine++
		return result
	}
}
