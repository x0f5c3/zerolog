//go:build !binary_log

package diode_test

import (
	"fmt"
	"os"

	"github.com/x0f5c3/zerolog"
	"github.com/x0f5c3/zerolog/diode"
	"github.com/x0f5c3/zerolog/internal/utils"
)

func ExampleNewWriter() {
	w := diode.NewWriter(os.Stdout, 1000, 0, func(missed int) {
		fmt.Printf("Dropped %d messages\n", missed)
	})
	log := zerolog.New(w)
	log.Print("test")

	utils.HandleErr(w.Close(), "Failed to close the diode writer")

	// Output: {"level":"debug","message":"test"}
}
