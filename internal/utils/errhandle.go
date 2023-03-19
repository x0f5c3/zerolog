package utils

import "github.com/x0f5c3/zerolog/log"

func HandleErr(err error, msg string, writeFunc ...func(error, string)) {
	f := func() func(error, string) {
		if len(writeFunc) > 0 {
			return writeFunc[0]
		} else {
			return defaultErrWrite
		}
	}()
	if err != nil {
		f(err, msg)
	}
}

func defaultErrWrite(err error, msg string) {
	log.Error().Err(err).Msg(msg)
}
