package aquostv

import "errors"

type AQUOS_MUTE struct {
	COMMAND string
	STATUS  string
	TOGGLE  string
	MUTE    string
	UNMUTE  string
}

func initMute() *AQUOS_MUTE {
	return &AQUOS_MUTE{
		COMMAND: "MUTE",
		STATUS:  "????",
		TOGGLE:  "0000",
		MUTE:    "0001",
		UNMUTE:  "0002",
	}
}

func (tv *TV) mute() {
	switch tv.responseRaw {
	case "OK":
		tv.responseDescription = "OK"
	case "ERR":
		tv.responseDescription = "ERR"
	case "1":
		tv.responseDescription = "MUTE"
	case "2":
		tv.responseDescription = "UNMUTE"
	default:
		tv.responseError = errors.New("Unknown ERROR.")
	}
}
