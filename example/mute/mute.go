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

	mute := tv.MUTE

	if raw, desc, err := tv.SendCommand(mute.COMMAND, mute.STATUS); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}
	if raw, desc, err := tv.SendCommand(mute.COMMAND, mute.TOGGLE); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}

	if raw, desc, err := tv.SendCommand(mute.COMMAND, mute.STATUS); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}
	if raw, desc, err := tv.SendCommand(mute.COMMAND, mute.MUTE); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}

	if raw, desc, err := tv.SendCommand(mute.COMMAND, mute.STATUS); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}
	if raw, desc, err := tv.SendCommand(mute.COMMAND, mute.UNMUTE); err != nil {
		log.Fatal(err)
	} else {
		log.Println(raw, desc)
	}
}
