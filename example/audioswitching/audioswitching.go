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

	acha := tv.AUDIO_SWITCHING

	if raw, desc, err := tv.SendCommand(acha.COMMAND, acha.TOGGLE); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}
}
