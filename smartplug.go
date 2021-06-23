package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iot"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
	smarthome "github.com/orktes/go-alexa-smarthome"
)

type ThinglerSmartPlug struct {
	val       interface{}
	iotClient *iotdataplane.IoTDataPlane
	config    *Config
	thingID   *string
}

func (smartplug *ThinglerSmartPlug) GetValue() (interface{}, error) {
	fmt.Printf("Getting value %+v\n", smartplug.val)
	return smartplug.val, nil
}

func (smartplug *ThinglerSmartPlug) SetValue(val interface{}) error {
	fmt.Printf("Received value %+v\n", val)

	smartplug.val = val
	smartplugValue := smartplug.val.(string)

	turnOn := &TurnOn{
		IOTClient: smartplug.iotClient,
		Topic:     &smartplug.config.IOTTopic,
		ThingID:   smartplug.thingID,
	}

	turnOff := &TurnOff{
		IOTClient: smartplug.iotClient,
		Topic:     &smartplug.config.IOTTopic,
		ThingID:   smartplug.thingID,
	}

	action, err := NewActionFactory().
		AddAction(turnOn).
		AddAction(turnOff).
		GetAction(&smartplugValue)
	if err == nil {
		err = action.Do()
	}

	return err
}

func (smartplug *ThinglerSmartPlug) UpdateChannel() <-chan interface{} {
	return nil
}

func getRegisteredSmartPlugs(sm *smarthome.Smarthome, clientIOT *iot.IoT, dataPlaneIOTClient *iotdataplane.IoTDataPlane, config *Config) {

	listThingsConfig := iot.ListThingsInput{
		// AttributeName:  aws.String("device"),
		// AttributeValue: aws.String("PowerController"),
		ThingTypeName: aws.String("Thingler_smartplug"),
	}

	things, err := clientIOT.ListThings(&listThingsConfig)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", things)

	for _, thing := range things.Things {
		smartPlugDevice := smarthome.NewAbstractDevice(
			*thing.Attributes["id"],
			*thing.Attributes["name"],
			*thing.Attributes["manufacturer"],
			*thing.Attributes["description"],
		)
		smartPlugDevice.AddDisplayCategory("SMARTPLUG")
		capability := smartPlugDevice.NewCapability("PowerController")
		capability.AddPropertyHandler("powerState", &ThinglerSmartPlug{
			val:       "ON",
			iotClient: dataPlaneIOTClient,
			config:    config,
			thingID:   thing.Attributes["id"],
		})
		sm.AddDevice(smartPlugDevice)
	}
}
