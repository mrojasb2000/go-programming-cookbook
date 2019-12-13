package csvformat

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

// Movie will hold out parsed CSV
type Movie struct {
	Title    string
	Director string
	Year     int
}

// ReadCSV gives shows some examples of processsing CSV
// that is passed in as an io.Reader
func ReadCSV(b io.Reader) ([]Movie, error) {
	r := csv.NewReader(b)

	// There area some optional configuration options
	r.Comma = ';'
	r.Comment = '-'

	var movies []Movie

	// grab and ignore the header for now
	// we may also want to use this for a dictionary key or
	// some other form of lookup
	_, err := r.Read()
	if err != nil {
		return nil, err
	}

	// loop until it's all processed
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		year, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			return nil, err
		}

		m := Movie{record[0], record[1], int(year)}
		movies = append(movies, m)
	}
	return movies, nil
}

// AddMoviesFromText uses the CSV parser with a string
func AddMoviesFromText() error {
	// this is an example of us taking a string, converting
	// it into a buffer, and reading it
	// with the csv package
	in := `
	- first our header
	movie title;directory;year released
	
	- then some data
	Guardians of the Galaxy Vol. 2;James Gunn;2017
	Start Wars: Episode VIII;Rain Johnson;2017
	`
	b := bytes.NewBufferString(in)
	m, err := ReadCSV(b)
	if err != nil {
		return err
	}
	fmt.Printf("%#vn", m)
	return nil
}
