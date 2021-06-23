package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/service/iotdataplane"
)

// TurnOff
type TurnOff struct {
	IOTClient *iotdataplane.IoTDataPlane
	Topic     *string
	ThingID   *string
}

// Name return the action name
func (action *TurnOff) Name() string {
	return "OFF"
}

// Do perform the Turn Off action
func (action *TurnOff) Do() error {

	log.Printf("action %s triggered", action.Name())

	message := &ActionMessage{
		Action:  "off",
		ThingID: *action.ThingID,
	}

	jsonString, _ := json.Marshal(message)

	publishInput := &iotdataplane.PublishInput{
		Topic:   action.Topic,
		Payload: jsonString,
	}

	_, err := action.IOTClient.Publish(publishInput)

	return err
}
