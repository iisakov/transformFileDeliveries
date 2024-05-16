package main

import (
	"os"
	sdk "transformFileDeliveries/SDK"
)

func main() {
	var fileIn, fileOut *os.File
	var err error
	var requestDateDeliveryShift map[string][]string

	if len(os.Args) > 1 {
		commands := os.Args[1:]
		sdk.Handle(commands)
	}

	if fileIn, err = os.Open(sdk.FileInPath); err != nil {
		sdk.PrintError(err.Error(), "Файл с данными точно существует?")
		sdk.Exit()
	}

	if fileOut, err = os.Create(sdk.FileOutPath); err != nil {
		sdk.PrintError(err.Error(), "Файл с данными точно существует?")
		sdk.Exit()
	}

	defer fileIn.Close()

	requestDateDeliveryShift = sdk.ReadFile(fileIn, sdk.Sep)
	sdk.WriteFile(fileOut, sdk.PatternData, sdk.CountOrders, requestDateDeliveryShift)

}
