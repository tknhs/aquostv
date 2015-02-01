package main

import (
	"log"
	"strings"

	"github.com/tknhs/aquostv"
)

func main() {
	tv, err := aquostv.New("/dev/tty.usbserial-DJ001UMD")
	if err != nil {
		log.Fatal(err)
	}

	volm := tv.VOLUME

	resp, err := tv.SendCommand(volm.COMMAND, volm.STATUS)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
	nowVolume := strings.Split(resp, "_")[1]

	if resp, err := tv.SendCommand(volm.COMMAND, volm.VOLUME_05); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}
	if resp, err := tv.SendCommand(volm.COMMAND, volm.STATUS); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}

	if resp, err := tv.SendCommand(volm.COMMAND, "00"+nowVolume); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}
	if resp, err := tv.SendCommand(volm.COMMAND, volm.STATUS); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}
}
