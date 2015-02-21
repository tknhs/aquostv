package aquostv

import "errors"

type AQUOS_INPUT_SWITCHING struct {
	ITGD *INPUT_SWITCHING_TOGGLE
	ITVD *INPUT_SWITCHING_TV
	IAVD *INPUT_SWITCHING_INPUT
	IDEG *INPUT_SWITCHING_BROADCAST_SWITCHING
}

type INPUT_SWITCHING_TOGGLE struct {
	COMMAND string
	TOGGLE  string
}

type INPUT_SWITCHING_TV struct {
	COMMAND string
	DEFAULT string
}

type INPUT_SWITCHING_INPUT struct {
	COMMAND string
	STATUS  string
	INPUT_1 string
	INPUT_2 string
	INPUT_3 string
	INPUT_4 string
	INPUT_5 string
}

type INPUT_SWITCHING_BROADCAST_SWITCHING struct {
	COMMAND string
	TOGGLE  string
}

func initInputSwitching() *AQUOS_INPUT_SWITCHING {
	return &AQUOS_INPUT_SWITCHING{
		ITGD: initInputSwitchingToggle(),
		ITVD: initInputSwitchingTV(),
		IAVD: initInputSwitchingInput(),
		IDEG: initInputSwitchingBroadcastSwitching(),
	}
}

func initInputSwitchingToggle() *INPUT_SWITCHING_TOGGLE {
	return &INPUT_SWITCHING_TOGGLE{
		COMMAND: "ITGD",
		TOGGLE:  "0000",
	}
}

func initInputSwitchingTV() *INPUT_SWITCHING_TV {
	return &INPUT_SWITCHING_TV{
		COMMAND: "ITVD",
		DEFAULT: "0000",
	}
}

func initInputSwitchingInput() *INPUT_SWITCHING_INPUT {
	return &INPUT_SWITCHING_INPUT{
		COMMAND: "IAVD",
		STATUS:  "????",
		INPUT_1: "0001",
		INPUT_2: "0002",
		INPUT_3: "0003",
		INPUT_4: "0004",
		INPUT_5: "0005",
	}
}

func initInputSwitchingBroadcastSwitching() *INPUT_SWITCHING_BROADCAST_SWITCHING {
	return &INPUT_SWITCHING_BROADCAST_SWITCHING{
		COMMAND: "IDEG",
		TOGGLE:  "0000",
	}
}

func (tv *TV) inputSwitchingDefault() {
	switch tv.responseRaw {
	case "OK":
		tv.responseDescription = "OK"
	case "ERR":
		tv.responseDescription = "ERR"
	default:
		tv.responseError = errors.New("Unknown ERROR.")
	}
}

func (tv *TV) inputSwitchingInput() {
	switch tv.responseRaw {
	case "OK":
		tv.responseDescription = "OK"
	case "ERR":
		tv.responseDescription = "ERR"
	case "1":
		tv.responseDescription = "INPUT_1"
	case "2":
		tv.responseDescription = "INPUT_2"
	case "3":
		tv.responseDescription = "INPUT_3"
	case "4":
		tv.responseDescription = "INPUT_4"
	case "5":
		tv.responseDescription = "INPUT_5"
	default:
		tv.responseError = errors.New("Unknown ERROR.")
	}
}
