package layout

import "time"

type Layout interface {
	Format(time time.Time, level, name, message string) string
}
