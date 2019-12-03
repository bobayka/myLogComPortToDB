package comPort

import (
	"context"
	"github.com/myLogComPortToDB/internal/myError"
	"github.com/pkg/errors"
	"github.com/tarm/serial"
	"log"
	"strings"
	"time"
)

const sizeInputBuf = 100
const dimension = 6

func ReadComPort(ch chan []string, port *serial.Port, timeout time.Duration) { //функция не доделана
	for {
		if err := readLine(ch, port, timeout); err != nil {
			log.Fatalf("can't read string: %s", err)
		}
	}
}

func WriteToComPort(port *serial.Port, msg string) {
	_, err := port.Write([]byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}

func readLine(ch chan []string, port *serial.Port, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	matchWholeLine := make([]string, 1)
	matchParam := make([]string, dimension)
	buf := make([]byte, 1)
	cont := make([]byte, 0, sizeInputBuf)
	finishCh := make(chan struct{})
	go func() {
		for {
			_, err := port.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			cont = append(cont, buf...)
			matchWholeLine = newLine.FindStringSubmatch(string(cont))
			if matchWholeLine != nil {
				//fmt.Printf("Cont: %+v\n", string(cont))
				matchParam = strings.Split(matchWholeLine[1], ",")
				break
			}
		}
		if err := port.Flush(); err != nil {
			log.Println("flush error:", err)
		}
		ch <- matchParam
		finishCh <- struct{}{}
	}()
	select {
	case <-finishCh:
	case <-ctx.Done():
		if strings.Contains(string(cont), myError.MPUDoesntConnect) {
			return errors.Wrap(myError.ErrNoMatches, myError.MPUDoesntConnect)
		} else if strings.Contains(string(cont), myError.MPUHalError) {
			return errors.Wrap(myError.ErrNoMatches, myError.MPUHalError)
		}
		return myError.ErrNoMatches
	}
	return nil
}
