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

	acsu := tv.SURROUND

	if raw, desc, err := tv.SendCommand(acsu.COMMAND, acsu.STATUS); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}
	if raw, desc, err := tv.SendCommand(acsu.COMMAND, acsu.TOGGLE); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}

	if raw, desc, err := tv.SendCommand(acsu.COMMAND, acsu.STATUS); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}
	if raw, desc, err := tv.SendCommand(acsu.COMMAND, acsu.OFF); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}

	if raw, desc, err := tv.SendCommand(acsu.COMMAND, acsu.STATUS); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}
	if raw, desc, err := tv.SendCommand(acsu.COMMAND, acsu.ON); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}
}
