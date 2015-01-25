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

	itgd := tv.INPUT_SWITCHING.ITGD
	itvd := tv.INPUT_SWITCHING.ITVD
	ideg := tv.INPUT_SWITCHING.IDEG
	iavd := tv.INPUT_SWITCHING.IAVD

	// ITGD
	if resp, err := tv.SendCommand(itgd.COMMAND, itgd.TOGGLE); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}

	// ITVD
	if resp, err := tv.SendCommand(itvd.COMMAND, itvd.DEFAULT); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}

	// IDEG
	if resp, err := tv.SendCommand(ideg.COMMAND, ideg.TOGGLE); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}

	// IAVD
	if resp, err := tv.SendCommand(iavd.COMMAND, iavd.INPUT_2); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}

	if resp, err := tv.SendCommand(iavd.COMMAND, iavd.STATUS); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}
}
