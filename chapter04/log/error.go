package log

import (
	"log"

	"github.com/pkg/errors"
)

// OriginalError returns the error original error
func OriginalError() error {
	return errors.New("error occurred")
}

// PassThroughError calls Original Error and
// forwads the error along after wrapping.
func PassThroughError() error {
	err := OriginalError()
	// no need to check error
	// since this works with nil
	return errors.Wrap(err, "in passthrougherror")
}

// FinalDestination deals with the error
// and doesn't forward it
func FinalDestination() {
	err := PassThroughError()
	if err != nil {
		// we log because an unexprected error occurred!
		log.Printf("an error occurred: %s\n", err.Error())
		return
	}
}
