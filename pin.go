package gpio

import (
	"fmt"
	"strconv"
)

type Level uint8

const (
	Low Level = iota
	High
)

type Pin interface {
	Free() error
	Out(Level) error
	In() (Level, error)
}

type pinImpl struct {
	name   string
	number string
}

func TakePin(pinNum uint8) (Pin, error) {
	if pinNum < 2 || 27 < pinNum {
		return nil, fmt.Errorf("unknown pin: %d", pinNum)
	}

	pin := pinImpl{
		name:   fmt.Sprintf("gpio%d", pinNum),
		number: strconv.Itoa(int(pinNum)),
	}

	err := export(pin.number)
	if err != nil {
		return nil, fmt.Errorf("cannot take pin=%s, %w", pin.name, err)
	}

	return &pin, nil
}

func (p *pinImpl) Free() error {
	err := unexport(p.number)
	if err != nil {
		return fmt.Errorf("cannot free pin=%s, %w", p.name, err)
	}

	return nil
}

func (p *pinImpl) Out(level Level) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("cannot set \"out\" level = %d, %w", level, err)
		}
	}()

	err = setDirection(p.name, out)
	if err != nil {
		return err
	}

	if level == Low {
		err = setState(p.name, low)
	} else {
		err = setState(p.name, high)
	}
	if err != nil {
		return err
	}

	return nil
}

func (p *pinImpl) In() (level Level, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("cannot read \"in\" level, %w", err)
		}
	}()

	err = setDirection(p.name, in)
	if err != nil {
		return 0, err
	}

	levelStr, err := state(p.name)
	if err != nil {
		return 0, err
	}

	if levelStr == low {
		return Low, nil
	}

	return High, nil
}
