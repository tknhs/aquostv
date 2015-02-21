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

	wide := tv.SCREEN_SIZE

	if raw, desc, err := tv.SendCommand(wide.COMMAND, wide.TOGGLE); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}
	if raw, desc, err := tv.SendCommand(wide.COMMAND, wide.STATUS); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}
}
