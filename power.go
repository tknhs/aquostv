package aquostv

import "errors"

type AQUOS_POWER struct {
	COMMAND string
	STATUS  string
	OFF     string
	ON      string
}

func initPower() *AQUOS_POWER {
	return &AQUOS_POWER{
		COMMAND: "POWR",
		STATUS:  "????",
		OFF:     "0000",
		ON:      "0001",
	}
}

func (tv *TV) power() {
	switch tv.responseRaw {
	case "OK":
		tv.responseDescription = "OK"
	case "ERR":
		tv.responseDescription = "ERR"
	case "0":
		tv.responseDescription = "OFF"
	case "1":
		tv.responseDescription = "ON"
	default:
		tv.responseError = errors.New("Unknown ERROR.")
	}
}
