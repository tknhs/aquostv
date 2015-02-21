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

func (tv *TV) inputSelection() {
	switch tv.responseRaw {
	case "OK":
		tv.responseDescription = "OK"
	case "ERR":
		tv.responseDescription = "ERR"
	case "0":
		tv.responseDescription = "AUTO"
	case "1":
		tv.responseDescription = "D_TERMINAL"
	case "3":
		tv.responseDescription = "S_TERMINAL"
	case "4":
		tv.responseDescription = "VIDEO_TERMINAL"
	default:
		tv.responseError = errors.New("Unknown ERROR.")
	}
}
