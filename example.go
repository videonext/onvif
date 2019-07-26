package main

import (
	"log"
	"time"

	"wsdlgo_test/devicemgmt"

	"github.com/hooklift/gowsdl/soap"
	"github.com/kr/pretty"
)

func main() {

	client := soap.NewClient(
		"http://10.168.0.89:8000/onvif/device_service",
		//"http://10.168.0.109/onvif/services",
		soap.WithTimeout(time.Second*5),
		soap.WithBasicAuth("root", "rootpass"),
		//		soap.WithTLS(&tls.Config{InsecureSkipVerify: true}),
	)

	client.AddHeader(soap.NewWSSSecurityHeader("root", "rootpass"))

	service := devicemgmt.NewDevice(client)
	{
		reply, err := service.GetDeviceInformation(&devicemgmt.GetDeviceInformation{})
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
		reply, err := service.GetCapabilities(&devicemgmt.GetCapabilities{})
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
		reply, err := service.GetServices(&devicemgmt.GetServices{})
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
