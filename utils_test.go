package parsembox

import "testing"

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
