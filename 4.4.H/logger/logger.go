package logger

import (
	"time"

	"4.4.H/logger/exporter"
	"4.4.H/logger/layout"
)

type Logger struct {
	LevelThreshold Level
	Name           string
	Parent         *Logger
	Layout         layout.Layout
	Exporter       exporter.Exporter
}

type LoggerOption struct {
	LevelThreshold Level
	Name           string
	Parent         *Logger
	Layout         layout.Layout
	Exporter       exporter.Exporter
}

func New(opt *LoggerOption) *Logger {
	name := opt.Name
	if name == "" {
		name = "root"
	}

	if opt.Layout == nil {
		opt.Layout = opt.Parent.Layout
	}

	if opt.LevelThreshold == 0 {
		opt.LevelThreshold = opt.Parent.LevelThreshold
	}

	if opt.Exporter == nil {
		opt.Exporter = opt.Parent.Exporter
	}

	return &Logger{
		LevelThreshold: opt.LevelThreshold,
		Name:           name,
		Parent:         opt.Parent,
		Layout:         opt.Layout,
		Exporter:       opt.Exporter,
	}
}

func (l *Logger) log(level Level, message string) {
	if l.LevelThreshold <= level {
		message := l.Layout.Format(time.Now(), level.String(), l.Name, message)
		l.Exporter.Export(message)
	}
}

func (l *Logger) Trace(message string) {
	l.log(LevelTrace, message)
}

func (l *Logger) Debug(message string) {
	l.log(LevelDebug, message)
}

func (l *Logger) Info(message string) {
	l.log(LevelInfo, message)
}

func (l *Logger) Warn(message string) {
	l.log(LevelWarn, message)
}

func (l *Logger) Error(message string) {
	l.log(LevelError, message)
}

func (l *Logger) Fatal(message string) {
	l.log(LevelFatal, message)
}
