package logger

// Loggable : Interface representation of logger
type Loggable interface {

	//Set name
	SetName(name string)
}

// Logger struct
type Logger struct {
	Name     string
	Handlers []string
}

// SetName : sets name of logger
func (lgr *Logger) SetName(name string) {
	lgr.Name = name
}
