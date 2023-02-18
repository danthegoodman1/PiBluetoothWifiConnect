package main

import (
	"log"
	"time"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

func main() {
	// Enable BLE interface.
	log.Println("enabling ble stack")
	must("enable BLE stack", adapter.Enable())

	// Define the peripheral device info.
	log.Println("configuring advertisement")
	adv := adapter.DefaultAdvertisement()
	must("config adv", adv.Configure(bluetooth.AdvertisementOptions{
		LocalName: "TangiaPPC",
	}))

	// Start advertising
	must("start adv", adv.Start())

	adapter.SetConnectHandler(func(device bluetooth.Addresser, connected bool) {
		log.Printf("connected %+v device %+v", connected, device)
	})

	log.Println("advertising...")
	for {
		// Sleep forever.
		time.Sleep(time.Hour)
	}
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}
