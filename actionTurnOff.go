package main

import (
	"log"

	"github.com/aws/aws-sdk-go/service/iotdataplane"
)

// TurnOff
type TurnOff struct {
	IOTClient *iotdataplane.IoTDataPlane
	Topic     *string
	deviceId  *string
}

// Name return the action name
func (action *TurnOff) Name() string {
	return "OFF"
}

// Do perform the Turn Off action
func (action *TurnOff) Do() error {

	log.Printf("action %s triggered", action.Name())

	publishInput := &iotdataplane.PublishInput{
		Topic:   action.Topic,
		Payload: []byte("off"),
	}

	_, err := action.IOTClient.Publish(publishInput)

	return err
}
