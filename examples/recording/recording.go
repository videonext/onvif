package main

import (
	"fmt"
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

	var token recording.RecordingReference
	// Create service instance and specify xaddr (which could be received in the devicemgmt.GetServices())
	r := recording.NewRecordingPort(client, "http://192.168.27.66/onvif/services")
	{
		cr := recording.CreateRecording{RecordingConfiguration: recording.RecordingConfiguration{}}
		cr.RecordingConfiguration.Content = "mycontent"
		cr.RecordingConfiguration.MaximumRetentionTime = "P7D" //"PT0M30S"
		cr.RecordingConfiguration.Source =
			recording.RecordingSourceInformation{
				SourceId:    "http://192.168.27.66/onvif/services",
				Name:        "mysourcename",
				Location:    "mysourcelocation",
				Description: "mysourcedescription",
				Address:     "http://192.168.27.66/onvif/services"}

		reply, err := r.CreateRecording(&cr)

		if err != nil {
			if serr, ok := err.(*soap.SOAPFault); ok {
				pretty.Println(serr)
			}
			log.Fatalf("Request failed: %s", err.Error())
		}
		pretty.Println(reply)
		token = reply.RecordingToken
		fmt.Printf("Token: %s\n", token)

		dr := recording.DeleteRecording{RecordingToken: token}
		reply2, err := r.DeleteRecording(&dr)
		if err != nil {
			if serr, ok := err.(*soap.SOAPFault); ok {
				pretty.Println(serr)
			}
			log.Fatalf("Request failed: %s", err.Error())
		}
		pretty.Println(reply2)

		fmt.Println("Getting all records")
		reply3, err := r.GetRecordings(&recording.GetRecordings{})
		if err != nil {
			if serr, ok := err.(*soap.SOAPFault); ok {
				pretty.Println(serr)
			}
			log.Fatalf("Request failed: %s", err.Error())
		}
		pretty.Println(reply3)

	}

	// .............

}
