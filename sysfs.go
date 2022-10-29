package gpio

import (
	"fmt"
	"io/ioutil"
)

const (
	basePath              = "/sys/class/gpio"
	exportPath            = basePath + "/export"
	unexportPath          = basePath + "/unexport"
	gpioPathTemplate      = basePath + "/%s"
	directionPathTemplate = gpioPathTemplate + "/direction"
	valuePathTemplate     = gpioPathTemplate + "/value"
)

func export(pinNum string) error {
	err := ioutil.WriteFile(exportPath, []byte(pinNum), 0)
	if err != nil {
		return fmt.Errorf("cannot export pinNum=%s, %w", pinNum, err)
	}

	return nil
}

func unexport(pinNum string) error {
	err := ioutil.WriteFile(unexportPath, []byte(pinNum), 0)
	if err != nil {
		return fmt.Errorf("cannot unexport pinNum=%s, %w", pinNum, err)
	}

	return nil
}

const (
	in  = "in"
	out = "out"
)

func setDirection(pinName, direction string) error {
	if direction != in && direction != out {
		return fmt.Errorf("unknown direction %s", direction)
	}

	directionPath := fmt.Sprintf(directionPathTemplate, pinName)
	err := ioutil.WriteFile(directionPath, []byte(direction), 0)
	if err != nil {
		return fmt.Errorf("cannot set direction %s, %w", direction, err)
	}

	return nil
}

func direction(pinName string) (string, error) {
	directionPath := fmt.Sprintf(directionPathTemplate, pinName)
	direction, err := ioutil.ReadFile(directionPath)
	if err != nil {
		return "", fmt.Errorf("cannot read direction, %w", err)
	}

	return string(direction), nil
}

const (
	low  = "0"
	high = "1"
)

func setState(pinName, level string) error {
	if level != low && level != high {
		return fmt.Errorf("unknown level %s", level)
	}

	valuePath := fmt.Sprintf(valuePathTemplate, pinName)
	err := ioutil.WriteFile(valuePath, []byte(level), 0)
	if err != nil {
		return fmt.Errorf("cannot set state %s, %w", level, err)
	}

	return nil
}

func state(pinName string) (string, error) {
	valuePath := fmt.Sprintf(valuePathTemplate, pinName)
	level, err := ioutil.ReadFile(valuePath)
	if err != nil {
		return "", fmt.Errorf("cannot read state, %w", err)
	}

	return string(level), nil
}
