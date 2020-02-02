package proc

import (
	"../gpio"
	"../mesgs"
	"../props"
	"../utils"
)

func Output(argPin string, argValueBool bool) {
	var gpioNum = props.LicheePins[argPin]
	var isEnabled bool = gpio.IsEnabled(gpioNum)
	var errEnableGpio error = nil
	var errSetMode error = nil
	var errSetValue error = nil

	if !isEnabled {
		errEnableGpio = gpio.EnableGPIO(gpioNum)
	}
	if errEnableGpio == nil {
		errSetMode = gpio.SetMode(gpioNum, true)
	} else {
		utils.PrintExit(mesgs.CreateErrInternal("EnableGpio", errEnableGpio.Error()))
	}
	if errSetMode == nil {
		errSetValue = gpio.SetValue(gpioNum, argValueBool)
	} else {
		utils.PrintExit(mesgs.CreateErrInternal("SetMode", errSetMode.Error()))
	}
	if errSetValue != nil {
		utils.PrintExit(mesgs.CreateErrInternal("SetValue", errSetValue.Error()))
	}
}
