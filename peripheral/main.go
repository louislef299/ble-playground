package main

import (
	"log"
	"time"

	"tinygo.org/x/bluetooth"
)

func main() {
	adapter := bluetooth.DefaultAdapter

	// Enable BLE interface.
	err := adapter.Enable()
	if err != nil {
		log.Fatal("failed to enable BLE stack:", err)
	}

	// Define the peripheral device info.
	adv := adapter.DefaultAdvertisement()
	err = adv.Configure(bluetooth.AdvertisementOptions{
		LocalName: "loupi5",
		Interval:  bluetooth.Duration(20000),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Start advertising
	err = adv.Start()
	if err != nil {
		log.Fatal(err)
	}

	println("advertising...")
	for {
		// Sleep forever.
		time.Sleep(time.Hour)
	}
}
