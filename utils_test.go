package parsembox

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsLetter(t *testing.T) {
	tests := []struct {
		desc string
		char rune
		want bool
	}{{
		desc: "Is a",
		char: 'a',
		want: true,
	}, {
		desc: "Is A",
		char: 'A',
		want: true,
	}, {
		desc: "Is 1",
		char: '1',
		want: false,
	}, {
		desc: "Is <period>",
		char: '.',
		want: false,
	}}

	for _, test := range tests {
		got := isLetter(test.char)
		if got != test.want {
			t.Errorf("[%v]: failed to got/want comparison got/want: %v/%v", test.desc, got, test.want)
		}
	}
}

func TestIsDigit(t *testing.T) {
	tests := []struct {
		desc string
		char rune
		want bool
	}{{
		desc: "Is 1",
		char: '1',
		want: true,
	}, {
		desc: "Is 0",
		char: '1',
		want: true,
	}, {
		desc: "Is <period>",
		char: '.',
		want: false,
	}, {
		desc: "Is f",
		char: 'f',
		want: false,
	}}

	for _, test := range tests {
		got := isDigit(test.char)
		if got != test.want {
			t.Errorf("[%v]: failed to got/want comparison got/want: %v/%v", test.desc, got, test.want)
		}
	}
}

func TestIsWhitespace(t *testing.T) {
	tests := []struct {
		desc string
		char rune
		want bool
	}{{
		desc: "Is <space>",
		char: ' ',
		want: true,
	}, {
		desc: "Is f",
		char: 'f',
		want: false,
	}, {
		desc: "Is <tab>",
		char: '	',
		want: true,
	}}

	for _, test := range tests {
		got := isWhitespace(test.char)
		if got != test.want {
			t.Errorf("[%v]: failed to got/want comparison got/want: %v/%v", test.desc, got, test.want)
		}
	}
}

func TestIsSpace(t *testing.T) {
	tests := []struct {
		desc string
		char rune
		want bool
	}{{
		desc: "Is <space>",
		char: ' ',
		want: true,
	}, {
		desc: "Is f",
		char: 'f',
		want: false,
	}, {
		desc: "Is <tab>",
		char: '	',
		want: false,
	}, {
		desc: "Is <newline>",
		char: '',
		want: false,
	}}

	for _, test := range tests {
		got := isSpace(test.char)
		if got != test.want {
			t.Errorf("[%v]: failed to got/want comparison got/want: %v/%v", test.desc, got, test.want)
		}
	}
}

func TestIsNewline(t *testing.T) {
	tests := []struct {
		desc string
		char rune
		want bool
	}{{
		desc: "Success",
		char: '\n',
		want: true,
	}, {
		desc: "Fail",
		char: 'n',
		want: false,
	}}

	for _, test := range tests {
		got := isNewline(test.char)
		if got != test.want {
			t.Errorf("[%v]: got/want mismatch, got: %v want: %v", test.desc, got, test.want)
		}
	}
}

func TestIsColon(t *testing.T) {
	tests := []struct {
		desc string
		char rune
		want bool
	}{{
		desc: "Success",
		char: ':',
		want: true,
	}, {
		desc: "Fail",
		char: 'n',
		want: false,
	}}

	for _, test := range tests {
		got := isColon(test.char)
		if got != test.want {
			t.Errorf("[%v]: got/want mismatch, got: %v want: %v", test.desc, got, test.want)
		}
	}
}

func TestIsOctothorpe(t *testing.T) {
	tests := []struct {
		desc string
		char rune
		want bool
	}{{
		desc: "Success",
		char: '#',
		want: true,
	}, {
		desc: "Fail",
		char: 'n',
		want: false,
	}}

	for _, test := range tests {
		got := isOctothorpe(test.char)
		if got != test.want {
			t.Errorf("[%v]: got/want mismatch, got: %v want: %v", test.desc, got, test.want)
		}
	}
}

