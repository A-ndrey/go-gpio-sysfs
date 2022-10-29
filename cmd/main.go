package main

import (
	"github.com/A-ndrey/go-gpio-sysfs"
	"log"
	"time"
)

func main() {
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
		}
	}()

	redPin, err := gpio.TakePin(20)
	if err != nil {
		return
	}
	defer redPin.Free()

	greenPin, err := gpio.TakePin(21)
	if err != nil {
		return
	}
	defer greenPin.Free()

	err = redPin.Out(1)
	if err != nil {
		return
	}

	err = greenPin.Out(1)
	if err != nil {
		return
	}

	for i := 0; i < 10; i++ {
		err = redPin.Out(0)
		if err != nil {
			return
		}
		time.Sleep(time.Second)
		err = redPin.Out(1)
		if err != nil {
			return
		}
		err = greenPin.Out(0)
		if err != nil {
			return
		}
		time.Sleep(time.Second)
		err = greenPin.Out(1)
		if err != nil {
			return
		}
	}

}
