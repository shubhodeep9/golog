package logger

import (
	"github.com/shubhodeep9/golog/logger/handler"
)

// Loggable : Interface representation of logger
type Loggable interface {

	//Set name
	SetName(name string)
}

// Logger struct
type Logger struct {
	Name     string
	Handlers []handler.Handler
}

// SetName : sets name of logger
func (lgr *Logger) SetName(name string) {
	lgr.Name = name
}

// PushHandler : pushes handler to list of handlers
func (lgr *Logger) PushHandler(handlr handler.Handler) {
	lgr.Handlers = append(lgr.Handlers, handlr)
}
