package aquostv

import "errors"

type AQUOS_INPUT_SELECTION struct {
	INP3 *INPUT_SELECTION_3
	INP4 *INPUT_SELECTION_4
}

type INPUT_SELECTION_3 struct {
	COMMAND        string
	STATUS         string
	AUTO           string
	D_TERMINAL     string
	VIDEO_TERMINAL string
}

type INPUT_SELECTION_4 struct {
	COMMAND        string
	STATUS         string
	AUTO           string
	S_TERMINAL     string
	VIDEO_TERMINAL string
}

func initInputSelection() *AQUOS_INPUT_SELECTION {
	return &AQUOS_INPUT_SELECTION{
		INP3: initInputSelection3(),
		INP4: initInputSelection4(),
	}
}

func initInputSelection3() *INPUT_SELECTION_3 {
	return &INPUT_SELECTION_3{
		COMMAND:        "INP3",
		STATUS:         "????",
		AUTO:           "0000",
		D_TERMINAL:     "0001",
		VIDEO_TERMINAL: "0004",
	}
}

func initInputSelection4() *INPUT_SELECTION_4 {
	return &INPUT_SELECTION_4{
		COMMAND:        "INP4",
		STATUS:         "????",
		AUTO:           "0000",
		S_TERMINAL:     "0003",
		VIDEO_TERMINAL: "0004",
	}
}

func (tv *TV) inputSelection() (string, error) {
	switch tv.readResponse {
	case "OK":
		return tv.readResponse, nil
	case "ERR":
		return tv.readResponse, errors.New("Response ERROR.")
	case "0":
		return "AUTO", nil
	case "1":
		return "D_TERMINAL", nil
	case "3":
		return "S_TERMINAL", nil
	case "4":
		return "VIDEO_TERMINAL", nil
	default:
		return "", errors.New("Unknown ERROR.")
	}
}
