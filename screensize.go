package aquostv

import "errors"

type AQUOS_SCREEN_SIZE struct {
	COMMAND    string
	STATUS     string
	TOGGLE     string
	NORMAL     string
	SMART_ZOOM string
	WIDE       string
	CINEMA     string
	FULL       string
	FULL_1     string
	FULL_2     string
	UNDER_SCAN string
	DOT_BY_DOT string
}

func initScreenSize() *AQUOS_SCREEN_SIZE {
	return &AQUOS_SCREEN_SIZE{
		COMMAND:    "WIDE",
		STATUS:     "????",
		TOGGLE:     "0000",
		NORMAL:     "0001",
		SMART_ZOOM: "0002",
		WIDE:       "0003",
		CINEMA:     "0004",
		FULL:       "0005",
		FULL_1:     "0006",
		FULL_2:     "0007",
		UNDER_SCAN: "0008",
		DOT_BY_DOT: "0009",
	}
}

func (tv *TV) screenSize() {
	switch tv.responseRaw {
	case "OK":
		tv.responseDescription = "OK"
	case "ERR":
		tv.responseDescription = "ERR"
	case "1":
		tv.responseDescription = "NORMAL"
	case "2":
		tv.responseDescription = "SMART_ZOOM"
	case "3":
		tv.responseDescription = "WIDE"
	case "4":
		tv.responseDescription = "CINEMA"
	case "5":
		tv.responseDescription = "FULL"
	case "6":
		tv.responseDescription = "FULL_1"
	case "7":
		tv.responseDescription = "FULL_2"
	case "8":
		tv.responseDescription = "UNDER_SCAN"
	case "9":
		tv.responseDescription = "DOT_BY_DOT"
	default:
		tv.responseError = errors.New("Unknown ERROR.")
	}
}
