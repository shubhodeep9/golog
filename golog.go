// Package golog provides logging services for web apps
package golog

import (
	"github.com/shubhodeep9/golog/logger"
)

// NewLogger creates logger instance
func NewLogger() *logger.Logger {
	lgr := &logger.Logger{}
	return lgr
}
