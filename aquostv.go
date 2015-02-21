package aquostv

import (
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/tarm/goserial"
)

type TV struct {
	serialObject        io.ReadWriteCloser
	responseRaw         string
	responseDescription string
	responseError       error
	POWER               *AQUOS_POWER
	INPUT_SWITCHING     *AQUOS_INPUT_SWITCHING
	INPUT_SELECTION     *AQUOS_INPUT_SELECTION
	AV_POSITION         *AQUOS_AV_POSITION
	VOLUME              *AQUOS_VOLUME
	SCREEN_SIZE         *AQUOS_SCREEN_SIZE
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
		serialObject:    s,
		POWER:           initPower(),
		INPUT_SWITCHING: initInputSwitching(),
		INPUT_SELECTION: initInputSelection(),
		AV_POSITION:     initAVPosition(),
		VOLUME:          initVolume(),
		SCREEN_SIZE:     initScreenSize(),
	}, nil
}

func (tv *TV) Close() error {
	err := tv.serialObject.Close()
	if err != nil {
		return err
	}
	return nil
}

func (tv *TV) writeSerial(data string) error {
	_, err := tv.serialObject.Write([]byte(data))
	if err != nil {
		return err
	}
	return nil
}

func (tv *TV) readSerial() (string, error) {
	buf := make([]byte, 4)
	n, err := tv.serialObject.Read(buf)
	if err != nil {
		return "", err
	}
	return string(buf[:n-1]), nil
}

func (tv *TV) SendCommand(aPart, bPart string) (string, string, error) {
	// initialize response values
	tv.responseRaw = ""
	tv.responseDescription = ""
	tv.responseError = nil

	command := fmt.Sprintf("%s%s\r", aPart, bPart)
	tv.writeSerial(command)
	time.Sleep(500 * time.Millisecond)
	resp, err := tv.readSerial()
	if err != nil {
		return "", "", err
	}
	tv.responseRaw = resp

	switch aPart {
	// power
	case tv.POWER.COMMAND:
		tv.power()
	// input switching
	case tv.INPUT_SWITCHING.ITGD.COMMAND,
		tv.INPUT_SWITCHING.ITVD.COMMAND,
		tv.INPUT_SWITCHING.IDEG.COMMAND:
		tv.inputSwitchingDefault()
	case tv.INPUT_SWITCHING.IAVD.COMMAND:
		tv.inputSwitchingInput()
	// input selection
	case tv.INPUT_SELECTION.INP3.COMMAND,
		tv.INPUT_SELECTION.INP4.COMMAND:
		tv.inputSelection()
	// av position
	case tv.AV_POSITION.COMMAND:
		tv.avPosition()
	// volume
	case tv.VOLUME.COMMAND:
		tv.volume()
	// screen size
	case tv.SCREEN_SIZE.COMMAND:
		tv.screenSize()
	default:
		tv.responseError = errors.New("Invalid Command.")
	}

	return tv.responseRaw, tv.responseDescription, tv.responseError
}
