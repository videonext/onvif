package main

import (
	"log"
	"time"

	"github.com/kr/pretty"
	"github.com/videonext/onvif/profiles/device"
	"github.com/videonext/onvif/soap"
)

func main() {

	client := soap.NewClient(
		//"http://10.168.0.109/onvif/services",
		soap.WithTimeout(time.Second*5),
		soap.WithBasicAuth("root", "rootpass"),
		//		soap.WithTLS(&tls.Config{InsecureSkipVerify: true}),
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
