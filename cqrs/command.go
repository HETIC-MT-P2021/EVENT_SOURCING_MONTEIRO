package cqrs

import (
	"fmt"
	"reflect"
)

type CommandMessage interface {
	Payload() interface{}
	CommandType() string
}

type CommandBus struct {
	handlers map[string]CommandHandler
}

func InitCommandBus() *CommandBus {
	return &CommandBus{
		handlers: map[string]CommandHandler{},
	}
}

func (bus *CommandBus) RegisterHandler(handler CommandHandler, cmd interface{}) error {
	typeName := TypeOf(cmd)
	if _, ok := bus.handlers[typeName]; ok {
		fmt.Println("Already exists")
		return fmt.Errorf("duplicate command handler registration with command bus for command of type: %s", typeName)
	}

	bus.handlers[typeName] = handler

	return nil
}

func (bus *CommandBus) Dispatch(cmd CommandMessage) error {
	if handler, ok := bus.handlers[cmd.CommandType()]; ok {
		return handler.Handle(cmd)
	}
	return fmt.Errorf("the command bus does not have a handler for commands of type: %s", cmd.CommandType())
}

type CommandDescriptor struct {
	command interface{}
}

func NewCommandMessage(command interface{}) *CommandDescriptor {
	return &CommandDescriptor{
		command: command,
	}
}

func (c *CommandDescriptor) CommandType() string {
	return TypeOf(c.command)
}

// Command returns the actual command payload of the message.
func (c *CommandDescriptor) Payload() interface{} {
	return c.command
}

func TypeOf(i interface{}) string {
	return reflect.TypeOf(i).Elem().Name()
}
