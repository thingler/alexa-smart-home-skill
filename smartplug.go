package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/iotdataplane"
	smarthome "github.com/orktes/go-alexa-smarthome"
)

type ThinglerSmartPlug struct {
	val       interface{}
	iotClient *iotdataplane.IoTDataPlane
	config    *Config
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
	}

	turnOff := &TurnOff{
		IOTClient: smartplug.iotClient,
		Topic:     &smartplug.config.IOTTopic,
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

func mockThinglerSmartPlug(sm *smarthome.Smarthome, IOTClient *iotdataplane.IoTDataPlane, config *Config) {

	smartPlugDevice := smarthome.NewAbstractDevice(
		"thingler-plug-1",
		"Thingler smart plug",
		"Thingler",
		"Thingler smart plug",
	)
	smartPlugDevice.AddDisplayCategory("SMARTPLUG")
	capability := smartPlugDevice.NewCapability("PowerController")
	capability.AddPropertyHandler("powerState", &ThinglerSmartPlug{
		val:       "ON",
		iotClient: IOTClient,
		config:    config,
	})

	sm.AddDevice(smartPlugDevice)
}
