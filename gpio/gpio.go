package gpio

import (
	"errors"
	"os/exec"
	"strconv"
	"unicode/utf8"
)

const TextNotPinSupport = "Pin is not Supported."
const TextUnknownError = "Unknown error."

func CheckMode(pinNumStr string) (bool, error) {
	var resGetMode, errGetMode = exec.Command("cat", "/sys/class/gpio/gpio"+pinNumStr+"/direction").Output()
	var resGetModeStr string = string(resGetMode)
	var resGetModeStrLen int = utf8.RuneCountInString(resGetModeStr)

	if errGetMode != nil {
		return false, errGetMode
	} else if string([]rune(resGetModeStr)[:resGetModeStrLen-1]) == "in" {
		return false, nil
	} else if string([]rune(resGetModeStr)[:resGetModeStrLen-1]) == "out" {
		return true, nil
	} else {
		return false, errors.New("checkMode: " + TextUnknownError)
	}
}

func SetMode(pinNum int, mode bool) error {
	var pinNumStr string = strconv.Itoa(pinNum)
	var resPreMode, errPreMode = CheckMode(pinNumStr)
	var modeStr string = ""

	if errPreMode != nil {
		return errPreMode
	} else if resPreMode == mode {
		return nil
	}

	if mode {
		modeStr = "out"
	} else {
		modeStr = "in"
	}

	var _, errSetMode = exec.Command("sh", "-c", "echo "+modeStr+" > /sys/class/gpio/gpio"+pinNumStr+"/direction").Output()

	return errSetMode
}

func SetValue(pinNum int, value bool) error {
	var pinNumStr string = strconv.Itoa(pinNum)
	var isExist bool = IsEnabled(pinNum)
	var valueNumStr string = "0"

	if value {
		valueNumStr = "1"
	} else {
		valueNumStr = "0"
	}

	if !isExist {
		return errors.New("Gpio sysfs not exist.")
	}

	var _, errSetValue = exec.Command("sh", "-c", "echo "+valueNumStr+" > /sys/class/gpio/gpio"+pinNumStr+"/value").Output()

	return errSetValue
}

func IsEnabled(pinNum int) bool {
	var _, errExist = exec.Command("ls", "/sys/class/gpio/gpio"+strconv.Itoa(pinNum)+"/").Output()

	return errExist == nil
}

func EnableGPIO(pinNum int) error {
	if IsEnabled(pinNum) {
		return errors.New("Already enabled.")
	}

	var pinNumStr string = strconv.Itoa(pinNum)
	var resExport, errExport = exec.Command("sh", "-c", "echo "+pinNumStr+" > /sys/class/gpio/export").Output()
	var resExist, errExist = exec.Command("ls", "/sys/class/gpio/gpio"+pinNumStr+"/").Output()

	if errExport != nil {
		return errors.New("Cannot export " + pinNumStr + ".")
	} else if errExist != nil {
		return errors.New("Exported but not exist gpio sysfs.")
	} else if resExport != nil && resExist != nil {
		return nil
	} else {
		return errors.New("enableGPIO: " + TextUnknownError)
	}
}
