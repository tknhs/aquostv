package aquostv

import (
	"errors"
	"regexp"
)

type AQUOS_VOLUME struct {
	COMMAND   string
	STATUS    string
	VOLUME_00 string
	VOLUME_01 string
	VOLUME_02 string
	VOLUME_03 string
	VOLUME_04 string
	VOLUME_05 string
	VOLUME_06 string
	VOLUME_07 string
	VOLUME_08 string
	VOLUME_09 string
	VOLUME_10 string
	VOLUME_11 string
	VOLUME_12 string
	VOLUME_13 string
	VOLUME_14 string
	VOLUME_15 string
	VOLUME_16 string
	VOLUME_17 string
	VOLUME_18 string
	VOLUME_19 string
	VOLUME_20 string
	VOLUME_21 string
	VOLUME_22 string
	VOLUME_23 string
	VOLUME_24 string
	VOLUME_25 string
	VOLUME_26 string
	VOLUME_27 string
	VOLUME_28 string
	VOLUME_29 string
	VOLUME_30 string
	VOLUME_31 string
	VOLUME_32 string
	VOLUME_33 string
	VOLUME_34 string
	VOLUME_35 string
	VOLUME_36 string
	VOLUME_37 string
	VOLUME_38 string
	VOLUME_39 string
	VOLUME_40 string
	VOLUME_41 string
	VOLUME_42 string
	VOLUME_43 string
	VOLUME_44 string
	VOLUME_45 string
	VOLUME_46 string
	VOLUME_47 string
	VOLUME_48 string
	VOLUME_49 string
	VOLUME_50 string
	VOLUME_51 string
	VOLUME_52 string
	VOLUME_53 string
	VOLUME_54 string
	VOLUME_55 string
	VOLUME_56 string
	VOLUME_57 string
	VOLUME_58 string
	VOLUME_59 string
	VOLUME_60 string
}

func initVolume() *AQUOS_VOLUME {
	return &AQUOS_VOLUME{
		COMMAND:   "VOLM",
		STATUS:    "????",
		VOLUME_00: "0000",
		VOLUME_01: "0001",
		VOLUME_02: "0002",
		VOLUME_03: "0003",
		VOLUME_04: "0004",
		VOLUME_05: "0005",
		VOLUME_06: "0006",
		VOLUME_07: "0007",
		VOLUME_08: "0008",
		VOLUME_09: "0009",
		VOLUME_10: "0010",
		VOLUME_11: "0011",
		VOLUME_12: "0012",
		VOLUME_13: "0013",
		VOLUME_14: "0014",
		VOLUME_15: "0015",
		VOLUME_16: "0016",
		VOLUME_17: "0017",
		VOLUME_18: "0018",
		VOLUME_19: "0019",
		VOLUME_20: "0020",
		VOLUME_21: "0021",
		VOLUME_22: "0022",
		VOLUME_23: "0023",
		VOLUME_24: "0024",
		VOLUME_25: "0025",
		VOLUME_26: "0026",
		VOLUME_27: "0027",
		VOLUME_28: "0028",
		VOLUME_29: "0029",
		VOLUME_30: "0030",
		VOLUME_31: "0031",
		VOLUME_32: "0032",
		VOLUME_33: "0033",
		VOLUME_34: "0034",
		VOLUME_35: "0035",
		VOLUME_36: "0036",
		VOLUME_37: "0037",
		VOLUME_38: "0038",
		VOLUME_39: "0039",
		VOLUME_40: "0040",
		VOLUME_41: "0041",
		VOLUME_42: "0042",
		VOLUME_43: "0043",
		VOLUME_44: "0044",
		VOLUME_45: "0045",
		VOLUME_46: "0046",
		VOLUME_47: "0047",
		VOLUME_48: "0048",
		VOLUME_49: "0049",
		VOLUME_50: "0050",
		VOLUME_51: "0051",
		VOLUME_52: "0052",
		VOLUME_53: "0053",
		VOLUME_54: "0054",
		VOLUME_55: "0055",
		VOLUME_56: "0056",
		VOLUME_57: "0057",
		VOLUME_58: "0058",
		VOLUME_59: "0059",
		VOLUME_60: "0060",
	}
}

func (tv *TV) volume() {
	re := regexp.MustCompile("[0-9]+")
	switch tv.responseRaw {
	case "OK":
		tv.responseDescription = "OK"
	case "ERR":
		tv.responseDescription = "ERR"
	case re.FindString(tv.responseRaw):
		val := tv.responseRaw
		if len(val) == 1 {
			val = "0" + val
		}
		tv.responseDescription = "VOLUME_" + val
	default:
		tv.responseError = errors.New("Unknown ERROR.")
	}
}
