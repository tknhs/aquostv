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

	inp3 := tv.INPUT_SELECTION.INP3
	inp4 := tv.INPUT_SELECTION.INP4

	// INP3
	if resp, err := tv.SendCommand(inp3.COMMAND, inp3.AUTO); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}
	if resp, err := tv.SendCommand(inp3.COMMAND, inp3.STATUS); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}

	// INP4
	if resp, err := tv.SendCommand(inp4.COMMAND, inp4.AUTO); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}
	if resp, err := tv.SendCommand(inp4.COMMAND, inp4.STATUS); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}
}
