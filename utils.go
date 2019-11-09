package parsembox

// isLetter validates that the rune is a letter.
func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// isDigit validates that the rune is a number.
func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

// isWhitespace validates if the rune is a space, newline, or tab.
func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

// isSpace validates if the rune is a space.
func isSpace(ch rune) bool {
	return ch == ' '
}

func isPunctuation(ch rune) bool {
	p := map[rune]bool{
		'!': true, '@': true, '#': true, '$': true,
		'%': true, '^': true, '&': true, '*': true,
		'(': true, ')': true, '_': true, '+': true,
		'-': true, '=': true, '[': true, ']': true,
		'\\': true, '{': true, '}': true, '|': true,
		':': true, ';': true, '"': true, '\'': true,
		'<': true, '>': true, ',': true, '.': true,
		'/': true, '?': true, '`': true, '~': true,
	}
	return p[ch]
}

// isNewline validates if the rune is a newline.
func isNewline(ch rune) bool { return ch == '\n' }

// isColon validates that the rune is a ':'.
func isColon(ch rune) bool { return (ch == colon) }

// isOctothorpe validates that the current rune is a '#'.
func isOctothorpe(ch rune) bool { return (ch == octothorpe) }

// consumeWS consumes leading whitespace from the reader.
func (p *Parser) consumeWS() error {
	for {
		if !isWhitespace(p.Peek()) {
			return nil
		}
		_, _, err := p.Read()
		if err != nil {
			return err
		}
	}
}

// consumeToNewline reads all content until a newline rune is encountered.
func (p *Parser) consumeToNewline() error {
	for {
		if isNewline(p.Peek()) {
			_, _, err := p.Read()
			if err != nil {
				return err
			}
			return nil
		}
		_, _, err := p.Read()
		if err != nil {
			return err
		}
	}
}
