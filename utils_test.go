package parsembox

import (
	"strings"
	"testing"
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
		got := IsLetter(test.char)
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
		got := IsDigit(test.char)
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
		got := IsWhitespace(test.char)
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
		got := IsSpace(test.char)
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
		got := IsNewline(test.char)
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
		got := IsColon(test.char)
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
		got := IsOctothorpe(test.char)
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
	}}

	ws := []string{" ", "\n", "\t"}

	for _, test := range tests {
		p := NewParser(strings.NewReader(test.input))
		err := p.ConsumeWS()
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
