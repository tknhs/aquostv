package main

import (
	"log"

	"github.com/tknhs/aquostv"
)

func main() {
	tv, err := aquostv.New("/dev/tty.usbserial-DJ001UMD")
	if err != nil {
		log.Fatal(err)
	}

	avpos := tv.AV_POSITION

	if raw, desc, err := tv.SendCommand(avpos.COMMAND, avpos.TOGGLE); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}

	if raw, desc, err := tv.SendCommand(avpos.COMMAND, avpos.STATUS); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}
}
