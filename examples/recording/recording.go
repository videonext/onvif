package main

import (
	"log"
	"time"

	"github.com/kr/pretty"
	"github.com/videonext/onvif/profiles/recording"
	"github.com/videonext/onvif/soap"
)

func main() {

	// Create soap client
	client := soap.NewClient(
		soap.WithTimeout(time.Second * 5),
	)
	client.AddHeader(soap.NewWSSSecurityHeader("root", "pass"))

	// Create service instance and specify xaddr (which could be received in the devicemgmt.GetServices())
	r := recording.NewRecordingPort(client, "http://192.168.27.66/onvif/services")
	{
		rc := recording.CreateRecording{RecordingConfiguration: recording.RecordingConfiguration{}}
		rc.RecordingConfiguration.Content = "mycontent"
		rc.RecordingConfiguration.MaximumRetentionTime = "0" //"P7D" //"PT0M30S"
		rc.RecordingConfiguration.Source =
			recording.RecordingSourceInformation{
				SourceId:    "http://192.168.27.66/onvif/services",
				Name:        "mysourcename",
				Location:    "mysourcelocation",
				Description: "mysourcedescription",
				Address:     "http://192.168.27.66/onvif/services"}

		reply, err := r.CreateRecording(&rc)

		if err != nil {
			if serr, ok := err.(*soap.SOAPFault); ok {
				pretty.Println(serr)
			}
			log.Fatalf("Request failed: %s", err.Error())
		}
		pretty.Println(reply)

		//		spew.Dump(reply)
	}

	// .............

}
