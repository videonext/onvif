package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kr/pretty"
	"github.com/videonext/onvif/discovery"
	"github.com/videonext/onvif/profiles/devicemgmt"
	"github.com/videonext/onvif/soap"
)

func main() {

	// discovery devices
	devices, err := discovery.StartDiscovery(5 * time.Second)
	if err != nil {
		fmt.Println(err.Error())
	}
	if len(devices) == 0 {
		fmt.Printf("No devices descovered\n")

		return
	}

	fmt.Printf("Discovered %d devices\n", len(devices))
	pretty.Println(devices)

	// Create soap client
	client := soap.NewClient(
		soap.WithTimeout(time.Second * 5),
	)
	client.AddHeader(soap.NewWSSSecurityHeader("root", "rootpass"))

	for _, d := range devices {
		// Create devicemgmt service instance and specify xaddr (which could be received in the discovery)
		dev := devicemgmt.NewDevice(client, d.XAddr)

		log.Println("devicemgmt.GetDeviceInformation", d.XAddr)
		{
			reply, err := dev.GetDeviceInformation(&devicemgmt.GetDeviceInformation{})
			if err != nil {
				if serr, ok := err.(*soap.SOAPFault); ok {
					pretty.Println(serr)
				}
				log.Fatalf("Request failed: %s", err.Error())
			}
			pretty.Println(reply)
		}

		log.Println("devicemgmt.GetServices", d.XAddr)
		{
			reply, err := dev.GetServices(&devicemgmt.GetServices{})
			if err != nil {
				if serr, ok := err.(*soap.SOAPFault); ok {
					pretty.Println(serr)
				}
				log.Fatalf("Request failed: %s", err.Error())
			}
			pretty.Println(reply)
		}

	}

}
