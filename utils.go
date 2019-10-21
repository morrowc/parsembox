package parsembox

// IsLetter validates that the rune is a letter.
func IsLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// IsDigit validates that the rune is a number.
func IsDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

// IsWhitespace validates if the rune is a space, newline, or tab.
func IsWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

// IsSpace validates if the rune is a space.
func IsSpace(ch rune) bool {
	return ch == ' '
}

// IsNewline validates if the rune is a newline.
func IsNewline(ch rune) bool { return ch == '\n' }

// IsColon validates that the rune is a ':'.
func IsColon(ch rune) bool { return (ch == colon) }

// IsOctothorpe validates that the current rune is a '#'.
func IsOctothorpe(ch rune) bool { return (ch == octothorpe) }

// Consume leading whitespace from the reader.
func (p *Parser) ConsumeWS() error {
	for {
		if !IsWhitespace(p.Peek()) {
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
		if IsNewline(p.Peek()) {
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
