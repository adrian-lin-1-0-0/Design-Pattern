package mem

import "4.4.H/logger"

var Logger = New()

type Mem struct {
	loggers map[string]*logger.Logger
}

func New() *Mem {
	return &Mem{
		loggers: make(map[string]*logger.Logger),
	}
}

func (m *Mem) Get(name string) *logger.Logger {
	if l, ok := m.loggers[name]; ok {
		return l
	}
	return nil
}

func (m *Mem) Set(name string, l *logger.Logger) {
	m.loggers[name] = l
}

func (m *Mem) SetLoggers(loggers ...*logger.Logger) {
	for _, l := range loggers {
		m.Set(l.Name, l)
	}
}
