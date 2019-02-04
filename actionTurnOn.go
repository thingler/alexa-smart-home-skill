package main

import (
	"log"

	"github.com/aws/aws-sdk-go/service/iotdataplane"
)

// TurnOn
type TurnOn struct {
	IOTClient *iotdataplane.IoTDataPlane
	Topic     *string
}

// Name return the action name
func (action *TurnOn) Name() string {
	return "ON"
}

// Do perform the Turn On action
func (action *TurnOn) Do() error {

	log.Printf("action %s triggered", action.Name())

	publishInput := &iotdataplane.PublishInput{
		Topic:   action.Topic,
		Payload: []byte("on"),
	}

	_, err := action.IOTClient.Publish(publishInput)

	return err
}
