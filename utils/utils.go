package utils

import (
	"errors"
	"fmt"
	"os"

	"../mesgs"
	"../props"
)

func PrintExit(mes string) {
	fmt.Println(mes)
	os.Exit(1)
}

func CheckPinSupport(licheePinStr string) error {
	if _, ok := props.LicheePins[licheePinStr]; ok {
		return nil
	} else {
		return errors.New(mesgs.CreateErrPinNotSupport(props.LicheePins))
	}
}
