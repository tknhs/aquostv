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

	if resp, err := tv.SendCommand(tv.AVMD.COMMAND, tv.AVMD.TOGGLE); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}

	if resp, err := tv.SendCommand(tv.AVMD.COMMAND, tv.AVMD.STATUS); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}
}
