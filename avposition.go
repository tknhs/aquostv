package aquostv

import "errors"

type AQUOS_AV_POSITION struct {
	COMMAND       string
	STATUS        string
	TOGGLE        string
	NORMAL        string
	MOVIE         string
	GAME          string
	AV_MEMORY     string
	DYNAMIC_FIXED string
	DYNAMIC       string
	PC            string
}

func initAVPosition() *AQUOS_AV_POSITION {
	return &AQUOS_AV_POSITION{
		COMMAND:       "AVMD",
		STATUS:        "????",
		TOGGLE:        "0000",
		NORMAL:        "0001",
		MOVIE:         "0002",
		GAME:          "0003",
		AV_MEMORY:     "0004",
		DYNAMIC_FIXED: "0005",
		DYNAMIC:       "0006",
		PC:            "0007",
	}
}

func (tv *TV) avPosition() {
	switch tv.responseRaw {
	case "OK":
		tv.responseDescription = "OK"
	case "ERR":
		tv.responseDescription = "ERR"
	case "1":
		tv.responseDescription = "NORMAL"
	case "2":
		tv.responseDescription = "MOVIE"
	case "3":
		tv.responseDescription = "GAME"
	case "4":
		tv.responseDescription = "AV_MEMORY"
	case "5":
		tv.responseDescription = "DYNAMIC_FIXED"
	case "6":
		tv.responseDescription = "DYNAMIC"
	case "7":
		tv.responseDescription = "PC"
	default:
		tv.responseError = errors.New("Unknown ERROR.")
	}
}
