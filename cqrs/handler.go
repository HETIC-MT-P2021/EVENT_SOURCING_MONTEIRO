package cqrs

type CommandHandler interface {
	Handle(CommandMessage) error
}
