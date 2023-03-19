package zerolog

import "github.com/x0f5c3/zerolog/log"

func HandleErr(err error, msg string, l ...*Logger) {
	if err != nil {
		if len(l) > 0 {
			for _, v := range l {
				v.Error().Err(err).Msg(msg)
			}
		} else {
			log.Error().Err(err).Msg(msg)
		}
	}
}
