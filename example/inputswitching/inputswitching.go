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
	if raw, desc, err := tv.SendCommand(itgd.COMMAND, itgd.TOGGLE); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}

	// ITVD
	if raw, desc, err := tv.SendCommand(itvd.COMMAND, itvd.DEFAULT); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}

	// IDEG
	if raw, desc, err := tv.SendCommand(ideg.COMMAND, ideg.TOGGLE); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}

	// IAVD
	if raw, desc, err := tv.SendCommand(iavd.COMMAND, iavd.INPUT_2); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}

	if raw, desc, err := tv.SendCommand(iavd.COMMAND, iavd.STATUS); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}
}
