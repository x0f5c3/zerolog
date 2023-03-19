package diode_test

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
	"time"

	"github.com/x0f5c3/zerolog"
	"github.com/x0f5c3/zerolog/diode"
	"github.com/x0f5c3/zerolog/internal/cbor"
)

func handleErr(err error, l *zerolog.Logger, msg string) {
	if err != nil {
		l.Error().Err(err).Msg(msg)
	}
}

func TestNewWriter(t *testing.T) {
	buf := bytes.Buffer{}
	w := diode.NewWriter(&buf, 1000, 0, func(missed int) {
		fmt.Printf("Dropped %d messages\n", missed)
	})
	l := zerolog.New(w)
	l.Print("test")

	handleErr(w.Close(), l, "Failed to close the diode writer")
	want := "{\"level\":\"debug\",\"message\":\"test\"}\n"
	got := cbor.DecodeIfBinaryToString(buf.Bytes())
	if got != want {
		t.Errorf("Diode New Writer Test failed. got:%s, want:%s!", got, want)
	}
}

func TestClose(t *testing.T) {
	buf := bytes.Buffer{}
	w := diode.NewWriter(&buf, 1000, 0, func(missed int) {})
	l := zerolog.New(w)
	l.Print("test")
	handleErr(w.Close(), l, "Failed to close the diode writer")
}

func Benchmark(b *testing.B) {
	log.SetOutput(io.Discard)
	defer log.SetOutput(os.Stderr)
	benchs := map[string]time.Duration{
		"Waiter": 0,
		"Pooler": 10 * time.Millisecond,
	}
	for name, interval := range benchs {
		b.Run(name, func(b *testing.B) {
			w := diode.NewWriter(io.Discard, 100000, interval, nil)
			l := zerolog.New(w)
			defer handleErr(w.Close(), l, "Failed to close the diode discard writer")

			b.SetParallelism(1000)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					l.Print("test")
				}
			})
		})
	}
}
