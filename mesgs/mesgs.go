package mesgs

const CommandUsage string = "Usage: gpioutils [command] [pin] [value]"

func CreateErrInvalidCommand(commands map[string]string) string {
	var textHelp string = "\n\nCommands: \n\n"

	for command, text := range commands {
		textHelp += ("\t" + command + ": " + text + "\n")
	}

	return "Invalid command.\n\n" + CommandUsage + textHelp
}

func CreateErrPinNotSupport(supports map[string]int) string {
	var textHelp string = "\n\nSupported Pins: \n\n\t"

	for licheePin, _ := range supports {
		textHelp += (licheePin + ", ")
	}

	return "Pin is not Supported.\n\n" + CommandUsage + textHelp
}

func CreateErrInvalidValue() string {
	return "Arg value is invalid. \n\nChoice: \n\t0: LOW\n\t1: HIGH"
}

func CreateErrInternal(factor string, detail string) string {
	return "Internal Error > By '" + factor + "': " + detail
}
