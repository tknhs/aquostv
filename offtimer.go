package aquostv

import "errors"
import "regexp"

type AQUOS_OFF_TIMER struct {
	COMMAND   string
	STATUS    string
	TIMER_OFF string
	TIMER_30  string
	TIMER_60  string
	TIMER_90  string
	TIMER_120 string
	TIMER_150 string
}

func initOffTimer() *AQUOS_OFF_TIMER {
	return &AQUOS_OFF_TIMER{
		COMMAND:   "OFTM",
		STATUS:    "????",
		TIMER_OFF: "0000",
		TIMER_30:  "0001",
		TIMER_60:  "0002",
		TIMER_90:  "0003",
		TIMER_120: "0004",
		TIMER_150: "0005",
	}
}

func (tv *TV) offTimer() {
	re := regexp.MustCompile("[0-9]+")
	switch tv.responseRaw {
	case "OK":
		tv.responseDescription = "OK"
	case "ERR":
		tv.responseDescription = "ERR"
	case re.FindString(tv.responseRaw):
		val := tv.responseRaw
		if val == "0" {
			tv.responseDescription = "TIMER_OFF"
		} else {
			tv.responseDescription = val + "_MINUTES"
		}
	default:
		tv.responseError = errors.New("Unknown ERROR.")
	}
}
