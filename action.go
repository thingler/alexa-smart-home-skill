package main

import (
	"errors"
	"fmt"
)

type ActionMessage struct {
	Action  string `json:"action"`
	ThingID string `json:"thingID"`
}

// Action interface
type Action interface {
	Name() string
	Do() error
}

// actionFactory contains a map[string] of all supported actions
// index value is the returned string from the action's Name() function
type actionFactory struct {
	actions map[string]Action
}

// NewActionFactory is the constructor function for the actions
func NewActionFactory() *actionFactory {
	factory := &actionFactory{}
	factory.actions = make(map[string]Action)
	return factory
}

// Add action
func (factory *actionFactory) AddAction(action Action) *actionFactory {
	factory.actions[action.Name()] = action
	return factory
}

// Return the specified action
func (factory *actionFactory) GetAction(name *string) (Action, error) {

	var err error

	action, registered := factory.actions[*name]
	if !registered {
		errorMessage := fmt.Sprintf(
			"Action, %s has not been implemented. Available actions:",
			*name,
		)
		for _, val := range factory.actions {
			errorMessage += fmt.Sprintf("\n %s", val.Name())
		}

		err = errors.New(errorMessage)
	}
	return action, err
}
