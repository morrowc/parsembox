package parsembox

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	testDataDir = "testdata"
	msgsFile    = "four_msgs.txt"
)

func TestFindFrom(t *testing.T) {
	tests := []struct {
		desc    string
		input   string
		want    string
		wantErr bool
	}{{
		desc:  "Success, simple test",
		input: "  \nFrom foo@bar.org Sept 11, 2001\n",
		want:  "foo@bar.org",
	}, {
		desc:    "Success, simple test, no date",
		input:   "  \nFrom foo@bar.org\n",
		wantErr: true,
	}, {
		desc:  "Success, more garbage before From test",
		input: "123klasjd8 2j1asd ds Fro ather this From: athing  \nFrom foo@bar.org Sept 11, 2001\n",
		want:  "foo@bar.org",
	}, {
		desc:    "Error, only From",
		input:   "From",
		wantErr: true,
	}, {
		desc:    "Error, only From<space>",
		input:   "From ",
		wantErr: true,
	}, {
		desc:    "Error raises, no from",
		input:   "123klasjd8 2j1asd ds Fro ather this From: athing  \noo@bar.org\n",
		wantErr: true,
	}, {
		desc:    "Error raises, no content",
		input:   "",
		wantErr: true,
	}, {
		desc:    "Error raises, F",
		input:   "F",
		wantErr: true,
	}, {
		desc:    "Error raises, Fr",
		input:   "Fr",
		wantErr: true,
	}, {
		desc:    "Error raises, Fro",
		input:   "Fro",
		wantErr: true,
	}, {
		desc: "Whitespace only",
		input: " 	 ",
		wantErr: true,
	}}

	for _, test := range tests {
		p := NewParser(strings.NewReader(test.input))
		_, gotFrom, _, err := p.FindFrom()

		switch {
		case err != nil && !test.wantErr:
			t.Errorf("[%v]: got an error when not expecting one: %v", test.desc, err)
		case err == nil && test.wantErr:
			t.Errorf("[%v]: did not get err when expecting one", test.desc)
		case err == nil:
			if gotFrom != test.want {
				t.Errorf("[%v]: failed to match got(%v) to want(%v)", test.desc, gotFrom, test.want)
			}
		}
	}
}

func TestNext(t *testing.T) {
	tests := []struct {
		desc     string
		testFile string
		want     string
		wantErr  bool
	}{{
		desc:     "Success test",
		testFile: msgsFile,
		want:     "root@db-server.fukuikenkei.jp",
	}}

	for _, test := range tests {
		fd, err := os.Open(filepath.Join(testDataDir, test.testFile))
		if err != nil {
			t.Fatalf("[%v]: failed to open test file(%v) for read: %v", test.desc, test.testFile, err)
		}
		p := NewParser(fd)
		got, err := p.Next()
		switch {
		case err != nil && !test.wantErr:
			t.Errorf("[%v]: got error when not expecting one: %v", test.desc, err)
		case err == nil && test.wantErr:
			t.Errorf("[%v]: did not get error when expecting one", test.desc)
		case err == nil:
			if *got != test.want {
				t.Errorf("[%v]: test got/want are not equal got/want:From %v/%v", test.desc, *got, test.want)
			}
		}
	}
}

func TestReadWord(t *testing.T) {
	tests := []struct {
		desc    string
		input   string
		want    []rune
		wantErr bool
	}{{
		desc:  "Success",
		input: " this ",
		want:  []rune("this"),
	}, {
		desc:  "Success no leading whitespace",
		input: "this ",
		want:  []rune("this"),
	}, {
		desc:  "Success punctatio in word",
		input: "thi!s ",
		want:  []rune("thi!s"),
	}, {
		desc:  "Success punctatio end of word",
		input: "this! ",
		want:  []rune("this!"),
	}, {
		desc:    "Error nothing but whitespace",
		input:   "    ",
		wantErr: true,
	}, {
		desc:    "Zero input should error",
		wantErr: true,
	}}

	for _, test := range tests {
		p := NewParser(strings.NewReader(test.input))
		got, err := p.readWord()
		switch {
		case err != nil && !test.wantErr:
		case err == nil && test.wantErr:
		case err == nil:
			if !cmp.Equal(got, test.want) {
				t.Errorf("[%v]: got/want mismatch: (+got/-want\n%v\n", test.desc, cmp.Diff(got, test.want))
			}
		}
	}
}
