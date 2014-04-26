package scanner

import (
	"fmt"
)

type Scanner struct {
	source *Source

	currPosition int

	currValue string
}

func OpenScanner(filename string) Scanner {
	var result Scanner
	result.source = new(Source)
	err := result.source.Open(filename)
	if err != nil {
		fmt.Printf("ERROR: %s", err.Error)
	}
	return result
}

func (scanner *Scanner) nextChar() string {
	if scanner.currPosition >= len(scanner.source.CurrentLine()) {
		return ""
	} else {
		var result string = string(scanner.CurrentLine()[scanner.currPosition])
		scanner.currPosition++
		return result
	}
}

func (scanner *Scanner) currChar() string {
	if scanner.currPosition >= len(scanner.source.CurrentLine()) {
		return ""
	} else {
		var result string = string(scanner.CurrentLine()[scanner.currPosition])
		return result
	}
}

func (scanner *Scanner) backChar() string {
	if scanner.currPosition == 0 && scanner.source.currLine == 0 {
		return ""
	} else if scanner.currPosition == 0 {
		scanner.source.currLine--
		scanner.currPosition = len(scanner.CurrentLine()) - 1
		return scanner.currChar()
	} else {
		scanner.currPosition--
		return scanner.currChar()
	}
}

func (scanner *Scanner) NextTokenValue() string {
	var result string
	inToken := true
	var currValue string = ""

	for inToken {
		c := scanner.nextChar()

		if c == "" {
			if currValue != "" {
				result = currValue
				break
			} else {
				result = "\n"
				scanner.source.NextLine()
				scanner.currPosition = 0
				break
			}
		} else if c == " " || c == "\t" || c == "\n" {
			if currValue != "" {
				result = currValue
				break
			}
		} else if c == "+" || c == "-" || c == "*" || c == "/" {
			if currValue != "" {
				result = currValue
				scanner.backChar()
				break
			} else {
				result = c
				break
			}
		} else {
			currValue += c
		}
	}

	return result
}

func (scanner *Scanner) SourceLineCount() int {
	if scanner.source != nil {
		return len(scanner.source.lines)
	} else {
		return 0
	}
}

func (scanner *Scanner) NextLine() string {
	return scanner.source.NextLine()
}

func (scanner *Scanner) CurrentLine() string {
	return scanner.source.CurrentLine()
}

func (*Token) String() {
	fmt.Printf("token:\n")
}

func NextToken() *Token {
	return new(Token)
}
