package main

import (
	"fmt"
	"log"

	"tinygo.org/x/bluetooth"
)

const targetDevice = "loupi5"

func main() {
	adapter := bluetooth.DefaultAdapter

	// Enable BLE interface.
	err := adapter.Enable()
	if err != nil {
		log.Fatal("failed to enable BLE stack:", err)
	}

	ch := make(chan bluetooth.ScanResult, 1)
	// Start scanning.
	println("scanning...")
	err = adapter.Scan(func(adapter *bluetooth.Adapter, result bluetooth.ScanResult) {
		//println("found device:", result.Address.String(), result.RSSI, result.LocalName())
		if result.AdvertisementPayload.LocalName() == targetDevice {
			adapter.StopScan()
			log.Printf("found target device %s with address %s\n", targetDevice, result.Address.String())
			ch <- result
		}
	})
	if err != nil {
		log.Fatal(err)
	}

	result := <-ch
	device, err := adapter.Connect(result.Address, bluetooth.ConnectionParams{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to ", result.Address.String())

	defer func() {
		err := device.Disconnect()
		if err != nil {
			log.Fatal("failed to disconnect:", err)
		}
	}()

	svcs, err := device.DiscoverServices(nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range svcs {
		log.Println("found service", s.String())
	}
}
