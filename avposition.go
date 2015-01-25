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

func (tv *TV) avPosition() (string, error) {
	switch tv.readResponse {
	case "OK":
		return tv.readResponse, nil
	case "ERR":
		return tv.readResponse, errors.New("Response ERROR.")
	case "1":
		return "NORMAL", nil
	case "2":
		return "MOVIE", nil
	case "3":
		return "GAME", nil
	case "4":
		return "AV_MEMORY", nil
	case "5":
		return "DYNAMIC_FIXED", nil
	case "6":
		return "DYNAMIC", nil
	case "7":
		return "PC", nil
	default:
		return "", errors.New("Unknown ERROR.")
	}
}
