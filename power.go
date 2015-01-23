package aquostv

import "errors"

type POWER struct {
	COMMAND string
	STATUS  string
	OFF     string
	ON      string
}

func initPOWER() *POWER {
	return &POWER{
		COMMAND: "POWR",
		STATUS:  "????",
		OFF:     "0000",
		ON:      "0001",
	}
}

func (tv *TV) power() (string, error) {
	switch tv.readResponse {
	case "OK":
		return tv.readResponse, nil
	case "ERR":
		return tv.readResponse, errors.New("Response ERROR.")
	case "0":
		return "OFF", nil
	case "1":
		return "ON", nil
	default:
		return "", errors.New("Unknown ERROR.")
	}
}
