//go:build !binary_log
// +build !binary_log

package diode_test

import (
	"fmt"
	"os"

	"github.com/x0f5c3/zerolog"
	"github.com/x0f5c3/zerolog/diode"
)

func ExampleNewWriter() {
	w := diode.NewWriter(os.Stdout, 1000, 0, func(missed int) {
		fmt.Printf("Dropped %d messages\n", missed)
	})
	log := zerolog.New(w)
	log.Print("test")

	w.Close()

	// Output: {"level":"debug","message":"test"}
}
