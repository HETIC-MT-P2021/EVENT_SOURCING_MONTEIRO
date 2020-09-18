package cqrs

type BusQuery struct {
	Handlers []QueryHandler
}

type Query struct {
	Payload interface{}
}

type QueryHandler struct{}
