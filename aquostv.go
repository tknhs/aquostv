package aquostv

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/tarm/goserial"
)

type TV struct {
	SerialObject    io.ReadWriteCloser
	readResponse    string
	POWER           *AQUOS_POWER
	INPUT_SWITCHING *AQUOS_INPUT_SWITCHING
	INPUT_SELECTION *AQUOS_INPUT_SELECTION
	AV_POSITION     *AQUOS_AV_POSITION
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
		SerialObject:    s,
		POWER:           initPower(),
		INPUT_SWITCHING: initInputSwitching(),
		INPUT_SELECTION: initInputSelection(),
		AV_POSITION:     initAVPosition(),
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
	// power
	case tv.POWER.COMMAND:
		resp, err = tv.power()
	// input switching
	case tv.INPUT_SWITCHING.ITGD.COMMAND,
		tv.INPUT_SWITCHING.ITVD.COMMAND,
		tv.INPUT_SWITCHING.IDEG.COMMAND:
		resp, err = tv.inputSwitchingDefault()
	case tv.INPUT_SWITCHING.IAVD.COMMAND:
		resp, err = tv.inputSwitchingInput()
	// input selection
	case tv.INPUT_SELECTION.INP3.COMMAND,
		tv.INPUT_SELECTION.INP4.COMMAND:
		resp, err = tv.inputSelection()
	// av position
	case tv.AV_POSITION.COMMAND:
		resp, err = tv.avPosition()
	}
	return resp, err
}