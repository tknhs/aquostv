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

	power := tv.POWER

	if raw, desc, err := tv.SendCommand(power.COMMAND, power.OFF); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}
	if raw, desc, err := tv.SendCommand(power.COMMAND, power.STATUS); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}

	if raw, desc, err := tv.SendCommand(power.COMMAND, power.ON); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}
	if raw, desc, err := tv.SendCommand(power.COMMAND, power.STATUS); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}
}
