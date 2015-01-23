package aquostv

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/tarm/goserial"
)

type TV struct {
	SerialObject io.ReadWriteCloser
	readResponse string
	POWR         *POWER
	INP3         *INPUT_SELECTION_3
	INP4         *INPUT_SELECTION_4
	AVMD         *AV_POSITION
}

func New(name string) (*TV, error) {
	c := &serial.Config{
		Name: name,
		Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		return &TV{}, err
	}
	return &TV{
		SerialObject: s,
		POWR:         initPOWER(),
		INP3:         initINPUTSELECTION3(),
		INP4:         initINPUTSELECTION4(),
		AVMD:         initAVPOSITION(),
	}, nil
}

func (tv *TV) Close() error {
	err := tv.SerialObject.Close()
	if err != nil {
		return err
	}
	return nil
}

func (tv *TV) writeSerial(data string) error {
	_, err := tv.SerialObject.Write([]byte(data))
	if err != nil {
		return err
	}
	return nil
}

func (tv *TV) readSerial() (string, error) {
	buf := make([]byte, 4)
	n, err := tv.SerialObject.Read(buf)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func (tv *TV) SendCommand(aPart, bPart string) (string, error) {
	command := fmt.Sprintf("%s%s\r", aPart, bPart)
	tv.writeSerial(command)
	time.Sleep(500 * time.Millisecond)
	resp, err := tv.readSerial()
	if err != nil {
		return "", err
	}
	tv.readResponse = strings.Split(resp, "\r")[0]

	switch aPart {
	case tv.POWR.COMMAND:
		resp, err = tv.power()
	case tv.INP3.COMMAND:
		resp, err = tv.inputselection3()
	case tv.INP4.COMMAND:
		resp, err = tv.inputselection4()
	case tv.AVMD.COMMAND:
		resp, err = tv.avposition()
	}
	return resp, err
}
