package bytestrings

import "testing"

import "bytes"

func TestBuffer(t *testing.T) {
	var bufferTests = []struct {
		input  string
		output *bytes.Buffer
	}{
		{"abc", bytes.NewBufferString("abc")},
		{"", bytes.NewBufferString("")},
		{"heard", bytes.NewBufferString("heard")},
	}

	for _, tt := range bufferTests {
		if got, want := Buffer(tt.input), tt.output; got.String() != want.String() {
			t.Errorf("got: %#v; want: %#v", got, want)
		}
	}
}
