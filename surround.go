package aquostv

import "errors"

type AQUOS_SURROUND struct {
	COMMAND string
	STATUS  string
	TOGGLE  string
	ON      string
	OFF     string
}

func initSurround() *AQUOS_SURROUND {
	return &AQUOS_SURROUND{
		COMMAND: "ACSU",
		STATUS:  "????",
		TOGGLE:  "0000",
		ON:      "0001",
		OFF:     "0002",
	}
}

func (tv *TV) surround() {
	switch tv.responseRaw {
	case "OK":
		tv.responseDescription = "OK"
	case "ERR":
		tv.responseDescription = "ERR"
	case "1":
		tv.responseDescription = "SURROUND_ON"
	case "2":
		tv.responseDescription = "SURROUND_OFF"
	default:
		tv.responseError = errors.New("Unknown ERROR.")
	}
}
