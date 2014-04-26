package scanner

type Token struct {
	value string

	lineNumber   int
	linePosition int
	source       string
}
