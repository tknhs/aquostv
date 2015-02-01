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

func (tv *TV) screenSize() (string, error) {
	switch tv.readResponse {
	case "OK":
		return tv.readResponse, nil
	case "ERR":
		return tv.readResponse, errors.New("Response ERROR.")
	case "1":
		return "NORMAL", nil
	case "2":
		return "SMART_ZOOM", nil
	case "3":
		return "WIDE", nil
	case "4":
		return "CINEMA", nil
	case "5":
		return "FULL", nil
	case "6":
		return "FULL_1", nil
	case "7":
		return "FULL_2", nil
	case "8":
		return "UNDER_SCAN", nil
	case "9":
		return "DOT_BY_DOT", nil
	default:
		return "", errors.New("Unknown ERROR.")
	}
}
