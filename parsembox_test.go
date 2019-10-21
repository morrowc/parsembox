package parsembox

import (
	"strings"
	"testing"
)

func TestFindFrom(t *testing.T) {
	tests := []struct {
		desc    string
		input   string
		want    string
		wantErr bool
	}{{
		desc:  "Success, simple test",
		input: "  \nFrom foo@bar.org\n",
		want:  "foo@bar.org",
	}, {
		desc:  "Success, more garbage before From test",
		input: "123klasjd8 2j1asd ds Fro ather this From: athing  \nFrom foo@bar.org\n",
		want:  "foo@bar.org",
	}, {
		desc:    "Error raises, no from",
		input:   "123klasjd8 2j1asd ds Fro ather this From: athing  \noo@bar.org\n",
		wantErr: true,
	}}

	for _, test := range tests {
		p := NewParser(strings.NewReader(test.input))
		got, err := p.findFrom()

		switch {
		case err != nil && !test.wantErr:
			t.Errorf("[%v]: got an error when not expecting one: %v", test.desc, err)
		case err == nil && test.wantErr:
			t.Errorf("[%v]: did not get err when expecting one", test.desc)
		case err == nil:
			if got != test.want {
				t.Errorf("[%v]: failed to match got(%v) to want(%v)", test.desc, got, test.want)
			}
		}
	}
}