func TestConsumeWS(t *testing.T) {
	tests := []struct {
		desc    string
		input   string
		wantErr bool
	}{{
		desc:  "Success",
		input: "  spaces before words",
	}, {
		desc:  "Success - with newline and space",
		input: " before words",
	}, {
		desc: "Success - with newlines",
		input: "before words",
	}, {
		desc: "Success - with tabs",
		input: "	before words",
	}, {
		desc: "Success - with no ws",
		input: "i	",
	}, {
		desc:  "Get Err from empty input",
		input: "",
	}}

	ws := []string{" ", "\n", "\t"}

	for _, test := range tests {
		p := NewParser(strings.NewReader(test.input))
		err := p.consumeWS()
		switch {
		case err != nil && !test.wantErr:
			t.Errorf("[%v]: test got error when not expecting one: %v", test.desc, err)
		case err == nil && test.wantErr:
			t.Errorf("[%v]: test did not get error when expecting one", test.desc)
		case err == nil:
			for _, c := range ws {
				if strings.HasPrefix(p.String(), c) {
					t.Errorf("[%v]: did not remove all expected WS(%v): *%v*", test.desc, ws, test.input)
				}
			}
		}
	}
}

func TestConsumeToNewline(t *testing.T) {
	tests := []struct {
		desc    string
		input   string
		want    string
		wantErr bool
	}{{
		desc:  "Success",
		input: " this is\nstop here\n",
		want:  "stop here\n",
	}, {
		desc:    "Report Error from reader beyond newline",
		input:   " ",
		wantErr: true,
	}, {
		desc:    "Report Error from reader death",
		input:   "",
		wantErr: true,
	}}

	for _, test := range tests {
		p := NewParser(strings.NewReader(test.input))
		err := p.consumeToNewline()
		switch {
		case err != nil && !test.wantErr:
			t.Errorf("[%v]: test got error when not expecting one: %v", test.desc, err)
		case err == nil && test.wantErr:
			t.Errorf("[%v]: test did not get error when expecting one", test.desc)
		case err == nil:
			got := p.String()
			if !cmp.Equal(got, test.want) {
				t.Errorf("[%v]: got/want mismatch: got/want\n%v", test.desc, cmp.Diff(got, test.want))
			}
		}
	}
}

func TestIsPunctuation(t *testing.T) {
	tests := []struct {
		desc string
		char rune
		want bool
	}{{
		desc: "excalamation",
		char: '!',
		want: true,
	}, {
		desc: "atsign",
		char: '@',
		want: true,
	}, {
		desc: "sharp",
		char: '#',
		want: true,
	}, {
		desc: "dollar",
		char: '$',
		want: true,
	}, {
		desc: "percent",
		char: '%',
		want: true,
	}, {
		desc: "carrot",
		char: '^',
		want: true,
	}, {
		desc: "ampersand",
		char: '&',
		want: true,
	}, {
		desc: "asterick",
		char: '*',
		want: true,
	}, {
		desc: "left paren",
		char: '(',
		want: true,
	}, {
		desc: "right paren",
		char: ')',
		want: true,
	}, {
		desc: "underscore",
		char: '_',
		want: true,
	}, {
		desc: "plus",
		char: '+',
		want: true,
	}, {
		desc: "minus",
		char: '-',
		want: true,
	}, {
		desc: "equal",
		char: '=',
		want: true,
	}, {
		desc: "left bracket",
		char: '[',
		want: true,
	}, {
		desc: "right bracket",
		char: ']',
		want: true,
	}, {
		desc: "backslash",
		char: '\\',
		want: true,
	}, {
		desc: "left curly",
		char: '{',
		want: true,
	}, {
		desc: "right curly",
		char: '}',
		want: true,
	}, {
		desc: "pipe",
		char: '|',
		want: true,
	}, {
		desc: "colon",
		char: ':',
		want: true,
	}, {
		desc: "semi-colon",
		char: ';',
		want: true,
	}, {
		desc: "double quote",
		char: '"',
		want: true,
	}, {
		desc: "single quote",
		char: '\'',
		want: true,
	}, {
		desc: "left angle",
		char: '<',
		want: true,
	}, {
		desc: "right angle",
		char: '>',
		want: true,
	}, {
		desc: "comma",
		char: ',',
		want: true,
	}, {
		desc: "period",
		char: '.',
		want: true,
	}, {
		desc: "fwd slash",
		char: '/',
		want: true,
	}, {
		desc: "question mark",
		char: '?',
		want: true,
	}, {
		desc: "backtick",
		char: '`',
		want: true,
	}, {
		desc: "tilde",
		char: '~',
		want: true,
	}, {
		desc: "letter",
		char: 'y',
		want: false,
	}, {
		desc: "numbe",
		char: '2',
		want: false,
	}, {
		desc: "space",
		char: ' ',
		want: false,
	}}

	for _, test := range tests {
		got := isPunctuation(test.char)
		if got != test.want {
			t.Errorf("[%v]: got/want mismatch: %v/%v", test.desc, got, test.want)
		}
	}
}
