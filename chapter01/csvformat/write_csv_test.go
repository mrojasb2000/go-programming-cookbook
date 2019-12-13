package csvformat

import (
	"bytes"
	"testing"
)

func TestBooks_ToCSV(t *testing.T) {
	tests := []struct {
		name    string
		books   *Books
		wantW   string
		wantErr bool
	}{
		// One instance
		{
			name: "base-case",
			books: &Books{
				Book{
					Author: "F Scott Fitzgerald",
					Title:  "The Great Gatsby",
				},
			},
			wantW:   "Author,Title\nF Scott Fitzgerald,The Great Gatsby\n",
			wantErr: false,
		},
		// Two instance
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := tt.books.ToCSV(w); (err != nil) != tt.wantErr {
				t.Errorf("Books.ToCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("Books.toCSV() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
