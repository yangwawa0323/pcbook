package utils

import (
	"github.com/fatih/color"
)

type Level int

const (
	DEBUG Level = iota
	WARN
	ERROR
	PANIC
)

type DebugOutput struct {
}

func NewDebugOutput() *DebugOutput {
	return &DebugOutput{}
}

var c *color.Color

func (debug *DebugOutput) Output(level Level, Message ...any) (coloredMessage string) {
	switch level {
	case DEBUG:
		c = color.New(color.FgHiYellow, color.BgCyan)
	case WARN:
		c = color.New(color.FgBlue)
	case ERROR:
		c = color.New(color.FgRed)
	case PANIC:
		c = color.New(color.FgRed, color.BgHiCyan)
	default:
		c = color.New(color.FgWhite)
	}
	formatString := Message[0]
	remainningParam := Message[1:]
	coloredMessage = c.Sprintf(formatString.(string), remainningParam...)
	return
}

func (debug *DebugOutput) Debug(Message ...any) string {
	return debug.Output(DEBUG, Message...)
}

func (debug *DebugOutput) Warn(Message ...any) string {
	return debug.Output(WARN, Message...)
}

func (debug *DebugOutput) Error(Message ...any) string {
	return debug.Output(ERROR, Message...)
}

func (debug *DebugOutput) Panic(Message ...any) string {
	return debug.Output(PANIC, Message...)
}
