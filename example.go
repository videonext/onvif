package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kr/pretty"
	"github.com/videonext/onvif/discovery"
	"github.com/videonext/onvif/profiles/device"
	"github.com/videonext/onvif/soap"
)

func main() {

	// discovery devices
	devices, err := discovery.StartDiscovery(5 * time.Second)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Discovered %d devices\n", len(devices))
	pretty.Println(devices)

	client := soap.NewClient(
		soap.WithTimeout(time.Second * 5),
	)
	client.AddHeader(soap.NewWSSSecurityHeader("root", "rootpass"))

	service := device.NewDevice(client, "http://10.168.0.89:8000/onvif/device_service")
	{
		reply, err := service.GetDeviceInformation(&device.GetDeviceInformation{})
		if err != nil {
			if serr, ok := err.(*soap.SOAPFault); ok {

				pretty.Println(serr)
				//			log.Fatalf("could't device info: %s", err.Error())

			}
			log.Fatalf("Could't get device info: %s", err.Error())
		}
		pretty.Println(reply)
	}
	{
		reply, err := service.GetCapabilities(&device.GetCapabilities{})
		if err != nil {
			if serr, ok := err.(*soap.SOAPFault); ok {

				pretty.Println(serr)
				//			log.Fatalf("could't device info: %s", err.Error())

			}
			log.Fatalf("Request failed: %s", err.Error())
		}
		pretty.Println(reply)
	}
	{
		reply, err := service.GetServices(&device.GetServices{})
		if err != nil {
			if serr, ok := err.(*soap.SOAPFault); ok {

				pretty.Println(serr)
				//			log.Fatalf("could't device info: %s", err.Error())

			}
			log.Fatalf("Request failed: %s", err.Error())
		}
		pretty.Println(reply)
	}

}
