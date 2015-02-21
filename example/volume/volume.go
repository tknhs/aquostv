package main

import (
	"log"
	"strings"

	"github.com/tknhs/aquostv"
)

func main() {
	tv, err := aquostv.New("/dev/tty.usbserial-DJ001UMD")
	if err != nil {
		log.Fatal(err)
	}

	volm := tv.VOLUME

	raw, desc, err := tv.SendCommand(volm.COMMAND, volm.STATUS)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(raw, desc)

	if raw, desc, err := tv.SendCommand(volm.COMMAND, volm.VOLUME_05); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}
	if raw, desc, err := tv.SendCommand(volm.COMMAND, volm.STATUS); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}

	if raw, desc, err := tv.SendCommand(volm.COMMAND, strings.Repeat("0", 4-len(raw))+raw); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}
	if raw, desc, err := tv.SendCommand(volm.COMMAND, volm.STATUS); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}
}
