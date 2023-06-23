package layout

import (
	"fmt"
	"time"
)

type Standard struct {
	timeFormat string
}

func (s *Standard) Format(time time.Time, level, name, message string) string {
	return fmt.Sprintf("%s |- %s %s - %s\n", time.Format(s.timeFormat), level, name, message)
}

func NewStandard() *Standard {
	return &Standard{
		timeFormat: "2006-01-02 15:04:05.123",
	}
}
