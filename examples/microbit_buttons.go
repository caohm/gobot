package main

import (
	"fmt"
	"os"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/microbit"
)

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ubit := microbit.NewButtonDriver(bleAdaptor)

	work := func() {
		ubit.On(microbit.ButtonA, func(data interface{}) {
			fmt.Println("button A")
		})

		ubit.On(microbit.ButtonB, func(data interface{}) {
			fmt.Println("button B")
		})
	}

	robot := gobot.NewRobot("buttonBot",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{ubit},
		work,
	)

	robot.Start()
}
