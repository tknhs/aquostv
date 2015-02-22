package aquostv

import "errors"

type AQUOS_AUDIO_SWITCHING struct {
	COMMAND string
	TOGGLE  string
}

func initAudioSwitching() *AQUOS_AUDIO_SWITCHING {
	return &AQUOS_AUDIO_SWITCHING{
		COMMAND: "ACHA",
		TOGGLE:  "0000",
	}
}

func (tv *TV) audioSwitching() {
	switch tv.responseRaw {
	case "OK":
		tv.responseDescription = "OK"
	case "ERR":
		tv.responseDescription = "ERR"
	default:
		tv.responseError = errors.New("Unknown ERROR.")
	}
}
