// Package parsembox will parse an mbox stream, returning individual messages.
package parsembox

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"
)

/*
 * parsembox reads a file stream, parsing individual messages
 * from the stream of unix mbox formatted text.
 * TODO(morrowc): more docs as more fleshed out content appears.
 */

const (
	msgStart   = "From "
	eof        = rune(0)
	octothorpe = rune('#')
	colon      = rune(':')
)

// Parser is the struct to manage access to the file/message stream.
type Parser struct {
	reader   *bufio.Reader
	msgCount int
}

// NewParser creates a new parser struct.
func NewParser(fd io.Reader) *Parser {
	return &Parser{
		reader:   bufio.NewReader(fd),
		msgCount: 0,
	}
}

func (p *Parser) String() string {
	b := make([]byte, 1024)
	l, _ := p.reader.Read(b)
	return string(b[:l])
}

// Read exports the ReadRune() method from bufio.Reader, used to read a single rune at a time.
func (p *Parser) Read() (rune, int, error) {
	return p.reader.ReadRune()
}

// Unread reverses the stream one rune only, by exporting bufio.Reader's UnreadRune().
func (p *Parser) Unread() error {
	return p.reader.UnreadRune()
}

// Peek peeks ahead one rune in the Reader.
func (p *Parser) Peek() rune {
	ch, _, err := p.Read()
	// Any error during Peek, which is reading an open file, is EOF, return eof.
	if err != nil {
		return eof
	}
	err = p.Unread()
	if err != nil {
		// Potentially logging the error here would be helpful.
		return eof
	}
	return ch
}

// FindFrom finds the start of an mbox message, returning the From address.
// Leave the read pointer at the newline before the messages headers.
// TODO(morrowc): This function seems like it could be recursive. Make it so.
func (p *Parser) FindFrom() (string, string, error) {
	var from, date bytes.Buffer
	// Start by consuming all leading whitespace.
	err := p.consumeWS()
	if err != nil {
		fmt.Printf("failed during consuming whitespace: %v\n", err)
		return "", "", err
	}

	// Next read chars until there are 5 chars: From<space>
	for {
		ch, _, err := p.Read()
		if err != nil {
			fmt.Printf("failed during attempt to find From<space>: %v\n", err)
			return "", "", err
		}
		if isLetter(ch) && ch == 'F' {
			ch, _, err := p.Read()
			if err != nil {
				fmt.Printf("failed to read letter after F: %v\n", err)
				return "", "", err
			}
			if isLetter(ch) && ch == 'r' {
				ch, _, err := p.Read()
				if err != nil {
					fmt.Printf("failed to read letter after r: %v\n", err)
					return "", "", err
				}
				if isLetter(ch) && ch == 'o' {
					ch, _, err := p.Read()
					if err != nil {
						fmt.Printf("failed to read letter after o: %v\n", err)
						return "", "", err
					}
					if isLetter(ch) && ch == 'm' {
						ch, _, err := p.Read()
						if err != nil {
							fmt.Printf("failed to read letter after m: %v\n", err)
							return "", "", err
						}
						if isSpace(ch) {
							// Read til the next whitespace char, storing in from as the address.
							for {
								ch, _, err := p.Read()
								if err != nil {
									fmt.Printf("read address failed, got(%v): %v\n", from.String(), err)
									return "", "", err
								}
								_, _ = from.WriteRune(ch)
								if isWhitespace(p.Peek()) {
									break
								}
							}

							// Read til a newline, store all data as date.
							// TODO(morrowc): decide if failing is appropriate if there is
							// no date data to return.
							for {
								ch, _, err := p.Read()
								if err != nil {
									fmt.Printf("read date failed, got(%v): %v\n", date.String(), err)
									return "", "", err
								}
								_, _ = date.WriteRune(ch)
								if isNewline(p.Peek()) {
									return from.String(), date.String(), nil
								}
							}
						}
					}
				}
			}
		}
	}
	return "", "", nil
}

// Next returns the next message in the mbox stream.
func (p *Parser) Next() (*string, error) {
	// email-from / date
	from, d, err := p.FindFrom()
	if err != nil {
		return nil, err
	}

	dstmp, err := time.Parse("Mon Jan 2 15:04:05 2006", strings.TrimLeft(d, " 	"))
	if err != nil {
		return nil, err
	}

	// If the next char is a newline, consume it and read until the next "From "
	if isNewline(p.Peek()) {
		fmt.Printf("Past From: %v\n", from)
		fmt.Printf("Date: %v\n", dstmp)
	}

	return &from, nil
}
