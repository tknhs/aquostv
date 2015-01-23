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

	if resp, err := tv.SendCommand(tv.INP3.COMMAND, tv.INP3.AUTO); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}

	if resp, err := tv.SendCommand(tv.INP3.COMMAND, tv.INP3.STATUS); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}

	if resp, err := tv.SendCommand(tv.INP4.COMMAND, tv.INP4.AUTO); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}

	if resp, err := tv.SendCommand(tv.INP4.COMMAND, tv.INP4.STATUS); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}
}
