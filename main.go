package main

import (
	"flag"
	"fmt"
	"strconv"

	"./mesgs"
	"./proc"
	"./props"
	"./utils"
)

func main() {
	flag.Parse()

	var argComm string = flag.Arg(0)
	var argPin string = flag.Arg(1)
	var argValue string = flag.Arg(2)
	var argValueBool bool = false

	if _, ok := props.Commands[argComm]; !ok {
		utils.PrintExit(mesgs.CreateErrInvalidCommand(props.Commands))
	}

	switch argComm {
	case "pins":
		var text string = "Support pins on LicheePi nano.\n\nLicheePi nano Pin > SOC Pin\n---------------------------\n"

		for licheePin, socPin := range props.LicheePins {
			text += "" + licheePin + " > " + strconv.Itoa(socPin) + "\n"
		}
		fmt.Println(text)
	default:
		var notPinSupport error = utils.CheckPinSupport(argPin)

		if notPinSupport != nil {
			utils.PrintExit(notPinSupport.Error())
		}

		if argValue != "0" && argValue != "1" {
			utils.PrintExit(mesgs.CreateErrInvalidValue())
		} else if argValue == "0" {
			argValueBool = false
		} else if argValue == "1" {
			argValueBool = true
		}

		switch argComm {
		case "output":
			proc.Output(argPin, argValueBool)
		}
	}
}
